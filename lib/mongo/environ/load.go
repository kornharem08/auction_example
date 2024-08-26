package environ

import (
	"log"

	"github.com/caarlos0/env/v6"
)

func Load[T any]() T {
	var cfg T
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
