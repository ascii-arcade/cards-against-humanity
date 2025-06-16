package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type czarScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newCzarScreen() *czarScreen {
	return &czarScreen{
		model: m,
		style: m.style,
	}
}

func (s *czarScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *czarScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil
	case tea.KeyMsg:
		switch {
		case keys.GameEndTurn.TriggeredBy(msg.String()):
			s.model.Game.NextTurn()
		case keys.GameRevealQuestion.TriggeredBy(msg.String()):
			if !s.model.Game.QuestionCard.IsRevealed {
				s.model.Game.RevealQuestionCard()
			}
		}
	}

	return s.model, nil
}

func (s *czarScreen) View() string {
	var content strings.Builder

	if !s.model.Game.QuestionCard.IsRevealed {
		content.WriteString(s.style.Render(s.model.lang().Get("board", "czar_card_not_revealed")) + "\n")
	}
	content.WriteString(s.style.Render(s.model.Game.QuestionCard.Text))

	content.WriteString("\n\n")
	for _, player := range s.model.Game.GetPlayers() {
		if s.model.Player == player {
			continue
		}
		var nameString string
		if player.Answer.IsLocked {
			nameString = "X " + player.Name
		} else {
			nameString = "  " + player.Name
		}
		content.WriteString(nameString)
	}

	return content.String() +
		"\n\n" + s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))
}
