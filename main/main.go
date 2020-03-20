package main

import (
	"basketball/app"
	"flag"
	"os"
)

func main() {
	env := flag.String("env", "local", "a string")
	flag.Parse()

	basketballPlayerRater := app.App{}
	ok := basketballPlayerRater.Initialize(*env)
	if !ok {
		os.Exit(2)
	}

	ok = basketballPlayerRater.Run()
	if !ok {
		os.Exit(3)
	}
}
