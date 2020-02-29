package app

import (
	"basketball/model"
	"fmt"
	"github.com/golang-collections/collections/trie"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

type App struct {
	store *trie.Trie
}

//#startup
func (a *App) Initialize() bool {
	log.Println("application initializing")
	a.store = trie.New()
	a.store.Init()
	err := fetchData(a.store)
	if err != nil {
		log.Println("unable to fetch player data")
		log.Fatal(err)
	}
	c := cron.New()
	_, err = c.AddFunc("CRON_TZ=America/New_York 00 11 * * *", func() {
		err := fetchData(a.store)
		if err != nil {
			log.Println("unable to set up fetch cron job")
			log.Println(err)
		}
	})

	log.Println("Up and Running!")

	return true
}

//#startup
func (a *App) Run() {
	//#upstream_pushing
	http.HandleFunc("/ratings", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ratings request received")
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("use root path to see API documentation"))
			return
		}

		var scoreConfig model.ScoreConfig
		err := yaml.NewDecoder(r.Body).Decode(&scoreConfig)
		if err != nil {
			msg := fmt.Sprintf("error reading request body: %v", err)
			log.Println(msg)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
			return
		}

		//this block of code is causing zeros to throw errors because of float32 0 value
		//err = validateScoreConfig(scoreConfig)
		//if err != nil {
		//	msg := fmt.Sprintf("invalid score config %v: %v", scoreConfig, err)
		//	log.Println(msg)
		//	w.WriteHeader(http.StatusBadRequest)
		//	_, _ = w.Write([]byte("all fields must be populated with a number between 0.0 and 10.0"))
		//	return
		//}

		log.Println(fmt.Sprintf("rating players for current config: %v", scoreConfig))
		//rate players using config
		players := ratePlayers(scoreConfig, a.store)

		//marshall response body
		bytes, err := yaml.Marshal(players)
		if err != nil {
			msg := fmt.Sprintf("error responding with rated players: %v", err)
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

	//#upstream_pushing
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("use root path to see API documentation"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/yaml")
		log.Println("Skole!")
		_, _ = w.Write([]byte("Skole!"))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Println("failed to boot server")
		log.Fatal(err)
	}
}
