package games

import (
	"slices"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

type Answer struct {
	AnswerCards    []deck.AnswerCard
	IsLocked       bool
	IsRevealed     bool
	Player         *Player
	WillBeRevealed bool
}

func (a *Answer) add(index int) {
	if index < 0 || index >= len(a.Player.Hand) {
		return
	}
	card := a.Player.Hand[index]
	a.Player.Hand.remove(card)
	a.Player.Answer.AnswerCards = append(a.Player.Answer.AnswerCards, card)
}

func (a *Answer) remove(index int) {
	if index < 0 || index >= len(a.AnswerCards) {
		return
	}
	card := a.AnswerCards[index]
	a.AnswerCards = slices.Delete(a.AnswerCards, index, index+1)
	a.Player.Hand.add(card)
}

func (a *Answer) String() string {
	answerStrings := make([]string, 0)
	for _, card := range a.AnswerCards {
		answerStrings = append(answerStrings, card.Text)
	}

	return strings.Join(answerStrings, ",\n")
}
