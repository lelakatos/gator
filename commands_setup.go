package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.cmds[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.cmds[cmd.name]
	if !exists {
		return errors.New("function not yet added")
	}

	err := f(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
