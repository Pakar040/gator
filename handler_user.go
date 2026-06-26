package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Pakar040/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects 1 arg: username")
	}
	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User has been set: %s\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects 1 arg: username")
	}
	name := cmd.args[0]

	newUser := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}

	createdUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(createdUser.Name); err != nil {
		return err
	}

	fmt.Printf("User has been created: %v\n", createdUser)
	return nil
}
