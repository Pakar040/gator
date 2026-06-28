package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Pakar040/gator/internal/config"
	"github.com/Pakar040/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	s := state{cfg: &cfg, db: dbQueries}

	cmds := commands{
		registry: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAggregate)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

	if len(os.Args) < 2 {
		fmt.Println("Command not recognized")
		os.Exit(1)
	}

	if err := cmds.run(&s, command{name: os.Args[1], args: os.Args[2:]}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
