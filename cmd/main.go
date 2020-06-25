package main

import (
	"os"
	"os/signal"

	"github.com/sergiosegrera/short/db/redisdb"
	"github.com/sergiosegrera/short/service"
	"github.com/sergiosegrera/short/transports/http"
	"go.uber.org/zap"
)

func main() {
	// Start logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Connect to db
	db, err := redisdb.New("localhost:6379")
	if err != nil {
		logger.Fatal("error connecting to db")
	}

	shortService := &service.ShortService{
		DB:     db,
		Logger: logger,
	}

	go func() {
		logger.Info("starting the http server", zap.String("port", "8080"))
		err := http.Serve(shortService)
		if err != nil {
			logger.Error("http server panic", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info("exited", zap.String("sig", sig.String()))
}
