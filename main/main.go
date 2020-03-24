package main

import (
	"basketball/app"
	"os"
)

func main() {
	basketballPlayerRater := app.App{}
	ok := basketballPlayerRater.Initialize()
	if !ok {
		os.Exit(1)
	}

	ok = basketballPlayerRater.Run()
	if !ok {
		os.Exit(1)
	}
}
