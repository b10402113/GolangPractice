package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck { //代表回傳deck格式
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // _代表我們不想用他
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
} //reciever就是形容會幫deck工作的小工具人 很像this在java self 在python

func deal(d deck, handSize int) (deck, deck) { //handSize是要拿幾張卡 後面代表要回傳兩個值
	return d[:handSize], d[handSize:]
} //跟上面的print不一樣，這是函數

func (d deck) toString() string {

	//cards.toString
	// []string(d) //convert back
	return strings.Join([]string(d), ",") //使用join
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //將bs轉成string
	return deck(s)
}
func (d deck) shuffle() {
	//type rand is a source of random numbers
	souce := rand.NewSource(time.Now().UnixNano())
	r := rand.New(souce)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
