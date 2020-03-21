package app

import (
	"basketball/env"
	"basketball/handlers"
	"basketball/storage"
	"github.com/gorilla/mux"
	"github.com/slham/toolbelt/l"
	"net/http"
)

type App struct {
	Config env.Config
	Router *mux.Router
}

func (a *App) Initialize(environment string) bool {
	l.Info(nil, "application initializing")

	config, ok := env.Load(environment)
	if !ok {
		return false
	}

	a.Config = config

	ok = l.Initialize(a.Config.L.Mode)
	if !ok {
		return false
	}

	ok = storage.Initialize(a.Config)
	if !ok {
		return false
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
	l.Info(nil, "Up and Running!")

	return true
}

func (a *App) Run() bool {
	if err := http.ListenAndServe(":"+a.Config.Runtime.Port, a.Router); err != nil {
		l.Error(nil, "failed to boot server: %v", err)
		return false
	}
	return true
}

func (a *App) initializeRoutes() {
	a.Router.Use(l.Logging)
	a.Router.Methods("GET").Path("/health").HandlerFunc(handlers.HealthCheck)

	a.Router.Methods("POST").Path("/ratings").HandlerFunc(handlers.RatePlayers)

	a.Router.Methods("PUT").Path("/players").HandlerFunc(handlers.StorePlayers)
}
