package games

import (
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/deck"
)

type Hand []deck.AnswerCard

func (h *Hand) String() string {
	var content strings.Builder
	content.WriteString("Your Cards:\n\n")
	for _, card := range *h {
		content.WriteString("* " + card.String() + "\n")
	}
	return content.String()
}
