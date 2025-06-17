package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/deck"
	"github.com/charmbracelet/lipgloss"
)

type questionCardComponent struct {
	card  *deck.QuestionCard
	model *Model
	style lipgloss.Style
}

func newQuestionCardComponent(model *Model, card *deck.QuestionCard) *questionCardComponent {
	return &questionCardComponent{
		card:  card,
		model: model,
		style: model.style,
	}
}

func (c *questionCardComponent) renderForPlayer(cards []deck.AnswerCard) string {
	if !c.card.IsRevealed {
		return c.cardStyle().Render("")
	}

	args := make([]any, len(cards))
	format := c.card.Text
	style := c.style

	if count := len(cards); count > 0 {
		format = strings.Replace(format, "_", "%s", count)
	}
	for i, card := range cards {
		args[i] = style.Bold(true).Render(card.Text)
	}
	return c.cardStyle().Render(fmt.Sprintf(format, args...))
}

func (c *questionCardComponent) renderForCzar() string {
	content := c.card.Text
	if !c.card.IsRevealed {
		content += "\n\n" + c.model.lang().Get("board", "czar_card_not_revealed")
	}

	return c.cardStyle().Render(content)
}

func (c *questionCardComponent) cardStyle() lipgloss.Style {
	return c.style.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colors.QuestionBorder).
		Width(c.model.contentWidth() / 3 * 2).
		Height(3).
		Padding(1)
}
