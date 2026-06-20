package main

import (
	"fmt"

	"github.com/Pakar040/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("Alek")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)
}
