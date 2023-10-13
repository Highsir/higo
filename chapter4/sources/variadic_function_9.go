package main

import "fmt"

type finishHouse struct {
	style                  int
	centralAirConditioning bool
	floorMaterial          string
	wallMaterial           string
}

type Option func(house *finishHouse)

func newFinishHouse(options ...Option) *finishHouse {
	h := &finishHouse{
		// default options
		style:                  0,
		centralAirConditioning: true,
		floorMaterial:          "wood",
		wallMaterial:           "paper",
	}
	for _, option := range options {
		option(h)
	}
	return h
}

func withStyle(style int) Option {
	return func(h *finishHouse) {
		h.style = style
	}
}

func withFloorMaterial(material string) Option {
	return func(h *finishHouse) {
		h.floorMaterial = material
	}
}

func withWallMaterial(material string) Option {
	return func(h *finishHouse) {
		h.wallMaterial = material
	}
}

func withCentralAirConditioning(centralAirConditioning bool) Option {
	return func(h *finishHouse) {
		h.centralAirConditioning = centralAirConditioning
	}
}

func main() {
	fmt.Printf("%+v\n", newFinishHouse()) //默认选项
	fmt.Printf("%+v\n", newFinishHouse(withStyle(1),
		withFloorMaterial("ground-tile"),
		withCentralAirConditioning(false)))
}
