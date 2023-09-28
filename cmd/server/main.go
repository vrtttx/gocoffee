package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vrtttx/gocoffee/database"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

// global port variable for both main and serve
var port = os.Getenv("PORT") // 8080

func (app *Application) Serve() error {
	fmt.Println("API listening on port", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	return srv.ListenAndServe()
}

func main() {
	var cfg Config
	cfg.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := database.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	defer dbConn.DB.Close()

	app := &Application {
		Config: cfg,
	}

	err = app.Serve()

	if err != nil {
		log.Fatal(err)
	}
}