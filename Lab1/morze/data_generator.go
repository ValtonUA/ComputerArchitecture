package morze

import (
	"math/rand"
	"strings"
)

// GetRandomPlainText generates a random text
// using english alphabet symbols
func GetRandomPlainText(length int) string {
	return generateData(length, false)
}

// GetRandomMorzeSequence generates a random sequence
// using Morze alphabet
func GetRandomMorzeSequence(length int) string {
	return generateData(length, true)
}

func generateData(length int, isMorzeSequence bool) string {
	var result string
	var idx int
	var arrayLen = len(alphabet)
	var array = make([]string, arrayLen, arrayLen)

	idx = 0
	for k, v := range alphabet {
		if isMorzeSequence {
			array[idx] = v
		} else {
			array[idx] = k
		}

		idx++
	}

	for i := 0; i < length; i++ {
		idx = rand.Intn(arrayLen + 1)
		if isMorzeSequence {
			if idx == arrayLen {
				result = strings.TrimSuffix(result, LetterSpace) + WordSpace
			} else {
				result += array[idx] + LetterSpace
			}
		} else {
			if idx == arrayLen {
				result += " "
			} else {
				result += array[idx]
			}
		}
	}

	return result
}
