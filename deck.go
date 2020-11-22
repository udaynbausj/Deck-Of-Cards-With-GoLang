package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

	var res deck

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(" Error occurred : ", err)
		os.Exit(1)
	}

	for _, card := range strings.Split(string(bytes), ",") {
		res = append(res, card)
	}
	return res
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
