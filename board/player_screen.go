package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/colors"
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
	s.model.clearError()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		switch {
		case s.model.Player.Answer.IsLocked:
			s.model.setError("turn_is_locked")
			return s.model, nil
		case keys.GameAddAnswer.TriggeredBy(msg.String()):
			index, _ := strconv.Atoi(msg.String())
			s.model.Game.AddAnswerCard(s.model.Player, index)
		case keys.GameUndo.TriggeredBy(msg.String()):
			s.model.Game.RemoveAnswerCard(s.model.Player)
		case keys.GameLockAnswer.TriggeredBy(msg.String()):
			err := s.model.Game.LockAnswer(s.model.Player)
			if err != nil {
				s.model.setError(err.Error())
			}
		}
	}

	return s.model, nil
}

func (s *playerScreen) View() string {
	questionContent := s.model.lang().Get("board", "card_not_revealed")
	if s.model.Game.QuestionCard.IsRevealed {
		questionContent = s.model.Game.QuestionCard.Text
	}

	if s.model.errorCode != "" {
		fmt.Println("Error code:", s.model.errorCode)
		questionContent += s.style.
			Foreground(colors.Error).
			Render("\n" + s.model.lang().Get("error", s.model.errorCode) + "\n")
	}

	var answerContent strings.Builder
	if s.model.Player.Answer.IsLocked {
		answerContent.WriteString(s.style.
			Render(s.model.lang().Get("board", "answer_locked") + "\n"))
	} else {
		answerContent.WriteString(s.style.
			Render(s.model.lang().Get("board", "answer_not_locked") + "\n"))
	}
	for i, answerCard := range s.model.Player.Answer.AnswerCards {
		answerContent.WriteString(fmt.Sprintf("[%d] %s\n", i, answerCard.Text))
	}

	return questionContent +
		"\n\n" + answerContent.String() +
		"\n\n" + s.model.Player.Hand.String() +
		"\n\n" + s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))
}
