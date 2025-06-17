package board

import (
	"strconv"

	"github.com/ascii-arcade/cards-against-humanity/games"
	"github.com/charmbracelet/lipgloss"
)

type playersComponent struct {
	model   *Model
	players []*games.Player
	style   lipgloss.Style
}

func newPlayersComponent(model *Model) *playersComponent {
	return &playersComponent{
		model:   model,
		players: model.Game.GetPlayers(),
		style:   model.style,
	}
}

func (c *playersComponent) render() string {
	players := make([]string, 0)
	style := c.style.Width(c.model.width - 10).Align(lipgloss.Center)

	for _, player := range c.players {
		content := player.Name + ": " + strconv.Itoa(player.Points)
		playerStyle := style

		if c.model.Player == player {
			playerStyle = playerStyle.Italic(true)
		}
		if c.model.Game.GetCurrentPlayer() == player {
			playerStyle = playerStyle.Bold(true)
		}

		players = append(players, playerStyle.Width(style.GetWidth()/5-2).Render(content))
	}

	var playerRows []string
	if len(players) <= 5 {
		playerRows = append(
			playerRows,
			lipgloss.JoinHorizontal(lipgloss.Left, players...),
		)
	} else {
		n := (len(players) + 1) / 2
		playerRows = append(
			playerRows,
			lipgloss.JoinHorizontal(lipgloss.Left, players[:n]...),
			lipgloss.JoinHorizontal(lipgloss.Left, players[n:]...),
		)
	}

	return style.
		MarginTop(2).
		Render(lipgloss.JoinVertical(lipgloss.Center, playerRows...))
}
