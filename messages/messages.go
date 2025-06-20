package messages

import (
	"github.com/ascii-arcade/cards-against-humanity/games"
	"github.com/ascii-arcade/cards-against-humanity/screen"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	SwitchScreenMsg  struct{ Screen screen.Screen }
	PlayerUpdate     int
)
