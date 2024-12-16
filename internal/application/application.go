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
		Config: *config.NewConfigFromEnv(),
	}
}

func (a *Application) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.CalcHandler)

	err := http.ListenAndServe(":"+fmt.Sprint(a.Config.Port), mux)

	log.Fatal(err)
	return nil
}
