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
	if strings.TrimRight(input, "\n") == "A" {return 1}
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



func playerTurn() (int, bool) {
	var ace bool
	var status string
	var newCard, sumOfCards int
	var hit = true
	var who = "player"
	status = checkForBlackJackOrBust(sumOfCards)

	for status == "continue" && hit == true {
		newCard = drawRandomCard(who)
		if newCard == 1 {
			ace = true
		}
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
		return sumOfCards, ace

	}

	func  dealersTurn() (int, bool){
		var ace bool
		rand.Seed(time.Now().UnixNano())
		fmt.Println("Now it's the dealer's turn.")
		time.Sleep(1 * time.Second)
		var who = "dealer"
		var status string
		var dealerHitOrStay = 1
		var newCard, sumOfCards int


		newCard = drawRandomCard(who)
		if newCard == 1 {
			ace = true
		}
		time.Sleep(1 * time.Second)
		sumOfCards += newCard
		for dealerHitOrStay != 0 {
			newCard = drawRandomCard(who)
			if newCard == 1 {
				ace = true
			}
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
		return sumOfCards, ace
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
		var sumPlayerCards int
		var sumDealerCards int
		var acePlayer bool
		var aceDealer bool

		welcomeMessage()
		time.Sleep(1 * time.Second)
		fmt.Println("Player, you will go first")
		time.Sleep(1 * time.Second)
		sumPlayerCards, acePlayer = playerTurn()
		sumDealerCards, aceDealer = dealersTurn()

		// check for Aces, adjust accordingly
		if acePlayer == true && sumPlayerCards < 12 {
			sumPlayerCards = sumPlayerCards + 10
			if sumPlayerCards == 21 {
				log.Fatalf("Black jack, black jack, black jack! You've won")
			}
		}
		if aceDealer == true && sumDealerCards < 12 {
			sumDealerCards = sumDealerCards + 10
		}

		if sumPlayerCards > sumDealerCards {
			fmt.Println("You've won!")
		} else if sumPlayerCards == sumDealerCards {
			fmt.Println("It's a tie!")
		} else {
			fmt.Println("You lose :(")
		}
	}

