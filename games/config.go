package games

type GameConfig struct {
	EndPoints int
	HandSize  int
}

func NewGameConfig() GameConfig {
	return GameConfig{
		EndPoints: 3,
		HandSize:  10,
	}
}
