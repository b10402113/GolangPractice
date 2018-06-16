package main

import (
	"fmt"
)

// Create a new type of 'deck'
//which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
} //reciever就是形容會幫deck工作的小工具人 很像this在java self 在python
