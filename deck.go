package main

// note the syntax to import multiple packages!
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type "deck"
// slice of strings
// it will "borrow/extend" the behaviour of a slice of strings

type deck []string

// we can create functions that are particularly attached to the new type!

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function syntax
// func (varName receiver) funcName {}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function with multiple return values!
// func funcName (argName argType, argName argType)(returnType, secondReturnType)
func deal(d deck, handSize int) (deck, deck) {
	// return two decks, one up to the specified int and one starting from that index!
	return d[:handSize], d[handSize:]
}

// a helper function that will turn the deck into a string
func (d deck) toString() string {
	// this works, becuase deck essentially is a []string
	return strings.Join([]string(d), ",")
}

// return type is error because that is what the ioutil function might return so that we dont have to hunt for it
// this function creates a plain text file containing the deck which calls it
func (d deck) saveToFile(filename string) error {
	// WriteFile syntax: (filename string, data [byte], perm os.FileMode)
	// 0666 is generic permissions that allow anyone read or write to the created file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// the ReadFile function returns a []byte (bs) and an error (err)

	// error handling: (note conditional syntax!)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log error and entirely quit the program
		fmt.Println("Error:", err)
		// exit the program with error code 1
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s) // we can do this conversion because deck is really a []string
}

func (d deck) shuffle() {
	// in order to create a truly random order, we need to randomise the seed of the math/rand package everytime
	// we use the current time and turn it into a int64 to create a unique source
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// use the math/rand package to generate a random index between 0 and the length of slice -1
		newPos := r.Intn(len(d) - 1)

		// assign the elem at newPos to i and the elem at i to newPos
		d[i], d[newPos] = d[newPos], d[i]
	}
}
