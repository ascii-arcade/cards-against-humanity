package games

func (s *Game) NextTurn() {
	s.withLock(func() {
		winner := s.GetWinner()
		if winner != nil {
			s.Winner = winner
		}

		if len(s.players) > s.CurrentTurnIndex+1 {
			s.CurrentTurnIndex++
		} else {
			s.CurrentTurnIndex = 0
		}
		s.deal()
	})
}
