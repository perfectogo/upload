package handlers

import (
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/pkg/logger"
	"github.com/perfectogo/upload/service"
)

type CfgHandler struct {
	Service service.InterfaceServer
	Config  config.Config
	Log     logger.Logger
}
type OutHandler struct {
	Service service.InterfaceServer
	Config  config.Config
	Log     logger.Logger
}

func NewHandler(c CfgHandler) *OutHandler {
	return &OutHandler{Service: c.Service}
}
