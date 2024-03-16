package main

import (
	"github.com/geejjoo/library_repository/internal/config"
	"github.com/geejjoo/library_repository/internal/infrastructure/logs"
	"github.com/geejjoo/library_repository/run"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// load .env
	err := godotenv.Load()
	// config
	conf := config.NewAppConf()
	// logger
	logger := logs.NewLogger(conf, os.Stdout)
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	// app config init
	conf.Init(logger)
	// app init
	app := run.NewApp(conf, logger)
	// exit code
	exitCode := app.Bootstrap().Run()
	os.Exit(exitCode)
}
