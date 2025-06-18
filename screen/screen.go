package screen

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	MenuSplash = iota
	MenuTitle
	MenuJoin
)

const (
	BoardLobby = iota
	BoardSettings
	BoardReveal
	BoardBuildAnswer
	BoardWinner
)

type Screen interface {
	Update(tea.Msg) (any, tea.Cmd)
	View() string
}
