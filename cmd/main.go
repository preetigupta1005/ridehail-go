package main

import (
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/sirupsen/logrus"
)

const shutDownTimeOut = 10 * time.Second

func main() {

	if err := database.ConnectDB(); err != nil {
		logrus.Panicf("failed to connect database: %+v", err)
	}

	if err := database.Close(); err != nil {
		logrus.WithError(err).Error("failed to close database connection")
	}
}
