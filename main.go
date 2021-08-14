package main

import (
	"log"
	"os"

	"github.com/1gkx/user-manager/internal/cmd"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
