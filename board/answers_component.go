package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/games"
	"github.com/charmbracelet/lipgloss"
)

type answersComponent struct {
	answer *games.Answer
	model  *Model
	style  lipgloss.Style
}

func newAnswersComponent(model *Model, answer *games.Answer) *answersComponent {
	return &answersComponent{
		answer: answer,
		model:  model,
		style:  model.style,
	}
}

func (c *answersComponent) render(isAllRevealed bool, willBeRevealed bool, index int) string {
	style := c.style.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colors.AnswerBorder).
		Width(c.model.contentWidth()/2 - 5).
		Height(3).
		Padding(1)

	var content strings.Builder

	if c.model.game.GetCurrentPlayer() == c.model.Player && willBeRevealed {
		content.WriteString("[r]eveal")
	}
	if isAllRevealed {
		content.WriteString(fmt.Sprintf("[%d] ", index))
	}
	if c.answer.IsRevealed {
		content.WriteString(c.answer.String())
	}

	return style.Render(content.String())
}
