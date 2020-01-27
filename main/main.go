package main

import (
	"basketball/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var players = make([]model.Player, 0)

	jsonFile, err := os.Open("sample.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened sample.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &players)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Unmarshalled sample.json")
	fmt.Println(len(players))
}
