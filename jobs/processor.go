package jobs

import (
	"context"

	"github.com/alvin-reyes/edge-urid/core"
)

var DELTA_UPLOAD_API = ""
var REPLICATION_FACTOR = "0"

type JobExecutable func() error
type IProcessor interface {
	Info() error
	Run() error
}

type ProcessorInfo struct {
	Name string
}
type Processor struct {
	context   *context.Context
	LightNode *core.LightNode
}
