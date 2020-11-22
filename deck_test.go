package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 200 {
		t.Errorf(" Expected len of 200, but got len %d", len(d))
	}

}
