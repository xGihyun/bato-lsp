package main

import (
	"bato-lsp/lsp"
	"bato-lsp/rpc"
	"bufio"
	"os"
)

func main() {
	logger := lsp.GetLogger("/home/gihyun/Documents/Programming/Go/bato-lsp/log.txt")
	logger.Println("LSP starting...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := lsp.NewState()
	writer := os.Stdout

	lsp.MainLoop(scanner, writer, state, logger)

	logger.Println("LSP closed.")
}
