package generaterandom

import (
	"fmt"
	"math/rand/v2"

	"github.com/ascii-arcade/cards-against-humanity/language"
)

func Name(lang *language.Language) string {
	a := lang.UsernameFirstWords[rand.IntN(len(lang.UsernameFirstWords))]
	b := lang.UsernameSecondWords[rand.IntN(len(lang.UsernameSecondWords))]

	return fmt.Sprintf("%s %s", a, b)
}
