package deck

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand/v2"
	"slices"
)

type AnswerCard struct {
	Text string `json:"text"`
}

type QuestionCard struct {
	IsRevealed bool
	Text       string `json:"text"`
	Pick       int    `json:"pick"`
}

type Pack struct {
	AnswerCards   []AnswerCard   `json:"white"`
	QuestionCards []QuestionCard `json:"black"`
}

var allAnswers []AnswerCard
var allQuestions []QuestionCard

//go:embed data/CAH.json
var jsonData []byte

func init() {
	var packs []Pack
	if err := json.Unmarshal(jsonData, &packs); err != nil {
		log.Fatal(err)
	}

	for _, pack := range packs {
		allAnswers = append(allAnswers, pack.AnswerCards...)
		allQuestions = append(allQuestions, pack.QuestionCards...)
	}
}

func NewDecks() ([]AnswerCard, []QuestionCard) {
	answerDeck := slices.Clone(allAnswers)
	shuffle(answerDeck)

	questionDeck := slices.Clone(allQuestions)
	shuffle(questionDeck)

	return answerDeck, questionDeck
}

func shuffle[S any](s []S) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
