package main

import "fmt"

type deck []string

func newDec() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamond", "Heart", "Cluste "}
	cardValues := []string{"Ace", "One", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
