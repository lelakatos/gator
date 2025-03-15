package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lelakatos/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: login <username>")
	}

	userName := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("cannot login to a user that does not exist: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Username has been set to %s\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("usage: register <username>")
	}

	userName := cmd.args[0]
	usr := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	user, err := s.db.CreateUser(context.Background(), usr)
	if err != nil {
		return err
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("New user '%s' created: %+v", userName, user)

	return nil

}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments allowed for reset")
	}

	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Resetting the database was successful")

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments allowed for users")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
