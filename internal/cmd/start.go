package cmd

import (
	"fmt"

	"github.com/1gkx/user-manager/internal/store"
)

func Start() error {

	store.Init()
	fmt.Println("Start programm")

	return nil
}
