package main

type Game struct {
	terrorists []*Player
	polices    []*Player
}

func (c *Game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *Game) addPolice(dressType string) {
	player := newPlayer("P", dressType)
	c.polices = append(c.polices, player)
	return
}
