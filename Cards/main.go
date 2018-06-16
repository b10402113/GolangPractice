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
func main() {
	cards := deck{"Ace of Diamonds", newCard()} //代表string組成的slice
	cards = append(cards, "Six of Spades")      //第一個參數放要加到的地方
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
