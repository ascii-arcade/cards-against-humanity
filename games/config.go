package games

type GameConfig struct {
	HandSize int
}

func NewGameConfig() GameConfig {
	return GameConfig{
		HandSize: 10,
	}
}
