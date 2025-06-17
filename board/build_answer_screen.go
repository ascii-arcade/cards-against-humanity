package board

import (
	"strconv"

	"github.com/ascii-arcade/cards-against-humanity/colors"
	"github.com/ascii-arcade/cards-against-humanity/keys"
	"github.com/ascii-arcade/cards-against-humanity/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type buildAnswerScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newBuildAnswerScreen() *buildAnswerScreen {
	return &buildAnswerScreen{
		model: m,
		style: m.style,
	}
}

func (s *buildAnswerScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *buildAnswerScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	s.model.clearError()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		switch {
		case s.model.Player.Answer.IsLocked:
			s.model.setError("turn_is_locked")
			return s.model, nil
		case keys.GameAddAnswer.TriggeredBy(msg.String()):
			index, _ := strconv.Atoi(msg.String())
			s.model.Game.AddAnswerCard(s.model.Player, index)
		case keys.GameUndo.TriggeredBy(msg.String()):
			s.model.Game.RemoveAnswerCard(s.model.Player)
		case keys.GameLockAnswer.TriggeredBy(msg.String()):
			err := s.model.Game.LockAnswer(s.model.Player)
			if err != nil {
				s.model.setError(err.Error())
			}
		}
	}

	return s.model, nil
}

func (s *buildAnswerScreen) View() string {
	errorMessage := ""
	if s.model.errorCode != "" {
		errorMessage = s.style.
			Foreground(colors.Error).
			Render("\n" + s.model.lang().Get("error", s.model.errorCode) + "\n")
	}

	return s.model.layoutStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.style.Render(newQuestionCardComponent(s.model, &s.model.Game.QuestionCard).renderForPlayer(s.model.Player.Answer.AnswerCards)),
			errorMessage,
			s.model.contentStyle().Render(s.cards()),
			s.style.Render(newPlayersComponent(s.model).render()),
		),
	)
}

func (s *buildAnswerScreen) cards() string {

	var renderedCards []string
	for i, card := range s.model.Player.Hand {
		renderedCards = append(renderedCards, newAnswerCardComponent(s.model, &card).render(i))
	}
	const maxRows = 5
	const cardsPerRow = 2
	var rows [][]string
	for i := 0; i < len(renderedCards); i += cardsPerRow {
		end := min(i+cardsPerRow, len(renderedCards))
		rows = append(rows, renderedCards[i:end])
		if len(rows) >= maxRows {
			break
		}
	}
	var rowViews []string
	for _, row := range rows {
		rowViews = append(rowViews, lipgloss.JoinHorizontal(lipgloss.Top, row...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rowViews...)
}
