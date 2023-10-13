package main

import "fmt"

type finishedHouse struct {
	style                  int    // 0: Chinese; 1: American; 2: European
	centralAirConditioning bool   // true或false
	floorMaterial          string // "ground-tile"或"wood"
	wallMaterial           string // "latex" "paper"或"diatom-mud"
}

type Options struct {
	Style                  int    // 0: Chinese, 1：American, 2:European
	CentralAirConditioning bool   // true or false
	FloorMaterial          string //ground-tile, wood
	WallMaterial           string // latex, paper, diatom-mud
	Color                  string // yellow, black
}

func newFinishedHouse(options *Options) *finishedHouse {
	var style int = 0
	var centralAirConditioning = true
	var floorMaterial = "wood"
	var wallMaterial = "paper"
	var color = "yellow"

	if options != nil {
		style = options.Style
		centralAirConditioning = options.CentralAirConditioning
		floorMaterial = options.FloorMaterial
		wallMaterial = options.WallMaterial
		color = options.Color
	}
	h := &finishedHouse{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
		color:                  color,
	}
	return h
}
func main() {
	fmt.Printf("%+v\n", newFinishedHouse(nil))
	fmt.Printf("%+v\n", newFinishedHouse(&Options{
		Style:                  1,
		CentralAirConditioning: false,
		FloorMaterial:          "ground-tile",
		WallMaterial:           "paper",
	}))
}
