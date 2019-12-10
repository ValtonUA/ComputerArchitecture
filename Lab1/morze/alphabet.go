package morze

const (
	// LetterPartSpace is number of spaces between parts of the same letter
	LetterPartSpace = " "
	// LetterSpace is number of spaces between letters
	LetterSpace = "   "
	// WordSpace is number of spaces between words
	WordSpace = "       "
	// SpaceCode is ASCII code of ' '
	SpaceCode = 32
	// UnexpectedSymbol replaces invalid symbols
	UnexpectedSymbol = "?"
)

var alphabet = map[string]string{
	"a": ". ---",
	"b": "--- . . .",
	"c": "--- . --- .",
	"d": "--- . .",
	"e": ".",
	"f": ". . --- .",
	"g": "--- --- .",
	"h": ". . . .",
	"i": ". .",
	"j": ". --- --- ---",
	"k": "--- . ---",
	"l": ". --- . .",
	"m": "--- ---",
	"n": "--- .",
	"o": "--- --- ---",
	"p": ". --- --- .",
	"q": "--- --- . ---",
	"r": ". --- .",
	"s": ". . .",
	"t": "---",
	"u": ". . ---",
	"v": ". . . ---",
	"w": ". --- ---",
	"x": "--- . . ---",
	"y": "--- . --- ---",
	"z": "--- --- . .",
	"1": ". --- --- --- ---",
	"2": ". . --- --- ---",
	"3": ". . . --- ---",
	"4": ". . . . ---",
	"5": ". . . . .",
	"6": "--- . . . .",
	"7": "--- --- . . .",
	"8": "--- --- --- . .",
	"9": "--- --- --- --- .",
	"0": "--- --- --- --- ---"}
