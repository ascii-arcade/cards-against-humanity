package games

import (
	"errors"
)

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

func (s *Game) LockAnswer(player *Player) error {
	return s.withErrLock(func() error {
		if !s.QuestionCard.IsRevealed {
			return errors.New("question_not_revealed")
		}
		if player.Answer.IsLocked {
			return nil
		}
		if len(player.Answer.AnswerCards) == s.QuestionCard.Pick {
			player.Answer.IsLocked = true
			return nil
		}
		return errors.New("not_enough_picks")
	})
}

func (s *Game) RevealNextAnswer() {
	s.withLock(func() {
		for _, player := range s.GetPlayers() {
			if s.GetCurrentPlayer() == player {
				continue
			}
			if player.Answer.IsRevealed {
				continue
			}
			player.Answer.IsRevealed = true
			break
		}
	})
}

func (s *Game) StageAnswer(index int) {
	s.withLock(func() {
		var answers []*Answer
		for _, player := range s.GetPlayers() {
			if player.Answer.IsLocked {
				answers = append(answers, &player.Answer)
			}
		}
		if index < 0 || index >= len(answers) {
			return
		}
		s.StagedAnswer = answers[index]
	})
}

func (s *Game) LockStagedAnswer() {
	s.withLock(func() {
		if s.StagedAnswer != nil {
			s.StagedAnswer.Player.incrementCount()
		}
	})
}
