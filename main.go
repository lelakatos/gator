package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/lelakatos/gator/internal/config"
	"github.com/lelakatos/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error loading the database: %v", err)
	}
	dbQueries := database.New(db)

	currentState := state{&cfg, dbQueries}
	cmds := commands{cmds: make(map[string]func(*state, command) error)}

	err = cmds.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering login command: %v", err)
	}

	err = cmds.register("register", handlerRegister)
	if err != nil {
		log.Fatalf("error registering the register command: %v", err)
	}

	err = cmds.register("reset", handlerReset)
	if err != nil {
		log.Fatalf("error resetting the database: %v", err)
	}

	err = cmds.register("users", handlerUsers)
	if err != nil {
		log.Fatalf("error getting the user list: %v", err)
	}

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("not enough arguments passed in")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&currentState, cmd)
	if err != nil {
		log.Fatalf("error running the command: %v", err)
	}
}
