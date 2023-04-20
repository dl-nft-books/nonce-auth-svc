package main

import (
	"os"

	"github.com/dl-nft-books/nonce-auth-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
