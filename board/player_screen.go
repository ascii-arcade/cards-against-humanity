package board

import (
	"fmt"

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
		}
	}

	return s.model, nil
}

func (s *playerScreen) View() string {
	questionContent := s.model.lang().Get("board", "card_not_revealed")
	if s.model.Game.QuestionCard.IsRevealed {
		questionContent = s.model.Game.QuestionCard.Text
	}

	return questionContent +
		"\n\n" + s.model.Player.Hand.String() +
		"\n\n" + s.style.Render(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))
}
