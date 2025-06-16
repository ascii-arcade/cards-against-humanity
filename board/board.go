package board

import (
	"fmt"
	"strings"

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
	style     lipgloss.Style
	errorCode string

	Player *games.Player
	Game   *games.Game
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

func (m *Model) lang() *language.Language {
	return m.Player.LanguagePreference.Lang
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.RefreshBoard:
		return m, waitForRefreshSignal(m.Player.UpdateChan)

	case tea.KeyMsg:
		switch {
		case keys.ExitApplication.TriggeredBy(msg.String()):
			m.Game.RemovePlayer(m.Player)
			return m, tea.Quit
		}
	}

	screenModel, cmd := m.activeScreen().Update(msg)
	return screenModel.(*Model), cmd
}

func (m Model) View() string {
	if m.width < config.MinimumWidth {
		return m.lang().Get("error", "window_too_narrow")
	}
	if m.height < config.MinimumHeight {
		return m.lang().Get("error", "window_too_short")
	}

	disconnectedPlayers := m.Game.GetDisconnectedPlayers()
	if len(disconnectedPlayers) > 0 {
		var names []string
		for _, p := range disconnectedPlayers {
			names = append(names, p.Name)
		}
		return m.style.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.style.Align(lipgloss.Center).MarginBottom(2).Render(m.Game.Code),
				fmt.Sprintf(m.lang().Get("board", "disconnected_player"), strings.Join(names, ", ")),
				m.style.Render(fmt.Sprintf(m.lang().Get("global", "quit"), keys.ExitApplication.String(m.style))),
			),
		)
	}

	return m.activeScreen().View()
}

func (m *Model) activeScreen() screen.Screen {
	if !m.Game.InProgress() {
		return m.newLobbyScreen()
	}

	switch {
	case m.Game.GetCurrentPlayer() == m.Player:
		return m.newCzarScreen()
	default:
		return m.newPlayerScreen()
	}
}

func waitForRefreshSignal(ch chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshBoard(<-ch)
	}
}

func (m *Model) setError(err string) {
	m.errorCode = err
}

func (m *Model) clearError() {
	m.errorCode = ""
}

func (m *Model) layoutStyle() lipgloss.Style {
	return m.style.
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center)
}

func (m *Model) contentStyle() lipgloss.Style {
	return m.style.
		Width(m.width/2).
		Height(m.height-10).
		Align(lipgloss.Left, lipgloss.Top)
}
