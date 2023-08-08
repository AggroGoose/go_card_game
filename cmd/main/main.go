package main

import (
	"fmt"

	c "github.com/StoicGoose/go_card_game/pkg/controllers"
	m "github.com/StoicGoose/go_card_game/pkg/models"
)

func main() {
	player1 := m.CreatePlayer("Player 1", false)
	computer1 := m.CreatePlayer("Computer 1", true)
	deck := c.CreateTraditionalDeck(1, 0)
	deck.ShuffleDeck()

	for i := 0; i < 7; i++ {
		fmt.Println("Cards are being dealt to players, total:", i+1)
		deck.DealCard(&player1)
		deck.DealCard(&computer1)
	}

	fmt.Println("Player has the following hand: {}", player1.Hand)
	fmt.Println("Computer has the following hand: {}", computer1.Hand)
	fmt.Println("Here are the cards in the deck: {}", deck.Drawpile)
}