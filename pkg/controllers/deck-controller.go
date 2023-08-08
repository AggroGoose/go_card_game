package controllers

import (
	"strconv"

	m "github.com/StoicGoose/go_card_game/pkg/models"
)

func CreateTraditionalDeck(numDecks int, jokers int) m.GameDeck {
	cardMap := make(map[int]m.Card)
	for i := 0; i < numDecks; i++ {
		generate52Cards(i * 52, cardMap)
	}
	cardTotal := numDecks * 52
	for i := cardTotal; i < cardTotal + jokers; i++ {
		cardMap[i] = m.NewCard("Joker", "Joker", 0, "none")
	}
	var cardArray []int
	for i := 0; i < len(cardMap); i++ {
		cardArray = append(cardArray, i)
	}

	return m.NewGameDeck(cardMap, cardArray)
}

func generate52Cards(startIndex int, cardMap map[int]m.Card) map[int]m.Card {
	iteration := startIndex

	//Creates all cards in the heart suit from Ace, Iterating from 2-10 for numbered Cards, to Jack, Queen, and King cards.

	suit := "heart"
	cardMap[iteration] = m.NewCard(suit, "Ace", 1, "none")
	iteration++
	for i:=2; i <= 10; i++ {
		cardMap[iteration] = m.NewCard(suit, strconv.Itoa(i), i, "none")
		iteration++
	}
	cardMap[iteration] = m.NewCard(suit, "Jack", 11, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "Queen", 12, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "King", 13, "none")
	iteration++

	//Creates all cards in the spade suit from Ace, Iterating from 2-10 for numbered Cards, to Jack, Queen, and King cards.

	suit = "spade"
	cardMap[iteration] = m.NewCard(suit, "Ace", 1, "none")
	iteration++
	for i:=2; i <= 10; i++ {
		cardMap[iteration] = m.NewCard(suit, strconv.Itoa(i), i, "none")
		iteration++
	}
	cardMap[iteration] = m.NewCard(suit, "Jack", 11, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "Queen", 12, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "King", 13, "none")
	iteration++

	//Creates all cards in the diamond suit from Ace, Iterating from 2-10 for numbered Cards, to Jack, Queen, and King cards.

	suit = "diamond"
	cardMap[iteration] = m.NewCard(suit, "Ace", 1, "none")
	iteration++
	for i:=2; i <= 10; i++ {
		cardMap[iteration] = m.NewCard(suit, strconv.Itoa(i), i, "none")
		iteration++
	}
	cardMap[iteration] = m.NewCard(suit, "Jack", 11, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "Queen", 12, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "King", 13, "none")
	iteration++

	//Creates all cards in the club suit from Ace, Iterating from 2-10 for numbered Cards, to Jack, Queen, and King cards.

	suit = "club"
	cardMap[iteration] = m.NewCard(suit, "Ace", 1, "none")
	iteration++
	for i:=2; i <= 10; i++ {
		cardMap[iteration] = m.NewCard(suit, strconv.Itoa(i), i, "none")
		iteration++
	}
	cardMap[iteration] = m.NewCard(suit, "Jack", 11, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "Queen", 12, "none")
	iteration++
	cardMap[iteration] = m.NewCard(suit, "King", 13, "none")
	iteration++
	
	return cardMap
}

