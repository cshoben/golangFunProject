package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func welcomeMessage () {
	fmt.Println("Welcome to Black Jack!")
}

// 	This is a variation of Black Jack, where the player is dealt all their cards, and
// 	then the dealer is dealt theirs. The dealers "choice" to hit or stay is based on a random number generator. 

func cardValue(input string) int{
	if strings.TrimRight(input, "\n") == "0" {return 0}
	if strings.TrimRight(input, "\n") == "2" {return 2}
	if strings.TrimRight(input, "\n") == "3" {return 3}
	if strings.TrimRight(input, "\n") == "4" {return 4}
	if strings.TrimRight(input, "\n") == "5" {return 5}
	if strings.TrimRight(input, "\n") == "6" {return 6}
	if strings.TrimRight(input, "\n") == "7" {return 7}
	if strings.TrimRight(input, "\n") == "8" {return 8}
	if strings.TrimRight(input, "\n") == "9" {return 9}
	if strings.TrimRight(input, "\n") == "10" {return 10}
	if strings.TrimRight(input, "\n") == "J" {return 10}
	if strings.TrimRight(input, "\n") == "Q" {return 10}
	if strings.TrimRight(input, "\n") == "K" {return 10}
	if strings.TrimRight(input, "\n") == "A" {return 11}
	return 99
}

func drawRandomCard(whoPlaying string) int {
	rand.Seed(time.Now().UnixNano())
	cards := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
		"A",
	}

	n := rand.Int() % len(cards)
	if whoPlaying == "player" {
		fmt.Printf("You've been dealt a(n) %s\n", cards[n])
	} else {
		fmt.Printf("The dealer has been dealt a(n) %s\n", cards[n])
	}
	return cardValue(cards[n])
}

func checkForBlackJackOrBust(sumOfCards int) string {
	//have true if BJ or bust, and false if under 21.
	time.Sleep(1 * time.Second)
	if sumOfCards > 21 {
		return "bust"
	} else if sumOfCards == 21 {
		log.Fatalf("BLACK JACK! Black Jack! black jack! You've won!")
	} else {
		return "continue"
	}
	return ""
}



func playerTurn () int {

	var status string
	var newCard, sumOfCards int
	var hit = true
	var who = "player"
	status = checkForBlackJackOrBust(sumOfCards)

	for status == "continue" && hit == true {
		newCard = drawRandomCard(who)
		sumOfCards += newCard
		status = checkForBlackJackOrBust(sumOfCards)
		if status == "blackJack" {
			log.Fatalln("Black jack, black jack, black jack! You've won")
		}
		if status == "bust" {
			log.Fatalf("BUST! Your card total is %v, which is over 21. Try again.", sumOfCards)
		}
		hit = hitOrStay()
	}
		return -1

	}

	func  dealersTurn() int{
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Now it's the dealer's turn.")
		time.Sleep(1 * time.Second)
		var who = "dealer"
		var status string
		var dealerHitOrStay = 1
		var newCard, sumOfCards int


		newCard = drawRandomCard(who)
		time.Sleep(1 * time.Second)
		sumOfCards += newCard
		for dealerHitOrStay != 0 {
			newCard = drawRandomCard(who)
			time.Sleep(1 * time.Second)
			sumOfCards += newCard
			status = checkForBlackJackOrBust(sumOfCards)
			time.Sleep(1 * time.Second)
			if status == "blackJack" {
				log.Fatalln("Black jack! The dealer has won!")
			}
			if status == "bust" {
				log.Fatalf("The Dealer BUSTED with a card total of %v, You've won!", sumOfCards)
			}
			dealerHitOrStay = rand.Int()%sumOfCards
		}
		return -1
	}

	func hitOrStay() bool{

		fmt.Println("Player, would you like a hit or to stay?")
		time.Sleep(1 * time.Second)
		fmt.Println("Enter 'h' for hit and 's' for stay")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil{
		fmt.Println("An error occurred while reading input. Please try again", err)
	}
		input = strings.TrimSuffix(input, "\n")

		if input == "h"{
		return true
	} else if input == "s"{
		fmt.Println("You have chosen to stay")
		return false
	} else{
		fmt.Println("Invalid entry, please type either 'h' or 's'.")
		return hitOrStay()
		}
	}

	func main() {
		var sumOfPlayerCards int
		var sumDealerCards int

		welcomeMessage()
		time.Sleep(1 * time.Second)
		fmt.Println("Player, you will go first")
		time.Sleep(1 * time.Second)
		sumOfPlayerCards = playerTurn()
		sumDealerCards = dealersTurn()
		if sumOfPlayerCards < sumDealerCards {
			fmt.Println("You've won!")
		} else if sumOfPlayerCards == sumDealerCards {
			fmt.Println("It's a tie!")
		} else {
			fmt.Println("You lose :(")
		}
	}



//Goal: write a program so a player can play black jack via the terminal.

/*
Rules of black jack
The player and the dealer receive two cards from a shuffled deck.
In our case, we'll use a single deck, though casinos usually use a
'shoe' consisting of six decks.

After the first two cards are dealt to dealer and player, the player
is asked if they'd like another card (called 'hitting'), or if they
are happy with the cards they have already (called 'staying'). The
object is to make the sum of your card values as close to 21, without
going over. If we make 21 exactly, we have blackjack, which can't be beat.
If we go over 21, we 'bust' and we lose the round. The player is allowed to
stop hitting at any point.

The number cards (2 through 10) are worth the number displayed, face
cards are worth 10, and an Ace can be worth either 1 or 11. For example,
if our first two cards are a Jack and an Ace, we'd want to count the Ace
as 11 since 10 + 11 = 21, and we'd have blackjack, but, if we had
already had a hand worth 18, decided to hit, and got an Ace, we'd
want to count it as 1, since counting it as 11 would put us at 29, and we'd bust.

Once our hand is finished, the dealer tries to do the same. The dealer must
keep hitting until they get to 17. If they get above 17 without busting, they can stay.

*/

