package board

import (
	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type settingsScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newSettingsScreen() *settingsScreen {
	return &settingsScreen{
		model: m,
		style: m.style,
	}
}

func (s *settingsScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *settingsScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case keys.PreviousScreen.TriggeredBy(msg.String()):
			s.model.screen = s.model.newLobbyScreen()
			return s.model, nil
		}
	}

	return s.model, nil
}

func (s *settingsScreen) View() string {
	return "settings!"
}
