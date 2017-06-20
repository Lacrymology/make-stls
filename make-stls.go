package main

import (
	"os"

	"github.com/Lacrymology/cams"
)

func main() {
	args := os.Args[1:]

	inputCsv := args[0]
	outputStlL := "out-l.stl"
	outputStlR := "out-r.stl"

	cams.CreateCams(
		5,
		0.045,
		0.045,
		cams.MakeCoordinate(-21750, -30536),
		cams.MakeCoordinate(21750, -30536),
		7060,
		inputCsv,
		outputStlL,
		outputStlR)
}
