package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"net/http"
)

func main() {
	logrus.Info("app starting...")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port is not set")
	}
	logrus.Info("listening on port " + port)

	gdtdevenv := os.Getenv("GDTDEV_ENV")
	logrus.Info("GDTDEV_ENV = " + gdtdevenv)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":"+ port, nil)
}