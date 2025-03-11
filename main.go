package main

import (
	"log"
	"os"

	"github.com/lelakatos/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{&cfg}
	cmds := commands{cmds: make(map[string]func(*state, command) error)}

	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering login command: %v", err)
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("not enough arguments passed in")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("error running the command: %v", err)
	}
}
