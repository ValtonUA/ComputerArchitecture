package morze

import (
	"strings"
	"testing"
)

func TestEncodeAllAlphabet(t *testing.T) {
	var text, actual, expected string
	for k, v := range alphabet {
		text += k
		expected += v + LetterSpace
	}

	expected = strings.TrimSuffix(expected, LetterSpace)

	actual = Encode(text)

	if actual != expected {
		t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
	}
}

func TestEncodeCustomText(t *testing.T) {
	const Text = "I am 20"
	var expected = ". ." +
		WordSpace +
		". ---" + LetterSpace + "--- ---" +
		WordSpace +
		". . --- --- ---" + LetterSpace + "--- --- --- --- ---"
	var actual = Encode(Text)

	if actual != expected {
		t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
	}
}

func TestConvertSymbolToSequenceAllAlphabet(t *testing.T) {
	var actual, expected string

	for k, v := range alphabet {
		actual = convertSymbolToSequence(k)
		expected = v

		if actual != expected {
			t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
		}
	}
}
