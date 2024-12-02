package main

import (
	"encoding/csv"
	"fmt"
	"github.com/stantowne/DeckMaker/deckMaker"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	count := os.Args[1] //string
	var numDecks int
	numDecks, _ = strconv.Atoi(count) //int
	var decks [][]string
	decks = deckMaker.DeckMaker(numDecks)
	fmt.Printf("number of decks is: %v\n", len(decks))
	now := time.Now()
	path := "C:/Users/stan/Dropbox/07-programming/program-output/deckMaker/"
	filename := path +
		"decks-made-" +
		strconv.Itoa(now.Year()) + "-" +
		strconv.Itoa(int(now.Month())) + "-" +
		strconv.Itoa(now.Day()) + "-" +
		strconv.Itoa(now.Hour()) + "-" +
		strconv.Itoa(now.Minute()) + "-" +
		"count-" + count + ".csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Cannot create csv file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Cannot close file:", err)
		}
	}(file)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(decks)
	if err != nil {
		log.Println("Cannot write to csv file:", err)
	}
}
