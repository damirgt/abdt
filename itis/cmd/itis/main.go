package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	//connstring := "postgres://postgres:!Superadmin3000@localhost:5432/Abdt.Kata.Api"
	//db, err := sql.Open("postgres", connstring)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	if db != nil {
		logrus.Info("Успешное соединение с бд!")
	}


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