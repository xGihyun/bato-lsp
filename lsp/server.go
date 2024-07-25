package lsp

import (
	"bato-lsp/rpc"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Server struct {
	method   string
	contents []byte
	writer   io.Writer
	logger   *log.Logger
	state    *State
}

type LSPMethod interface {
	initialize()
	open()
	hover()
}

func (s *Server) handleMessage() {
	s.logger.Printf("Received message with method: %s", s.method)

	switch s.method {
	case "initialize":
		s.initialize()
		break
	case "textDocument/didOpen":
		s.open()
		break
	case "textDocument/didChange":
		s.change()
		break
	case "textDocument/hover":
		s.hover()
		break
	case "textDocument/completion":
		s.completion()
		break
	}

}

func (s *Server) initialize() {
	var request InitRequest

	if err := json.Unmarshal(s.contents, &request); err != nil {
		s.logger.Printf("initialize - Failed to parse contents: %s", err)
	}

	client := request.Params.ClientInfo

	s.logger.Printf("Connected to: %s %s", client.Name, client.Version)

	msg := NewInitResponse(request.ID)
	writeResponse(s.writer, msg)
}

func (s *Server) open() {
	var request DidOpenTextDocumentNotification

	if err := json.Unmarshal(s.contents, &request); err != nil {
		s.logger.Printf("open - Failed to parse contents: %s", err)
	}

	document := request.Params.TextDocument

	s.logger.Printf("Opened: %s", document.URI)

	diagnostics := s.state.OpenDocument(document.URI, document.Text)
	notification := PublishDiagnosticsNotification{
		Notification: Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		Params: PublishDiagnosticsParams{
			URI:         document.URI,
			Diagnostics: diagnostics,
		},
	}

	writeResponse(s.writer, notification)
}

func (s *Server) change() {
	var request DidChangeNotification

	if err := json.Unmarshal(s.contents, &request); err != nil {
		s.logger.Printf("change - Failed to parse contents: %s", err)
	}

	document := request.Params.TextDocument

	s.logger.Printf("Changed: %s", document.URI)

	for _, change := range request.Params.ContentChanges {
		diagnostics := s.state.UpdateDocument(document.URI, change.Text)
		notification := PublishDiagnosticsNotification{
			Notification: Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: PublishDiagnosticsParams{
				URI:         document.URI,
				Diagnostics: diagnostics,
			},
		}

		writeResponse(s.writer, notification)
	}
}

func (s *Server) hover() {
	var request HoverRequest

	if err := json.Unmarshal(s.contents, &request); err != nil {
		s.logger.Printf("hover - Failed to parse contents: %s", err)
	}

	response := s.state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
	writeResponse(s.writer, response)
}

func (s *Server) completion() {
	var request CompletionRequest

	if err := json.Unmarshal(s.contents, &request); err != nil {
		s.logger.Printf("completion - Failed to parse contents: %s", err)
	}

	response := s.state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
	writeResponse(s.writer, response)
}

func GetLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Logger file not found.")
	}

	return log.New(logfile, "[batolsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

