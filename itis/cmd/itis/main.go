package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Info("app starting...")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}
	logrus.Info("listening on port " + port)
}