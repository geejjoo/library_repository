package component

import (
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/infrastructure/middleware/auth/guarder"
	"github.com/geejjoo/library_repository/internal/infrastructure/responder"
	"go.uber.org/zap"
)

type Components struct {
	Conf      config.AppConf
	Responder responder.Responder
	Logger    *zap.Logger
	Guarder   guarder.Guarder
}

func NewComponents(conf config.AppConf, responder responder.Responder, logger *zap.Logger, guarder guarder.Guarder) *Components {
	return &Components{
		Conf:      conf,
		Responder: responder,
		Logger:    logger,
		Guarder:   guarder,
	}
}
