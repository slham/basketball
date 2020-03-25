package app

import (
	"basketball/env"
	"basketball/handlers"
	"basketball/storage"
	"github.com/gorilla/mux"
	"github.com/slham/toolbelt/l"
	"net/http"
	"os"
)

type App struct {
	Config env.Config
	Router *mux.Router
}

func (a *App) Initialize() bool {
	l.Info(nil, "application initializing")

	a.Config.Env = os.Getenv("ENVIRONMENT")
	a.Config.L.Level = l.Level(os.Getenv("LOG_LEVEL"))
	a.Config.Runtime.Port = os.Getenv("RUNTIME_PORT")
	a.Config.Storage.Bucket = os.Getenv("STORAGE_BUCKET")
	a.Config.Storage.Prefix = os.Getenv("STORAGE_PREFIX")
	a.Config.Storage.FileName = os.Getenv("STORAGE_FILENAME")

	ok := l.Initialize(a.Config.L.Level)
	if !ok {
		l.Error(nil, "failed to initialize logging middleware")
		return false
	}

	ok = storage.Initialize(a.Config)
	if !ok {
		l.Error(nil, "failed to initialize storage")
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
