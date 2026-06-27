package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Reset command expects 0 args")
	}

	feedUrl := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)

	return nil
}
