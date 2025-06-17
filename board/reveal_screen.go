package board

import (
	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type revealScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newRevealScreen() *revealScreen {
	return &revealScreen{
		model: m,
		style: m.style,
	}
}

func (s *revealScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *revealScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	s.model.clearError()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil
	case tea.KeyMsg:
		switch {
		case s.model.Game.GetCurrentPlayer() != s.model.Player:
			s.model.setError("not_your_turn")
			return s.model, nil
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

func (s *revealScreen) View() string {
	errorMessage := ""
	if s.model.errorCode != "" {
		errorMessage = s.style.
			Foreground(colors.Error).
			Render("\n" + s.model.lang().Get("error", s.model.errorCode) + "\n")
	}

	return s.model.layoutStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.style.Render(newQuestionCardComponent(s.model, &s.model.Game.QuestionCard).renderForCzar()),
			errorMessage,
			s.style.Render(newPlayersComponent(s.model).render()),
		),
	)
}
