package handlers

import (
	"github.com/perfectogo/upload/config"
	"github.com/perfectogo/upload/service"
)

type CfgHandler struct {
	Service service.InterfaceServer
	Config  config.Config
}
type OutHandler struct {
	Service service.InterfaceServer
	Config  config.Config
}

func NewHandler(c CfgHandler) *OutHandler {
	return &OutHandler{Service: c.Service}
}
