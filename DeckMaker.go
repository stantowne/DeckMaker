package main

import (
	"math/rand"
	"strconv"
)

func DeckMaker(numberOfDecksToBeMade int) [][]string {
	type card [2]int //rank, suit
	var baseDeck []card
	for suit := 0; suit < 4; suit++ {
		for rank := 1; rank < 14; rank++ {
			var c card
			c[0] = rank
			c[1] = suit
			baseDeck = append(baseDeck, c)
		}
	}
	var decks [][]string

	for i := 0; i < numberOfDecksToBeMade; i++ {
		rand.Shuffle(len(baseDeck), func(i, j int) { baseDeck[i], baseDeck[j] = baseDeck[j], baseDeck[i] })
		var deck []string
		for _, c := range baseDeck {
			deck = append(deck, strconv.Itoa(c[0]))
			deck = append(deck, strconv.Itoa(c[1]))
		}
		decks = append(decks, deck)
	}

	return decks
}
