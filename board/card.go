package board

import (
	"github.com/ascii-arcade/cards-against-humanity/deck"
	"github.com/ascii-arcade/cards-against-humanity/games"
)

type Answer struct {
	Answers    []*deck.AnswerCard
	IsRevealed bool
	Player     *games.Player
}
