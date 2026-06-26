package main

import (
	"github.com/Pakar040/gator/internal/config"
	"github.com/Pakar040/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
