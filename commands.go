package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	funcHandler, exists := c.registry[cmd.name]
	if !exists {
		return fmt.Errorf("This command is not registered")
	}
	return funcHandler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}
