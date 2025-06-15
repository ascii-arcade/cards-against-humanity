package games

import (
	"errors"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

const (
	minimumPlayers = 2
	maximumPlayers = 8
)

func (s *Game) Begin() error {
	return s.withErrLock(func() error {
		if error := s.IsPlayerCountOk(); error != nil {
			return error
		}

		for i, p := range s.players {
			p.SetTurnOrder(i)
		}

		s.AnswerDeck, s.QuestionDeck = deck.NewDecks()
		s.resetPlayerHands()
		s.deal()
		s.CurrentTurnIndex = 0
		s.inProgress = true
		return nil
	})
}

func (s *Game) IsPlayerCountOk() error {
	if len(s.players) > maximumPlayers {
		return errors.New("too_many_players")
	}
	if len(s.players) < minimumPlayers {
		return errors.New("not_enough_players")
	}
	return nil
}

func (s *Game) resetPlayerHands() {
	for _, player := range s.players {
		player.Hand = make(Hand, 0)
	}
}

func (s *Game) deal() {
	s.QuestionCard = s.QuestionDeck[0]
	s.QuestionDeck = s.QuestionDeck[1:]

	for _, player := range s.players {
		for len(player.Hand) < s.Config.HandSize && len(s.AnswerDeck) > 0 {
			player.Hand = append(player.Hand, s.AnswerDeck[0])
			s.AnswerDeck = s.AnswerDeck[1:]
		}
	}
}
