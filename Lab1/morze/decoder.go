package morze

import "strings"

// Decode converts morze sequence to plain text
func Decode(text string) string {
	var result string

	var words = getLexems(text, len(WordSpace))

	for wordIdx := range words {
		var letters = getLexems(words[wordIdx], len(LetterSpace))

		for letterIdx := range letters {
			result += convertSequenceToSymbol(letters[letterIdx])
		}

		result += " "
	}

	return strings.TrimSuffix(result, " ")
}

func getLexems(text string, spaceLength int) []string {
	var spaces = strings.Repeat(" ", spaceLength)
	var capacity = strings.Count(text, spaces) + 1 // Number of lexems
	var lexems = make([]string, 0, capacity)
	var lexem string

	for i := 0; i < len(text); i++ {
		lexem += string(text[i])

		if strings.HasSuffix(lexem, spaces) || i+1 == len(text) {
			lexems = append(lexems, strings.TrimSuffix(lexem, spaces))
			lexem = ""
		}
	}

	return lexems
}

func convertSequenceToSymbol(sequence string) string {
	for key, val := range alphabet {
		if val == sequence {
			return key
		}
	}

	return UnexpectedSymbol
}
