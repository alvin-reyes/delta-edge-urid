// It creates a new Echo instance, adds some middleware, creates a new WhyPFS node, creates a new GatewayHandler, and then
// adds a route to the Echo instance
package main

import (
	"github.com/application-research/edge-urid/config"
	_ "net/http"
	"os"

	"github.com/application-research/edge-urid/cmd"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

var (
	log = logging.Logger("edge-ur")
)

var Commit string
var Version string

func main() {

	cfg := config.InitConfig()

	// get all the commands
	var commands []*cli.Command
	commands = append(commands, cmd.DaemonCmd(&cfg)...)

	app := &cli.App{
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
