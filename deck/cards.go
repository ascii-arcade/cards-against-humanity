package deck

type AnswerCard struct {
	ID         int
	IsRevealed bool
	Text       string `json:"text"`
}

type QuestionCard struct {
	ID         int
	IsRevealed bool
	Text       string `json:"text"`
	Pick       int    `json:"pick"`
}
