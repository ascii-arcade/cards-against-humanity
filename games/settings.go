package games

type Settings struct {
	EndPoints  int
	HandSize   int
	MinPlayers int
	MaxPlayers int
}

func NewSettings() Settings {
	return Settings{
		EndPoints:  1,
		HandSize:   10,
		MinPlayers: 2,
		MaxPlayers: 10,
	}
}
