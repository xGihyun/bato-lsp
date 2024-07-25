package lsp

import (
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

func (s *State) OpenDocument(uri, text string) []Diagnostic {
	s.Documents[uri] = text

	return getDiagnostics(text)
}

func (s *State) UpdateDocument(uri, text string) []Diagnostic {
	s.Documents[uri] = text

	return getDiagnostics(text)
}

// TODO: This should look up the types and stuff
func (s *State) Hover(id int, uri string, position Position) HoverResponse {
	document := s.Documents[uri]

	word := getWordUnderCursor(document, position)

	return HoverResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: HoverResult{
			Contents: CompletionMap[strings.Trim(word, " ")],
		},
	}
}

// NOTE: 
// It would be difficult to determine the data type since that would 
// require some magic from converting Ruby -> Go lang types since Bato
// is written in Ruby
func (s *State) TextDocumentCompletion(id int, uri string) CompletionResponse {
	items := MapToCompletionItems(CompletionMap)

	// Ask static analysis tools to figure out good completions
	response := CompletionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}

// TODO: Make this dynamic
func getDiagnostics(text string) []Diagnostic {
	diagnostics := []Diagnostic{}

	// TODO: Show error for syntax errors
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "Error") {
			idx := strings.Index(line, "Error")
			diagnostics = append(diagnostics, Diagnostic{
				Range:    lineRange(uint(row), uint(idx), uint(idx+len("Error"))),
				Severity: 1,
				Source:   "Imagination",
				Message:  "Fix this error! (๑˃ᴗ˂)ﻭ",
			})
		}
	}

	return diagnostics
}
