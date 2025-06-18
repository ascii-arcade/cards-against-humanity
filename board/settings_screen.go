package board

import (
	"fmt"

	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type settingsScreen struct {
	form  *huh.Form
	model *Model
	style lipgloss.Style
}

func (m *Model) newSettingsScreen() *settingsScreen {
	return &settingsScreen{
		form:  newSettingsForm(),
		model: m,
		style: m.style,
	}
}

func (s *settingsScreen) Init() tea.Cmd {
	return s.form.Init()
}

func (s *settingsScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		switch {
		case keys.PreviousScreen.TriggeredBy(msg.String()):
			s.model.Player.ActiveScreenCode = screen.BoardLobby
			return s.model, nil
		}
	}

	form, cmd := s.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		s.form = f
	}

	return s.model, cmd
}

func (s *settingsScreen) View() string {
	if s.form.State == huh.StateCompleted {
		level := s.form.GetInt("points")
		return fmt.Sprintf("You selected: %d", level)
	}
	return s.form.View()
}

func newSettingsForm() *huh.Form {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Key("points").
				Options(huh.NewOptions(3, 4, 5, 6, 7, 8, 9, 10)...).
				Title("Points to Win"),
		),
	)
	return form
}
