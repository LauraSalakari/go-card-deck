package main

// VSCode recognises packages and automatically imports them if we haven't declared them

func main() {
	// long form way to declare a variable
	// var card string = "Ace of Spades"
	// abbreviated variable declaration
	// card := "Ace of Spades"
	// reassign the variable
	// card = "Five of Diamonds"

	// use newCard() to create a card
	// card := newCard()
	// fmt.Println(card)

	// create a slice using newCard(), append to slice
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// assign cards the return value of the newDeck() function
	// cards := newDeck()
	// cards.print()

	// "deal a hand" with deal()
	// store eachof the return values in a variable, keep track of the order of returns!!
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// save a deck in the system storage
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my-cards")

	// retrieve a deck from system storage
	// cards := newDeckFromFile("my-car")
	// cards.print()

	// shuffle a deck using shuffle()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// function return type goes between the args and the body (ie between round and squiggly brackets)
// func newCard() string {
// 	return "Five of Diamonds"
// }
