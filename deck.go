package main

import "fmt"

type deck []string

func newdeck() deck {

	cards := deck{}

	cardsuits := []string{"Spades", "Hearts", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardsuits {
		for _, value := range cardvalues {
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

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
