package models

import (
	"math/rand"
	"time"
)

type GameDeck struct {
	Cards    map[int]Card
	Drawpile CardCollection
}

func (g *GameDeck) AddCard(id int) {
	g.Drawpile.AddCard(id)
}

func (g *GameDeck) RemoveCard(index int) {
	g.Drawpile.RemoveCard(index)
}

func (g *GameDeck) DealCard(player *Player) {
	card := g.Drawpile.cards[len(g.Drawpile.cards)-1]
	player.AddCard(card)
	g.RemoveCard(len(g.Drawpile.cards) - 1)
}

func (g *GameDeck) GetCards() []int {
	return g.Drawpile.cards
}

func (g *GameDeck) ShuffleDeck() {
	g.Drawpile.ShuffleCards()
}

func NewGameDeck(cards map[int]Card, pile []int) GameDeck {
	return GameDeck{Cards: cards, Drawpile: NewCardCollection(pile, "deck")}
}

type CardCollection struct {
	cards    []int
	position string
}

func (c *CardCollection) AddCard(id int) {
	c.cards = append(c.cards, id)
}

func (c *CardCollection) RemoveCard(index int) {
	if index == 0 {
		c.cards = c.cards[1:]
	} else if index == len(c.cards)-1 {
		c.cards = c.cards[:index]
	} else {
		start := make([]int, 0)
		start = append(start, c.cards[:index]...)
		c.cards = append(start, c.cards[index+1:]...)
	}
}

func (c *CardCollection) ShuffleCards() {
	rand.NewSource(time.Now().Unix())
	rand.Shuffle(len(c.cards), func(i, j int) {c.cards[i], c.cards[j] = c.cards[j], c.cards[i]})
}

func NewCardCollection(cards []int, position string) CardCollection {
	collection := CardCollection{cards: cards, position: position}
	return collection
}

type Card struct {
	suit       string
	name       string
	value      int
	visibility string
}

func (c *Card) ChangeVisibility(visibility string) {
	c.visibility = visibility
}

func NewCard(suit string, name string, value int, visibility string) Card {
	card := Card{suit: suit, name: name, value: value, visibility: visibility}
	return card
}