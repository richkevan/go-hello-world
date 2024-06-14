package main

import "fmt"

// Declares a deck type
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle_deck() {}
