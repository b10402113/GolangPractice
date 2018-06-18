package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsnum := []string{"Ace", "Two", "Three", "Four"}
	cardtype := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for _, typp := range cardtype {
		for _, num := range cardsnum {
			cards = append(cards, num+" of "+typp)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	d = d[handSize:]
	return d[:handSize], d
}
func (d deck) toString() string {

	//cards.toString
	// []string(d) //convert back
	a := []string(d)
	return strings.Join(a, ",") //使用join
}
