package board

import (
	"fmt"

	"github.com/ascii-arcade/cards-against-humanity/deck"
	"github.com/charmbracelet/lipgloss"
)

type answerCardComponent struct {
	card  *deck.AnswerCard
	style lipgloss.Style
}

func newAnswerCardComponent(model *Model, card *deck.AnswerCard) *answerCardComponent {
	return &answerCardComponent{
		card:  card,
		style: model.style,
	}
}

func (c *answerCardComponent) render(index int) string {
	style := c.style.
		Border(lipgloss.RoundedBorder()).
		Width(45).
		Padding(1)

	content := fmt.Sprintf("[%d] %s\n", index, c.card.Text)

	return style.Render(content)
}
