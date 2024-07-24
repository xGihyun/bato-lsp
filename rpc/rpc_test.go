package rpc_test

import (
	"bato-lsp/rpc"
	"fmt"
	"testing"
)

type EncodeTest struct {
	Method string `json:"method"`
}

var testContent = "{\"method\":\"hello\"}"

func TestEncodeMessage(t *testing.T) {
	expected := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(testContent), testContent)
	actual := rpc.EncodeMessage(EncodeTest{Method: "hello"})

	if expected != actual {
		t.Fatalf("Expected: %s, Got %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	message := fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(testContent), testContent)
	method, content, err := rpc.DecodeMessage([]byte(message))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLength != len(testContent) {
		t.Fatalf("Expected: %d, Got %d", len(testContent), contentLength)
	}

	if method != "hello" {
		t.Fatalf("Expected: %s, Got %s", "hello", method)
	}
}
