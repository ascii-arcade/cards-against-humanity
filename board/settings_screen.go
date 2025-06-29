package board

import (
	"strconv"

	"github.com/ascii-arcade/cards-against-humanity/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type settingsScreen struct {
	model     *Model
	style     lipgloss.Style
	form      *huh.Form
	endpoints string
}

func (m *Model) newSettingsScreen() *settingsScreen {
	s := &settingsScreen{
		model:     m,
		style:     m.style,
		endpoints: strconv.Itoa(m.game.Settings.EndPoints),
	}
	s.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("End Points").
				Value(&s.endpoints),
		),
	)
	return s
}

func (s *settingsScreen) Init() tea.Cmd {
	return s.form.Init()
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

	form, cmd := s.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		s.form = f
	}

	if s.form.State == huh.StateCompleted {
		var err error
		s.model.game.Settings.EndPoints, err = strconv.Atoi(s.endpoints)
		if err != nil {
			s.model.setError("Invalid end points value")
			return s.model, cmd
		}

		s.model.screen = s.model.newLobbyScreen()
	}

	return s.model, cmd
}

func (s *settingsScreen) View() string {
	return "Settings:\n\n" + s.form.View()
}
