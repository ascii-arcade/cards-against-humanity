package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/config"
	"github.com/ascii-arcade/cards-against-humanity/games"
	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/language"
	"github.com/ascii-arcade/cards-against-humanity/messages"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width     int
	height    int
	screen    screen.Screen
	style     lipgloss.Style
	errorCode string

	Player *games.Player
	game   *games.Game
}

func NewModel(width, height int, style lipgloss.Style, player *games.Player) Model {
	m := Model{
		width:  width,
		height: height,
		style:  style,
		Player: player,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return waitForRefreshSignal(m.Player.UpdateChan)
}

func (m *Model) SetGame(game *games.Game) {
	m.screen = m.newLobbyScreen()
	m.game = game
}

func (m *Model) lang() *language.Language {
	return m.Player.LanguagePreference.Lang
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.PlayerUpdate:
		return m.handlePlayerUpdate(int(msg))

	case tea.KeyMsg:
		switch {
		case keys.ExitApplication.TriggeredBy(msg.String()):
			m.game.RemovePlayer(m.Player)
			return m, tea.Quit
		}
	}

	screenModel, cmd := m.screen.Update(msg)
	return screenModel.(*Model), cmd
}

func (m Model) View() string {
	if m.width < config.MinimumWidth {
		return m.lang().Get("error", "window_too_narrow")
	}
	if m.height < config.MinimumHeight {
		return m.lang().Get("error", "window_too_short")
	}

	disconnectedPlayers := m.game.GetDisconnectedPlayers()
	if len(disconnectedPlayers) > 0 {
		var names []string
		for _, p := range disconnectedPlayers {
			names = append(names, p.Name)
		}
		return m.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.style.Align(lipgloss.Center).MarginBottom(2).Render(m.game.Code),
				fmt.Sprintf(m.lang().Get("board", "disconnected_player"), strings.Join(names, ", ")),
				m.style.Render(fmt.Sprintf(m.lang().Get("global", "quit"), keys.ExitApplication.String(m.style))),
			),
		)
	}

	return m.screen.View()
}

func waitForRefreshSignal(ch chan int) tea.Cmd {
	return func() tea.Msg {
		return messages.PlayerUpdate(<-ch)
	}
}

func (m *Model) setError(err string) {
	m.errorCode = err
}

func (m *Model) clearError() {
	m.errorCode = ""
}

func (m *Model) renderedError() string {
	errorMessage := ""
	if m.errorCode != "" {
		errorMessage = m.style.
			Foreground(colors.Error).
			Render("\n" + m.lang().Get("error", m.errorCode) + "\n")
	}

	return errorMessage
}

func (m *Model) layoutStyle() lipgloss.Style {
	return m.style.
		Width(m.width).
		Height(m.height).
		PaddingTop(3).
		Align(lipgloss.Center)
}

func (m *Model) contentStyle() lipgloss.Style {
	return m.style.
		Width(m.contentWidth()).
		Align(lipgloss.Left, lipgloss.Top)
}

func (m *Model) contentWidth() int {
	return max(config.MinimumWidth-10, m.width-30)
}

func (m *Model) handlePlayerUpdate(msg int) (tea.Model, tea.Cmd) {
	switch msg {
	case messages.BuildAnswerScreen:
		m.screen = m.newBuildAnswerScreen()
	case messages.RevealScreen:
		m.screen = m.newRevealScreen()
	case messages.SettingsScreen:
		m.screen = m.newSettingsScreen()
	case messages.WinnerScreen:
		m.screen = m.newWinnerScreen()
	}
	return m, waitForRefreshSignal(m.Player.UpdateChan)
}
