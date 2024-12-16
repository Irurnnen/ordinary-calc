package application

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Irurnnen/ordinary-calc/internal/config"
	"github.com/Irurnnen/ordinary-calc/internal/handler"
)

type Application struct {
	Config config.Config
}

func New() *Application {
	return &Application{
		Config: *config.NewConfigExample(),
	}
}

func (a *Application) Run() error {
	// Todo: Add init code for http server
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.CalcHandler)
	err := http.ListenAndServe(":"+fmt.Sprint(a.Config.Port), mux)
	log.Fatal(err)
	return nil
}
