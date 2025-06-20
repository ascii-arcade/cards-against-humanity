package games

type Settings struct {
	EndPoints  int
	HandSize   int
	MinPlayers int
	MaxPlayers int
}

func NewSettings() Settings {
	return Settings{
		EndPoints:  3,
		HandSize:   10,
		MinPlayers: 3,
		MaxPlayers: 10,
	}
}
