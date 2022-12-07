package app

import (
	"ordersmod/api"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/logger"
)

type app struct {
	conf    *config.Config
	oServer api.Server
	l       logger.Logger
}

type App interface {
	StartApp() error
}

func (a *app) StartApp() error {
	return a.oServer.Start()
}
func NewApp(oServer api.Server, conf *config.Config, l logger.Logger) (App, error) {
	return &app{oServer: oServer, conf: conf, l: l}, nil
}
