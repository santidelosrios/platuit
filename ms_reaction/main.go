package main

import (
	"database/sql"
	"net/http"

	"github.com/santidelosrios/platuit/ms_reaction/app"
	"github.com/santidelosrios/platuit/ms_reaction/app/handler"
	"github.com/santidelosrios/platuit/ms_reaction/cmd"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var appName = "reaction-service"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Starting %v\n", appName)

	defaultConfiguration := cmd.DefaultConfiguration()

	client := &http.Client{}

	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}

	client.Transport = transport

	db, err := sql.Open("mysql", "platuit-user:platuit-password@tcp(mysql-db:3306)/platuit")

	if err != nil {
		logrus.WithError(err).Fatal("Error trying to connect to Mysql")
	}

	handler := handler.NewHandler(client, db)

	s := app.NewServer(defaultConfiguration, handler)
	s.SetupRoutes()

	s.Start()

}
