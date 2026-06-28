package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Pakar040/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Follow command expects 1 arg: url")
	}
	url := cmd.args[0]

	feedToFollow, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedToFollow.ID,
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return err
	}

	fmt.Printf("%s now follows %s\n", user.Name, feedToFollow.Name)
	fmt.Printf("Created Feed Follow Record: \n%v\n", createdFeedFollow)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("Following command expects 0 args")
	}

	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollowing {
		fmt.Printf(" * %s\n", feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Unfollow command expects 1 arg: url")
	}
	url := cmd.args[0]

	feedToUnfollow, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	deletedFeedFollow, err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedToUnfollow.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s had unfollowed %s\n", user.Name, feedToUnfollow.Name)
	fmt.Printf("Deleted Feed Follow Record:\n%v", deletedFeedFollow)
	return nil
}
