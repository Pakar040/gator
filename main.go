package main

import (
	"fmt"
	"os"

	"github.com/Pakar040/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	s := state{config: &cfg}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Command not recognized")
		os.Exit(1)
	}

	if err := cmds.run(&s, command{name: os.Args[1], args: os.Args[2:]}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
