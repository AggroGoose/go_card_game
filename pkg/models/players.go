package models

type Player struct {
	Name       string
	IsComputer bool
	IsDealer   bool
	IsPlaying  bool
	Hand       CardCollection
}

func (p *Player) AddCard(id int) {
	p.Hand.AddCard(id)
}

func (p *Player) RemoveCard(index int) {
	p.Hand.RemoveCard(index)
}

func CreatePlayer(name string, isComputer bool) Player {
	return Player{
		Name:       name,
		IsComputer: isComputer,
		IsDealer:   false,
		IsPlaying:  true,
		Hand:       NewCardCollection(make([]int, 0), name),
	}
}