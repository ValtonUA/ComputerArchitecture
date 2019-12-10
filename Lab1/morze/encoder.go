package morze

import "strings"

// Encode converts plain text to morze sequence
func Encode(text string) string {
	var result string

	text = strings.ToLower(text)

	for i := range text {
		if text[i] == SpaceCode {
			result += WordSpace
		} else if i+1 < len(text) && text[i+1] != SpaceCode {
			result += convertSymbolToSequence(string(text[i])) + LetterSpace
		} else {
			result += convertSymbolToSequence(string(text[i]))
		}

	}

	return result
}

func convertSymbolToSequence(symbol string) string {
	if val, keyExist := alphabet[symbol]; keyExist {
		return val
	}

	return UnexpectedSymbol
}
