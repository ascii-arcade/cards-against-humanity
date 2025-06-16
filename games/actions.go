package games

func (s *Game) Count(player *Player) {
	s.withLock(func() {
		player.incrementCount()
	})
}

func (s *Game) RevealQuestionCard() {
	s.withLock(func() {
		s.QuestionCard.IsRevealed = true
	})
}

func (s *Game) AddAnswerCard(player *Player, index int) {
	s.withLock(func() {
		if s.QuestionCard.IsRevealed && len(player.Answer.AnswerCards) < s.QuestionCard.Pick {
			player.Answer.add(index)
		}
	})
}

func (s *Game) RemoveAnswerCard(player *Player) {
	s.withLock(func() {
		player.Answer.remove(len(player.Answer.AnswerCards) - 1)
	})
}
