package main

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := dressFactorySingleInstance.getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}
