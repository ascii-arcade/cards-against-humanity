package board

import (
	"math/rand/v2"
	"strconv"

	"github.com/ascii-arcade/cards-against-humanity/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type revealScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newRevealScreen() *revealScreen {
	return &revealScreen{
		model: m,
		style: m.style,
	}
}

func (s *revealScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	s.model.clearError()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil
	case tea.KeyMsg:
		switch {
		case s.model.game.GetCurrentPlayer() != s.model.Player:
			s.model.setError("not_your_turn")
			return s.model, nil
		case keys.GamePick.TriggeredBy(msg.String()):
			if s.isAllRevealed() {
				index, _ := strconv.Atoi(msg.String())
				s.model.game.StageAnswer(index)
			}
		case keys.GameLock.TriggeredBy(msg.String()):
			if s.model.game.StagedAnswer != nil {
				s.model.game.LockStagedAnswer()
			}
		case keys.GameReveal.TriggeredBy(msg.String()):
			if !s.model.game.QuestionCard.IsRevealed {
				s.model.game.RevealQuestionCard()
				return s.model, nil
			}
			if s.isAllLocked() {
				s.model.game.RevealNextAnswer()
			}
		}
	}

	return s.model, nil
}

func (s *revealScreen) View() string {
	question := newQuestionCardComponent(
		s.model,
		&s.model.game.QuestionCard,
	).renderForReveal()

	return s.model.layoutStyle().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.style.Render(question),
			s.model.renderedError(),
			s.answers(),
			s.style.Render(newPlayersComponent(s.model).render()),
		),
	)
}

func (s *revealScreen) answers() string {
	var answers []*answersComponent
	for _, player := range s.model.game.GetPlayers() {
		if player.Answer.IsLocked {
			answers = append(answers, newAnswersComponent(s.model, &player.Answer))
		}
	}
	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})

	canReveal := s.isAllLocked()
	var answerComponents []string
	for i, answer := range answers {
		willReveal := !answer.answer.IsRevealed && canReveal
		answerComponents = append(answerComponents, answer.render(s.isAllRevealed(), willReveal, i))
		if !answer.answer.IsRevealed && canReveal {
			canReveal = false
		}
	}

	var answerRowsStyled []string
	cols := 2
	for i := 0; i < len(answers); i += cols {
		end := min(i+cols, len(answerComponents))
		row := lipgloss.JoinHorizontal(lipgloss.Top, answerComponents[i:end]...)
		answerRowsStyled = append(answerRowsStyled, row)
	}

	return lipgloss.JoinVertical(lipgloss.Left, answerRowsStyled...)
}

func (s *revealScreen) isAllRevealed() bool {
	for _, player := range s.model.game.GetPlayers() {
		if s.model.game.GetCurrentPlayer() == player {
			continue
		}
		if !player.Answer.IsRevealed {
			return false
		}
	}
	return true
}

func (s *revealScreen) isAllLocked() bool {
	for _, player := range s.model.game.GetPlayers() {
		if s.model.game.GetCurrentPlayer() == player {
			continue
		}
		if !player.Answer.IsLocked {
			return false
		}
	}
	return true
}
