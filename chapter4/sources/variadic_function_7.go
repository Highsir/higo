package main

import "fmt"

type FinishedHouse struct {
	style                  int    //// 0: Chinese; 1: American; 2: European
	centralAirConditioning bool   // true or false
	floorMaterial          string //ground-tile wood
	wallMaterial           string // latex, paper or diatom-mud
}

func NewFinishedHouse(style int, centralAirConditioning bool, floorMaterial string,
	wallMaterial string) *FinishedHouse {

	h := &FinishedHouse{
		style:                  style,
		centralAirConditioning: centralAirConditioning,
		floorMaterial:          floorMaterial,
		wallMaterial:           wallMaterial,
	}
	return h
}

func main() {
	fmt.Printf("%+v\n", NewFinishedHouse(0, true, "wood", "paper"))
}
