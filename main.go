package main

func main() {

	cards := newdeck()
	cards = append(cards, "Six of Spades")

	cards1, cards2 := deal(cards, 5)

	cards1.print()
	cards2.print()

	//fmt.Println(cards)

}

func newcard() string {
	return "Five of Diamonds"
}
