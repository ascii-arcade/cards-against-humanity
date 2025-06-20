package board

import (
	"strconv"
	"strings"

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
		players: model.game.GetPlayers(),
		style:   model.style,
	}
}

func (c *playersComponent) render() string {
	players := make([]string, 0)
	style := c.style.Width(c.model.width - 10).Align(lipgloss.Center)

	for _, player := range c.players {
		playerStyle := style
		var content strings.Builder

		if c.model.game.GetCurrentPlayer() == player {
			playerStyle = playerStyle.Bold(true)
		} else if !player.Answer.IsLocked {
			content.WriteString("")
		}
		if c.model.Player == player {
			playerStyle = playerStyle.Italic(true)
		}

		content.WriteString(player.Name)
		content.WriteString(": ")
		content.WriteString(strconv.Itoa(player.Points))

		players = append(players, playerStyle.Width(style.GetWidth()/5-2).Render(content.String()))
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
