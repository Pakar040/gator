package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects 1 arg: username")
	}
	s.config.SetUser(cmd.args[0])
	fmt.Println("User has been set")
	return nil
}
