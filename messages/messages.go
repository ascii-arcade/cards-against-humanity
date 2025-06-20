package messages

import (
	"github.com/ascii-arcade/cards-against-humanity/games"
)

type (
	SwitchToBoardMsg struct{ Game *games.Game }
	PlayerUpdate     int
)
