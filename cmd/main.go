package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/server"
	"github.com/preetigupta1005/ridehail-go/websocket"

	"github.com/sirupsen/logrus"
)

const shutDownTimeOut = 10 * time.Second

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var hub = websocket.NewHub()

	if err := database.ConnectDB(); err != nil {
		logrus.Panicf("failed to connect database: %+v", err)
	}

	srv := server.SetUpRoutes(hub)

	go func() {
		if err := srv.Run(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Panicf("failed to run server: %+v", err)
		}
	}()
	logrus.Print("server started at port:8080")

	<-done
	logrus.Info("shutting down server")

	if err := database.Close(); err != nil {
		logrus.WithError(err).Error("failed to close database connection")
	}

	if err := srv.Shutdown(shutDownTimeOut); err != nil {
		logrus.WithError(err).Panic("failed to gracefully shutdown server")
	}
}
