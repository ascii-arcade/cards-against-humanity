package games

import (
	"errors"
	"math/rand/v2"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

func (s *Game) Begin() error {
	return s.withErrLock(func() error {
		if error := s.IsPlayerCountOk(); error != nil {
			return error
		}

		rand.Shuffle(len(s.players), func(i, j int) {
			s.players[i], s.players[j] = s.players[j], s.players[i]
		})
		for i, p := range s.players {
			p.SetTurnOrder(i)
		}

		s.AnswerDeck, s.QuestionDeck = deck.NewDecks()
		s.resetPlayerHands()
		s.deal()
		s.LockedAnswers = make([]*Answer, 0)
		s.CurrentTurnIndex = 0
		s.inProgress = true
		s.setPlayerScreens()
		return nil
	})
}

func (s *Game) IsPlayerCountOk() error {
	if len(s.players) > s.Settings.MaxPlayers {
		return errors.New("too_many_players")
	}
	if len(s.players) < s.Settings.MinPlayers {
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
	s.StagedAnswer = nil

	for _, player := range s.players {
		player.Answer = Answer{Player: player}
		for len(player.Hand) < s.Settings.HandSize && len(s.AnswerDeck) > 0 {
			player.Hand.add(s.AnswerDeck[0])
			s.AnswerDeck = s.AnswerDeck[1:]
		}
	}
}
