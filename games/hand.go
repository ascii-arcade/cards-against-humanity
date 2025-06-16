package games

import (
	"slices"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

type Hand []deck.AnswerCard

func (h *Hand) add(card deck.AnswerCard) {
	*h = append(*h, card)
}

func (h *Hand) remove(card deck.AnswerCard) {
	for i, c := range *h {
		if c.Text == card.Text {
			*h = slices.Delete(*h, i, i+1)
		}
	}
}
