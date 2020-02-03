package main

import "basketball/app"

func main() {
	basketballPlayerRater := app.App{}
	basketballPlayerRater.Initialize()
	basketballPlayerRater.Run()
}
