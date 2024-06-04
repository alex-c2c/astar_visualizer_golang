package main

import (
	"flag"
	"fyne.io/fyne/v2/app"
	"visuals"
)

func main() {
	colFlag := flag.Int("col", -1, "map col size")
	rowFlag := flag.Int("row", -1, "map row size")

	flag.Parse()

	app := app.New()

	if *colFlag > 0 && *rowFlag > 0 {
		visuals.CreateVisualizerWindow(app, *colFlag, *rowFlag)
	} else {
		visuals.CreateMapSizeInputWindow(app)
	}

	app.Run()
}
