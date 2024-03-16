package server

import (
	"context"
	"errors"
	"github.com/geejjoo/library_repository/internal/config"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server interface {
	Serve(ctx context.Context) error
}

type HttpServer struct {
	conf   config.Server
	logger *zap.Logger
	srv    *http.Server
}

func NewHttpServer(conf config.Server, server *http.Server, logger *zap.Logger) *HttpServer {
	return &HttpServer{
		conf:   conf,
		logger: logger,
		srv:    server,
	}
}

func (h *HttpServer) Serve(ctx context.Context) error {
	var err error
	errCh := make(chan error)

	go func() {
		h.logger.Info("server started", zap.String("port", h.conf.Port))
		if err = h.srv.ListenAndServe(); !errors.Is(http.ErrServerClosed, err) {
			h.logger.Error("http listen and serve error", zap.Error(err))
			errCh <- err
		}
	}()

	select {
	case <-errCh:
		return err
	case <-ctx.Done():
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), h.conf.ShutdownTimeout*time.Second)
	defer cancel()
	err = h.srv.Shutdown(ctxShutdown)
	return err
}
