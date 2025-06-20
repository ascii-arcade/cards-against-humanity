package games

import (
	"github.com/ascii-arcade/cards-against-humanity/messages"
)

func (s *Game) setPlayerScreens() {
	winner := s.GetWinner()
	if winner != nil {
		for _, p := range s.players {
			p.update(messages.WinnerScreen)
		}
		return
	}

	for _, p := range s.players {
		switch {
		case s.GetCurrentPlayer() == p:
			p.update(messages.RevealScreen)
		case p.Answer.IsLocked:
			p.update(messages.RevealScreen)
		default:
			p.update(messages.BuildAnswerScreen)
		}
	}
}
