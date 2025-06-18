package games

import "github.com/ascii-arcade/cards-against-humanity/screen"

func (s *Game) NextTurn() {
	s.withLock(func() {
		winner := s.GetWinner()
		if winner != nil {
			for _, player := range s.players {
				player.updateScreen(screen.BoardWinner)
			}
			return
		}

		if len(s.players) > s.CurrentTurnIndex+1 {
			s.CurrentTurnIndex++
		} else {
			s.CurrentTurnIndex = 0
		}
		s.deal()
		s.updateScreens()
	})
}
