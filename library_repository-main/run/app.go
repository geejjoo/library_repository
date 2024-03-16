package run

import (
	"context"
	"errors"
	"fmt"
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/db"
	"github.com/geejjoo/library_repository/internal/infrastructure/component"
	"github.com/geejjoo/library_repository/internal/infrastructure/db/migrate"
	"github.com/geejjoo/library_repository/internal/infrastructure/middleware/auth/guarder"
	"github.com/geejjoo/library_repository/internal/infrastructure/responder"
	"github.com/geejjoo/library_repository/internal/infrastructure/router"
	"github.com/geejjoo/library_repository/internal/infrastructure/server"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/geejjoo/library_repository/internal/modules"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
)

type Application interface {
	Runner
	Bootstraper
}

type Runner interface {
	Run() int
}
type Bootstraper interface {
	Bootstrap(opts ...interface{}) Runner
}

type App struct {
	conf     config.AppConf
	logger   *zap.Logger
	srv      server.Server
	Sig      chan os.Signal
	Storages *modules.Storages
	Services *modules.Services
}

func NewApp(conf config.AppConf, logger *zap.Logger) *App {
	return &App{
		conf:   conf,
		logger: logger,
		Sig:    make(chan os.Signal, 1),
	}
}

func (a *App) Run() int {
	ctx, calcel := context.WithCancel(context.Background())
	defer calcel()
	errGroup, ctx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		sigInt := <-a.Sig
		a.logger.Info("signal interrupt recieved", zap.Stringer("os_signal", sigInt))
		calcel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("app: server error", zap.Error(err))
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return 2
	}
	return 0
}

func (a *App) Bootstrap(opts ...interface{}) Runner {
	// responder
	responseManager := responder.NewResponder(a.logger)
	// components
	components := component.NewComponents(a.conf, responseManager, a.logger, guarder.NewGuarder())
	// SqlDB
	dbx, sqlAdapter, err := db.NewSqlDB(a.conf.DB, a.logger)
	if err != nil {
		a.logger.Fatal("error init db", zap.Error(err))
	}
	// migrator
	migrator := migrate.NewMigrator(dbx, a.conf.DB)
	err = migrator.Migrate(&models.AuthorDTO{}, &models.BookDTO{}, &models.UserDTO{})
	if err != nil {
		a.logger.Fatal("migrator err", zap.Error(err))
	}
	// storages
	storages := modules.NewStorages(sqlAdapter)
	a.Storages = storages
	// services
	services := modules.NewServices(storages, components)
	a.Services = services
	// controllers
	controllers := modules.NewControllers(services, components)
	// router
	r := router.NewRouter(controllers, components)
	// server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", a.conf.Server.Port),
		Handler: r,
	}
	// server init
	a.srv = server.NewHttpServer(a.conf.Server, srv, a.logger)
	// return app
	return a
}
