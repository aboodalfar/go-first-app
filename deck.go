package main

import "fmt"

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"bla", "bla2", "bla3"}
	cardValues := []string{"hhh", "dsd", "dsds"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards

}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func deal2(d deck, size int) deck {
	return d[size:]
}

// func (d deck) print() {
// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }
