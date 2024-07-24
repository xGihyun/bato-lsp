package analysis

import (
	"bato-lsp/lsp"
	_ "fmt"
	"strings"
	"unicode"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnostics(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnostics(text)
}


// TODO: This should look up the types and stuff
func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]

	word := getWordUnderCursor(document, position)

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: lsp.HoverDocs[strings.Trim(word, " ")],
		},
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	// items := []lsp.CompletionItem{
	// 	{
	// 		Label:         "Honkai Impact 3rd",
	// 		Detail:        "Tuna",
	// 		Documentation: "(´｡• ᵕ •｡`)",
	// 	},
	// }

  items := lsp.MapToCompletionItems(lsp.HoverDocs)

	// Ask static analysis tools to figure out good completions
	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

// TODO: Make this dynamic
func getDiagnostics(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}

	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "Error") {
			idx := strings.Index(line, "Error")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(uint(row), uint(idx), uint(idx+len("Error"))),
				Severity: 1,
				Source:   "Imagination",
				Message:  "Fix this error! (๑˃ᴗ˂)ﻭ",
			})
		}
	}

	return diagnostics
}

func getWordUnderCursor(document string, position lsp.Position) string {
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
