package main

import (
	"context"

	"github.com/Pakar040/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		loggedInUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return err
		}

		if err := handler(s, c, loggedInUser); err != nil {
			return err
		}

		return nil
	}
}
