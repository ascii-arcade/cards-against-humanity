package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type playerScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newPlayerScreen() *playerScreen {
	return &playerScreen{
		model: m,
		style: m.style,
	}
}

func (s *playerScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *playerScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		switch {
		case keys.GameAddAnswer.TriggeredBy(msg.String()):
			index, _ := strconv.Atoi(msg.String())
			s.model.Game.AddAnswerCard(s.model.Player, index)
		case keys.GameUndo.TriggeredBy(msg.String()):
			s.model.Game.RemoveAnswerCard(s.model.Player)
		}
	}

	return s.model, nil
}

func (s *playerScreen) View() string {
	questionContent := s.model.lang().Get("board", "card_not_revealed")
	if s.model.Game.QuestionCard.IsRevealed {
		questionContent = s.model.Game.QuestionCard.Text
	}

	var answerContent strings.Builder
	for i, answerCard := range s.model.Player.Answer.AnswerCards {
		answerContent.WriteString(fmt.Sprintf("[%d] %s\n", i, answerCard.Text))
	}

	return questionContent +
		"\n\n" + answerContent.String() +
		"\n\n" + s.model.Player.Hand.String() +
		"\n\n" + s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))
}
