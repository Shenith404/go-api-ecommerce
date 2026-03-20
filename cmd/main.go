package main

import (
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{},
	}
	api := application{
		config: cfg,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		os.Exit(1)
	}

}