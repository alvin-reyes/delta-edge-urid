package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alvin-reyes/edge-urid/core"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/xerrors"
)

var (
	OsSignal chan os.Signal
	log      = logging.Logger("router")
)

type HttpError struct {
	Code    int    `json:"code,omitempty"`
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

func (he HttpError) Error() string {
	if he.Details == "" {
		return he.Reason
	}
	return he.Reason + ": " + he.Details
}

type HttpErrorResponse struct {
	Error HttpError `json:"error"`
}

type AuthResponse struct {
	Result struct {
		Validated bool   `json:"validated"`
		Details   string `json:"details"`
	} `json:"result"`
}

func GetDefaultTagPolicy(ln *core.LightNode) error {

	// remove the current default tag policy
	if err := ln.DB.Model(&core.Policy{}).Where("name = ?", ln.Config.Node.DefaultCollectionName).Delete(&core.Policy{}).Error; err != nil {
		return xerrors.Errorf("failed to remove default tag policy: %w", err)
	}

	// create a new default tag policy
	newTagPolicy := core.Policy{
		Name:       ln.Config.Node.DefaultCollectionName,
		BucketSize: ln.Config.Common.BucketAggregateSize,
		SplitSize:  ln.Config.Common.SplitSize,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	ln.DB.Model(&core.Policy{}).Create(&newTagPolicy)

	return nil
}

// RouterConfig configures the API node
func InitializeEchoRouterConfig(ln *core.LightNode) {
	// Echo instance
	// Echo instance
	e := echo.New()
	e.File("/", "templates/index.html")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = ErrorHandler

	defaultOpenRoute := e.Group("")
	ConfigureGatewayRouter(defaultOpenRoute, ln) // access to light node
	ConfigureStatsRouter(defaultOpenRoute, ln)
	ConfigureHealthCheckRouter(defaultOpenRoute, ln)
	ConfigureNodeInfoRouter(defaultOpenRoute, ln)

	apiGroup := e.Group("/api/v1")

	ConfigureRetrieveRouter(apiGroup, ln)
	ConfigureUploadRouter(apiGroup, ln)
	ConfigureBucketsRouter(defaultOpenRoute, ln)
	ConfigureCollectionsRouter(defaultOpenRoute, ln)
	ConfigureStatusCheckRouter(apiGroup, ln)
	ConfigureStatusOpenCheckRouter(defaultOpenRoute, ln)

	// Start server

	addrPort := fmt.Sprintf("0.0.0.0:%d", ln.Config.Node.Port)
	e.Logger.Fatal(e.Start(addrPort)) // configuration
}

func ErrorHandler(err error, c echo.Context) {
	var httpRespErr *HttpError
	if xerrors.As(err, &httpRespErr) {
		log.Errorf("handler error: %s", err)
		if err := c.JSON(httpRespErr.Code, HttpErrorResponse{Error: *httpRespErr}); err != nil {
			log.Errorf("handler error: %s", err)
			return
		}
		return
	}

	var echoErr *echo.HTTPError
	if xerrors.As(err, &echoErr) {
		if err := c.JSON(echoErr.Code, HttpErrorResponse{
			Error: HttpError{
				Code:    echoErr.Code,
				Reason:  http.StatusText(echoErr.Code),
				Details: echoErr.Message.(string),
			},
		}); err != nil {
			log.Errorf("handler error: %s", err)
			return
		}
		return
	}

	log.Errorf("handler error: %s", err)
	if err := c.JSON(http.StatusInternalServerError, HttpErrorResponse{
		Error: HttpError{
			Code:    http.StatusInternalServerError,
			Reason:  http.StatusText(http.StatusInternalServerError),
			Details: err.Error(),
		},
	}); err != nil {
		log.Errorf("handler error: %s", err)
		return
	}
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")
}
