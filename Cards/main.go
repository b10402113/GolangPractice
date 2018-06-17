package main

//015
// var card string = "Ace of Spades"
// card := newCard() //不須打出型態go會自己認
// fmt.Println(card)
// newCard()
//016
// cards := []string{"Ace of Diamonds", newCard()} //代表string組成的slice
// cards = append(cards, "Six of Spades")          //第一個參數放要加到的地方
// for i, card := range cards {
// 	fmt.Println(i, card)
// } //對每個元素做迭代第一個arg是index第二個arg是元素
//017
// cards := newDeck()
// // cards.print()

// hand, remainingDeck := deal(cards, 5) //因為在同一個package可以直接呼叫
//為何不用cards.deal? 因為這樣代表直接修改牌組
// hand.print()
// remainingDeck.print()
func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cs")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
