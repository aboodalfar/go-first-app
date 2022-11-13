package main

func main() {
	// var card string = "hi"
	// card = "12" + "12"
	// card := newCard()

	// cards := []string{
	// 	"hi",
	// 	"hdsadsa",
	// 	"byr"}

	//cards := deck{"test"}

	///cards = append(cards, "dsadsa")

	// fmt.Println(cards)

	// fmt.Println(card)
	// fmt.Println(cards2)

	cards := newDeck()

	//cards.print()

	hand, remiingCards := deal(cards, 2)

	hand.print()
	println("--------------------------------------------")
	remiingCards.print()
	println("--------------------------------------------")
	hand2 := deal2(cards, 6)
	cards.print()
	println("--------------------------------------------")
	hand2.print()

	//cards := newDeck()
	//cards.print()
}

// func newCard() string {
// 	return "hhhhhhhhhh"
// }
