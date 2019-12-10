package morze

import (
	"strings"
	"testing"
)

func TestDecodeAllAlphabet(t *testing.T) {
	var text, actual, expected string
	for k, v := range alphabet {
		text += v + LetterSpace
		expected += k
	}

	text = strings.TrimSuffix(text, LetterSpace)

	actual = Decode(text)

	if actual != expected {
		t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
	}
}

func TestDecodeCustomText(t *testing.T) {
	var expected = "i am   20"
	var text = ". ." +
		WordSpace +
		". ---" + LetterSpace + "--- ---" +
		WordSpace + WordSpace + WordSpace +
		". . --- --- ---" + LetterSpace + "--- --- --- --- ---"
	var actual = Decode(text)

	if actual != expected {
		t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
	}
}

func TestConvertSequenceToSymbol(t *testing.T) {
	var actual, expected string

	for k, v := range alphabet {
		actual = convertSequenceToSymbol(v)
		expected = k

		if actual != expected {
			t.Errorf("Not equal\n Expected: %v\n Actual: %v", expected, actual)
		}
	}
}
