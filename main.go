package main

import (
	"bato-lsp/analysis"
	"bato-lsp/lsp"
	"bato-lsp/rpc"
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/gihyun/Documents/Programming/Go/bato-lsp/log.txt")
	logger.Println("LSP Starting...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)

		if err != nil {
			logger.Printf("ERROR: %s", err)
			continue
		}

		messageHandler := MessageHandler{
			method:   method,
			contents: contents,
			writer:   writer,
			logger:   logger,
			state:    &state,
		}

		messageHandler.handleMessage()
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

// TODO: Put the rest of the code below on a separate file 
type MessageHandler struct {
	method   string
	contents []byte
	writer   io.Writer
	logger   *log.Logger
	state    *analysis.State
}

type LSPMethod interface {
	initialize()
	open()
	hover()
}

func (m *MessageHandler) handleMessage() {
	m.logger.Printf("Received message with method: %s", m.method)

	switch m.method {
	case "initialize":
		m.initialize()
		break
	case "textDocument/didOpen":
		m.open()
		break
	case "textDocument/didChange":
		m.change()
		break
	case "textDocument/hover":
		m.hover()
		break
	case "textDocument/completion":
		m.completion()
		break
	}

}

func (m *MessageHandler) initialize() {
	var request lsp.InitRequest

	if err := json.Unmarshal(m.contents, &request); err != nil {
		m.logger.Printf("initialize - Failed to parse contents: %s", err)
	}

	client := request.Params.ClientInfo

	m.logger.Printf("Connected to: %s %s", client.Name, client.Version)

	msg := lsp.NewInitResponse(request.ID)
	writeResponse(m.writer, msg)
}

func (m *MessageHandler) open() {
	var request lsp.DidOpenTextDocumentNotification

	if err := json.Unmarshal(m.contents, &request); err != nil {
		m.logger.Printf("open - Failed to parse contents: %s", err)
	}

	document := request.Params.TextDocument

	m.logger.Printf("Opened: %s", document.URI)

	diagnostics := m.state.OpenDocument(document.URI, document.Text)
	notification := lsp.PublishDiagnosticsNotification{
		Notification: lsp.Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		Params: lsp.PublishDiagnosticsParams{
			URI:         document.URI,
			Diagnostics: diagnostics,
		},
	}

	writeResponse(m.writer, notification)
}

func (m *MessageHandler) change() {
	var request lsp.DidChangeNotification

	if err := json.Unmarshal(m.contents, &request); err != nil {
		m.logger.Printf("change - Failed to parse contents: %s", err)
	}

	document := request.Params.TextDocument

	m.logger.Printf("Changed: %s", document.URI)

	for _, change := range request.Params.ContentChanges {
		diagnostics := m.state.UpdateDocument(document.URI, change.Text)
		notification := lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: lsp.PublishDiagnosticsParams{
				URI:         document.URI,
				Diagnostics: diagnostics,
			},
		}

		writeResponse(m.writer, notification)
	}
}

func (m *MessageHandler) hover() {
	var request lsp.HoverRequest

	if err := json.Unmarshal(m.contents, &request); err != nil {
		m.logger.Printf("hover - Failed to parse contents: %s", err)
	}

	response := m.state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
	writeResponse(m.writer, response)
}

func (m *MessageHandler) completion() {
	var request lsp.CompletionRequest

	if err := json.Unmarshal(m.contents, &request); err != nil {
		m.logger.Printf("completion - Failed to parse contents: %s", err)
	}

	response := m.state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)
	writeResponse(m.writer, response)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("Logger file not found.")
	}

	return log.New(logfile, "[batolsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
