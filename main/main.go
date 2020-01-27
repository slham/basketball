package main

import (
	"basketball/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var urls = make(map[string]string, 0)
	urls["players"] = "https://api-nba-v1.p.rapidapi.com/players/teamId/%v" //1-30
	urls["player"] = "https://api-nba-v1.p.rapidapi.com/players/playerId/%v"
	urls["stats"] = "https://api-nba-v1.p.rapidapi.com/statistics/players/playerId/%d"

	var playerStats = make(map[string]model.Player, 0)

	/*
	for each team
	  get each teams players
	  marshal into []Players
	  for each player
	    get their stats
	 */

	for i := 1; i <= 30; i++{
		url := fmt.Sprintf(urls["players"], i)
		var playersRes model.Response
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "0ace5be181msh231d6cf57a0164cp13596cjsn1d6af7ad8df0")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()

		if res.StatusCode != 200{
			log.Fatal(res.Status)
		}
		err = json.NewDecoder(res.Body).Decode(&playersRes)
		if err != nil {
			log.Fatal(err)
		}

		for _, player := range playersRes.Api.Players {
			url := fmt.Sprintf(urls["player"], player.Id)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil{
				log.Fatal(err)
			}
			req.Header.Add("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
			req.Header.Add("x-rapidapi-key", "0ace5be181msh231d6cf57a0164cp13596cjsn1d6af7ad8df0")

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			if res.StatusCode != 200{
				log.Fatal(res.Status)
			}

			var playerRes model.Response
			err = json.NewDecoder(res.Body).Decode(&playerRes)
			if err != nil{
				log.Fatal(err)
			}
			player.FirstName = playerRes.Api.Players[0].FirstName
			player.LastName = playerRes.Api.Players[0].LastName
			player.CreatedDateTime = time.Now()
			player.UpdatedDateTime = time.Now()
			playerStats[player.Id] = player
		}
	}

	for k, v := range playerStats {
		url := fmt.Sprintf(urls["stats"], k)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil{
			log.Fatal(err)
		}
		req.Header.Add("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", "0ace5be181msh231d6cf57a0164cp13596cjsn1d6af7ad8df0")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		if res.StatusCode != 200{
			log.Fatal(res.Status)
		}

		var statsRes model.Response
		err = json.NewDecoder(res.Body).Decode(&statsRes)
		if err != nil{
			log.Fatal(err)
		}

		for _, player := range statsRes.Api.Statistics{
			v.NumGames ++
			v.Min = player.Min
			v.Fgm = player.Fgm
			v.Fga = player.Fga
			v.Fgp = player.Fgp
			v.Ftm = player.Fgm
			v.Fta = player.Fga
			v.Fgp = player.Fgp
			v.Tpm = player.Tpm
			v.Tpa = player.Tpa
			v.Tpp = player.Tpp
			v.Reb = player.Reb
			v.Ass = player.Ass
			v.Stl = player.Stl
			v.Blk = player.Blk
			v.Turnovers = player.Turnovers
			//double doubles
			v.CheckDds()
			v.Pts = player.Pts
			v.UpdatedDateTime = time.Now()
		}
	}
	log.Println(fmt.Sprintf("All done. Players: %v", playerStats))
	//url := "https://api-nba-v1.p.rapidapi.com/statistics/players/playerId/75"
	//
	//req, _ := http.NewRequest("GET", url, nil)
	//
	//req.Header.Add("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
	//req.Header.Add("x-rapidapi-key", "0ace5be181msh231d6cf57a0164cp13596cjsn1d6af7ad8df0")
	//
	//res, _ := http.DefaultClient.Do(req)
	//
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//
	//fmt.Println(res)
	//fmt.Println(string(body))

	//for page <= max {
	//	players, meta = fetchPlayers(page, limit)
	//	if meta.TotalPages > max {
	//		max = meta.TotalPages
	//	}
	//	log.Println(fmt.Sprintf("found %v players on page %v", len(players), page))
	//	for _, a := range players {
	//		log.Println(fmt.Sprintf("id:%d::name:%v %v", a.Id, a.FirstName, a.LastName))
	//		a.CreatedDateTime = time.Now()
	//		a.UpdatedDateTime = time.Now()
	//		store[a.Id] = a
	//		c := fetchSeasonAvgs(a.Id)
	//		c.CreatedDateTime = time.Now()
	//		c.UpdatedDateTime = time.Now()
	//		player := model.Merge(a, c)
	//		log.Println(fmt.Sprintf("player:%v", player))
	//	}
	//
	//	page++
	//}
	//
	//log.Println(fmt.Sprintf("total players: %v", len(store)))
	//log.Println(fmt.Sprintf("deets:%v::deets:%v", store[0], store[len(store)]))
}

//func fetchPlayers(page, limit int) (players []model.Player, meta model.Meta) {
//	playersRes := model.PlayersRes{}
//	meta = model.Meta{}
//	players = make([]model.Player, 0)
//	url := fmt.Sprintf("https://www.balldontlie.io/api/v1/players?page=%v&per_page=%v", page, limit)
//	res, err := http.Get(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.NewDecoder(res.Body).Decode(&playersRes)
//	_ = res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	players, meta = playersRes.Data, playersRes.Meta
//	return players, meta
//}
//
//func fetchPlayer(id int) (player model.Player) {
//	var playerRes model.PlayerRes
//	player = model.Player{Id: -1}
//	url := fmt.Sprintf("https://www.balldontlie.io/api/v1/players/%d", id)
//	res, err := http.Get(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.NewDecoder(res.Body).Decode(&playerRes)
//	_ = res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(fmt.Sprintf("player res:%v", playerRes))
//	player = playerRes.Player
//	return player
//}

//func fetchSeasonAvgs(id int) (player model.Player) {
//	var seasonRes model.SeasonRes
//	player = model.Player{Id: -1}
//	url := fmt.Sprintf("https://balldontlie-example.herokuapp.com/player/average?season=2019&playerIds[]=%d", id)
//	res, err := http.Get(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = json.NewDecoder(res.Body).Decode(&seasonRes)
//	_ = res.Body.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//	if len(seasonRes.Player) == 1 {
//		log.Println(fmt.Sprintf("player:%v", seasonRes.Player))
//		for _, v := range seasonRes.Player {
//			player = v
//		}
//		player.Id = player.Data.PlayerId
//		log.Println(fmt.Sprintf("SUGARTITS:%v", player))
//	} else {
//		log.Println(fmt.Sprintf("BUBKIS:%v", player))
//	}
//	return player
//}
