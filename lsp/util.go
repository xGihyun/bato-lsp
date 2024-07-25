package lsp

import (
	"strings"
	"unicode"
)

func lineRange(line, start, end uint) Range {
	return Range{
		Start: Position{
			Line:      line,
			Character: start,
		},
		End: Position{
			Line:      line,
			Character: end,
		},
	}
}

func getWordUnderCursor(document string, position Position) string {
	if document == "" {
		return ""
	}

	lines := strings.Split(document, "\n")

	if position.Line >= uint(len(lines)) {
		return ""
	}

	line := lines[position.Line]
	if position.Character >= uint(len(line)) {
		return ""
	}

	start := position.Character
	for start > 0 && isValidChar(rune(line[start-1])) {
		start--
	}

	end := position.Character
	for end < uint(len(line)) && isValidChar(rune(line[end])) {
		end++
	}

	return line[start:end]
}

func isValidChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}
