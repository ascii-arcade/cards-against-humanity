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

func (c *questionCardComponent) renderForBuild(cards []deck.AnswerCard) string {
	if !c.card.IsRevealed {
		return c.cardStyle().Render("")
	}

	return c.cardStyle().Render(c.String(cards))
}

func (c *questionCardComponent) renderForReveal() string {
	content := c.card.Text
	if !strings.Contains(content, "_") {
		content += "\n\n_"
	}
	if !c.card.IsRevealed {
		content += "\n\n" + c.model.lang().Get("board", "czar_question_card_not_revealed")
	}
	if c.model.game.StagedAnswer != nil {
		content = c.String(c.model.game.StagedAnswer.AnswerCards)
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

func (c *questionCardComponent) String(cards []deck.AnswerCard) string {
	if len(cards) == 0 {
		return c.card.Text
	}

	args := make([]any, len(cards))
	content := c.card.Text
	style := c.style

	if !strings.Contains(content, "_") {
		content += "\n\n_"
	}
	if count := len(cards); count > 0 {
		content = strings.Replace(content, "%", "%%", count)
		content = strings.Replace(content, "_", "%s", count)
	}
	for i, card := range cards {
		args[i] = style.Bold(true).Render(card.Text)
	}
	return fmt.Sprintf(content, args...)
}
