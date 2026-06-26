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
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
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

func handlerUsers(s *state, cmd command) error {
	currentUserName := s.cfg.CurrentUsername

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		fmt.Printf("* %s", u.Name)

		if u.Name == currentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Reset command expects 0 args")
	}

	if err := s.db.DeleteAll(context.Background()); err != nil {
		return err
	}

	fmt.Println("All users have been deleted")
	return nil
}
