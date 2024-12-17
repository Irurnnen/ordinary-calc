package application

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Irurnnen/ordinary-calc/docs"
	"github.com/Irurnnen/ordinary-calc/internal/config"
	"github.com/Irurnnen/ordinary-calc/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Application struct {
	Config config.Config
	Debug  bool
}

func New() *Application {
	return &Application{
		Config: *config.NewConfigFromEnv(),
		Debug:  false,
	}
}

func NewDebug() *Application {
	return &Application{
		Config: *config.NewConfigFromEnv(),
		Debug:  true,
	}
}

func (a *Application) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	if a.Debug {
		r.Get("/swagger/*", httpSwagger.Handler(
		// httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", a.Config.Port)),
		))
	}

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/calculate", handler.CalcHandler)
		})
	})

	// mux := http.NewServeMux()
	// mux.HandleFunc("/api/v1/calculate", handler.CalcHandler)

	err := http.ListenAndServe(":"+fmt.Sprint(a.Config.Port), r)

	log.Fatal(err)
	return nil
}
