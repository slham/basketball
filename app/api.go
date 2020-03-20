package app

import (
	"basketball/handlers"
	"basketball/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() bool {
	log.Println("application initializing")
	a.Router = mux.NewRouter()
	initializeRoutes(a)
	err := storage.Initialize()
	if err != nil {
		log.Println("unable to fetch player data")
		log.Fatal(err)
	}

	log.Println("Up and Running!")

	return true
}

func (a *App) Run() {
	if err := http.ListenAndServe(":80", a.Router); err != nil {
		log.Println("failed to boot server")
		log.Fatal(err)
	}
}

func initializeRoutes(a *App) {
	a.Router.Methods("GET").Path("/health").HandlerFunc(handlers.HealthCheck)

	a.Router.Methods("POST").Path("/ratings").HandlerFunc(handlers.RatePlayers)

	a.Router.Methods("PUT").Path("/players").HandlerFunc(handlers.StorePlayers)
}
