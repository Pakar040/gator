package main

import (
	"context"
	"fmt"
)

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
