package games

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

type Hand []deck.AnswerCard

func (h *Hand) String() string {
	var content strings.Builder
	content.WriteString("Your Cards:\n\n")
	for i, card := range *h {
		content.WriteString(fmt.Sprintf("[%d] %s\n", i, card.Text))
	}
	return content.String()
}

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
