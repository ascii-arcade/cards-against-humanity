package deck

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
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
	rendered, err := glamour.Render(card.Text, "dark")
	if err != nil {
		fmt.Println("error rendering markdown:", err)
		os.Exit(1)
	}
	return rendered
}

func (card *QuestionCard) String() string {
	rendered, err := glamour.Render(card.Text, "dark")
	if err != nil {
		fmt.Println("error rendering markdown:", err)
		os.Exit(1)
	}
	return rendered
}
