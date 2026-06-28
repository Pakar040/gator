package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Pakar040/gator/internal/database"
	"github.com/google/uuid"
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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("AddFeed command expects 2 args: name, url")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currUser.ID,
	}

	createdFeed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: newFeed.CreatedAt,
		UpdatedAt: newFeed.UpdatedAt,
		UserID:    currUser.ID,
		FeedID:    newFeed.ID,
	}

	if _, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow); err != nil {
		return err
	}

	fmt.Println(createdFeed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Feeds command expects 0 args")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf(" - Feed: %s; URL: %s; Username: %s\n", feed.Feedname, feed.Url, feed.Username)
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Follow command expects 1 arg: url")
	}
	url := cmd.args[0]

	feedToFollow, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return err
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currUser.ID,
		FeedID:    feedToFollow.ID,
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return err
	}

	fmt.Printf("%s now follows %s\n", currUser.Name, feedToFollow.Name)
	fmt.Printf("Created Feed Follow Record: \n%v\n", createdFeedFollow)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Following command expects 0 args")
	}

	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollowing {
		fmt.Printf(" * %s\n", feed.FeedName)
	}

	return nil
}
