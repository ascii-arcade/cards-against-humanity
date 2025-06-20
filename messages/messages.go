package messages

import (
	"github.com/ascii-arcade/cards-against-humanity/games"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	PlayerUpdate     int
)
