package app

import (
	"fmt"
	"github.com/golang-collections/collections/trie"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

type App struct {
	store *trie.Trie
}

func (a *App) Initialize() bool {
	a.store = trie.New()
	a.store.Init()
	err := fetchData(a.store)
	if err != nil {
		log.Fatal(err)
	}
	//fill trie
	//set cron job
	return true
}

func (a *App) Run() {
	http.HandleFunc("/ratings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST"{
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("use root path to see API documentation"))
			return
		}

		scoreConfig, err := validateScoreConfig(r)
		if err != nil {
			msg := fmt.Sprintf("error reading request body: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
			return
		}

		//rate players using config
		players := ratePlayers(scoreConfig, a.store)

		//marshall response body
		bytes, err := yaml.Marshal(players)
		if err != nil {
			msg := fmt.Sprintf("error rating players: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("failed to rate players"))
			return
		}

		//respond with players
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/yaml")
		_, _ = w.Write(bytes)
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
