# Go Card Game

I'm writing this card game to familiarize myself more with the GO programming language. This application is being built from scratch without guidance, which is not as terrifying as it sounds.
Functions within may be researched for their specific scope such as how to use the math/rand library or how to add methods to structs, but otherwise the core structs and logic are my own fiddlings. This concept is to challenge myself and be intentional about my research.
That said, let's get into the goal for the app and then we'll move into my key findings that I learned while building the app.

## Card Game App

The goal of the app is to build a skeleton of card functionality that can be worked into multiple types of card games. The structures used for decks and card values are flexible and can be adjusted per game.
Some definitions will be clearly defined such as a Card having a suit, name, value, and visibility, however what 'suits' can be used or what names and values can be given to a card will not be constrained to traditional cards allowing more future flexibility.
Card collections are defined as different locations that cards can be held such as draw piles, player hands, field cards, discard piles, etc. Keeping a general 'CardCollection' field makes it so games can be flexible in the number of locations cards are stored.
_By default, a game is initiated with a core draw pile and a hand for each player_
Cards individually hold a visibility 'string' to allow flexibility in notating who they're visible to. Are they hidden from all players? Visible to all players? Visible only to one player? Visible only to some players? etc...
Visibility was initially going to be defined by collection, but some games like Blackjack or Poker may have mixed visibility like one or more cards visible in a card location while other cards in the location were not visible.
**_More Game Logic to Follow Soon_**

## Findings

Here's a list of key topics I've learned while constructing the card game app:

- **Map Usage** - Originally I was going to design card collections as arrays of card objects, but I figured at scale it was better to store card objects in a map with 'int' keys. The CardCollections would then hold ints to reference the map keys and if card details need to be displayed to the user, a call could be made to reference the cards to the GameDeck's card map. This seems like it would be more performant than constantly moving Card objects around behind the scenes. Had not really understood GO's map function prior to this set up so it was a fun exercise.
- **Pointer Usage** - While I'm still learning how pointers work (prefacing function parameters with & and *) and when to use them in GO, I made use of pointers for the 'card dealing' feature. On initial testing, I had the game deal 7 cards each to the player and computer, and while the main deck was short 14 cards, the player and computer hands remained empty. By prefacing &player in the function call and *Player as the parameter type in the function implementation, it remedied this by adding cards to the player in memory's hand.
