package messages

type (
	PlayerUpdate int
)

const (
	Refresh = iota
	BuildAnswerScreen
	RevealScreen
	WinnerScreen
)
