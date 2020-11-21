package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "hearts", "clubs", "diamonds"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFileFromDisc(filename string) deck {

	bytes , err := ioutil.ReadFile(filename)

	if err != nil {
	
	}

	var res deck
	for _ , card := range ([]string(bytes)) {
		res = append(res,card)
	}

	return res
}