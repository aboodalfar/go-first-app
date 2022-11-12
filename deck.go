package main

import "fmt"

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func newDeck() deck {
// 	cards := deck{}
// 	cardSuits := []string{"bla", "bla2", "bla3"}
// 	cardValues := []string{"hhh", "dsd", "dsds"}

// 	for _, suit := range cardSuits {
// 		for _, value := range cardValues {
// 			cards = append(cards, suit+" of "+value)
// 		}
// 	}
// 	return cards

// }

// func (d deck) print() {
// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }
