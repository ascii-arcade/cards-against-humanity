package deck

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand/v2"
	"slices"
)

type Pack struct {
	AnswerCards   []AnswerCard   `json:"white"`
	QuestionCards []QuestionCard `json:"black"`
}

var allAnswerCards []AnswerCard
var allQuestionCards []QuestionCard

//go:embed CAH.json
var jsonData []byte

func init() {
	var packs []Pack
	if err := json.Unmarshal(jsonData, &packs); err != nil {
		log.Fatal(err)
	}

	id := 1

	for _, pack := range packs {
		for _, card := range pack.QuestionCards {
			card.ID = id
			id++
			allQuestionCards = append(allQuestionCards, card)
		}
		for _, card := range pack.AnswerCards {
			card.ID = id
			id++
			allAnswerCards = append(allAnswerCards, card)
		}
	}
}

func NewDecks() ([]AnswerCard, []QuestionCard) {
	answerDeck := slices.Clone(allAnswerCards)
	shuffle(answerDeck)

	questionDeck := slices.Clone(allQuestionCards)
	shuffle(questionDeck)

	return answerDeck, questionDeck
}

func shuffle[S any](s []S) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
