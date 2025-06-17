package games

func (s *Game) NextTurn() {
	s.withLock(func() {
		winner := s.GetWinner()
		if winner != nil {
			for _, p := range s.players {
				select {
				case p.UpdateChan <- 1:
				default:
				}
			}
		}

		if len(s.players) > s.CurrentTurnIndex+1 {
			s.CurrentTurnIndex++
		} else {
			s.CurrentTurnIndex = 0
		}
		s.deal()
	})
}
