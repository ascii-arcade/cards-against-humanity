package board

import (
	"fmt"

	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/deck"
	"github.com/charmbracelet/lipgloss"
)

type answerCardComponent struct {
	card  *deck.AnswerCard
	model *Model
	style lipgloss.Style
}

func newAnswerCardComponent(model *Model, card *deck.AnswerCard) *answerCardComponent {
	return &answerCardComponent{
		card:  card,
		model: model,
		style: model.style,
	}
}

func (c *answerCardComponent) render(index int) string {
	style := c.style.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colors.AnswerBorder).
		Width(c.model.contentWidth()/2 - 5).
		Height(3).
		Padding(1)

	content := fmt.Sprintf("[%d] %s", index, c.card.Text)

	return style.Render(content)
}
