package deck

import (
	"strings"

	"github.com/charmbracelet/glamour"
)

const style = `
{
  "document": {
    "margin": 0,
    "padding": 0
  },
  "paragraph": {
    "margin": 0,
    "padding": 0
  }
}
`

var renderer, _ = glamour.NewTermRenderer(
	glamour.WithStylesFromJSONBytes([]byte(style)),
)

type AnswerCard struct {
	ID   int
	Text string `json:"text"`
}

type QuestionCard struct {
	ID   int
	Text string `json:"text"`
	Pick int    `json:"pick"`
}

func (card *AnswerCard) String() string {
	return markdownRender(card.Text)
}

func (card *QuestionCard) String() string {
	return markdownRender(card.Text)
}

func markdownRender(text string) string {
	rendered, err := renderer.Render(text)
	if err != nil {
		return text
	}
	return strings.TrimSpace(rendered)
}
