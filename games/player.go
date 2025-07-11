package games

import (
	"context"

	"github.com/ascii-arcade/cards-against-humanity/language"
	"github.com/charmbracelet/ssh"
)

type Player struct {
	Name      string
	Points    int
	Answer    Answer
	Hand      Hand
	TurnOrder int

	isHost    bool
	connected bool

	UpdateChan         chan int
	LanguagePreference *language.LanguagePreference

	Sess ssh.Session

	onDisconnect []func()
	ctx          context.Context
}

func (p *Player) SetName(name string) *Player {
	p.Name = name
	return p
}

func (p *Player) SetTurnOrder(order int) *Player {
	p.TurnOrder = order
	return p
}

func (p *Player) MakeHost() *Player {
	p.isHost = true
	return p
}

func (p *Player) IsHost() bool {
	return p.isHost
}

func (p *Player) OnDisconnect(fn func()) {
	p.onDisconnect = append(p.onDisconnect, fn)
}

func (p *Player) incrementCount() {
	p.Points++
}

func (p *Player) update(code int) {
	select {
	case p.UpdateChan <- code:
	default:
	}
}
