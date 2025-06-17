package board

import (
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

	return s.model.layoutStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.style.Render(newQuestionCardComponent(s.model, &s.model.Game.QuestionCard).renderForCzar()),
			s.style.Render(newPlayersComponent(s.model).render()),
		),
	)
}
