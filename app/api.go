package app

import (
	"basketball/env"
	"basketball/handlers"
	"basketball/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Config env.Config
	Router *mux.Router
}

func (a *App) Initialize(environment string) bool {
	config, ok := env.Load(environment)
	if !ok {
		return false
	}

	log.Println("application initializing")

	a.Config = config
	a.Router = mux.NewRouter()
	a.initializeRoutes()

	ok = storage.Initialize(a.Config)
	if !ok {
		return false
	}

	log.Println("Up and Running!")

	return true
}

func (a *App) Run() bool {
	if err := http.ListenAndServe(":"+a.Config.Runtime.Port, a.Router); err != nil {
		log.Println("failed to boot server")
		log.Println(err)
		return false
	}
	return true
}

func (a *App) initializeRoutes() {
	a.Router.Methods("GET").Path("/health").HandlerFunc(handlers.HealthCheck)

	a.Router.Methods("POST").Path("/ratings").HandlerFunc(handlers.RatePlayers)

	a.Router.Methods("PUT").Path("/players").HandlerFunc(handlers.StorePlayers)
}
