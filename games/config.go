package games

type GameConfig struct {
	EndPoints  int
	HandSize   int
	MinPlayers int
	MaxPlayers int
}

func NewGameConfig() GameConfig {
	return GameConfig{
		EndPoints:  3,
		HandSize:   10,
		MinPlayers: 3,
		MaxPlayers: 10,
	}
}
