package app

import (
	"fmt"

	"github.com/khapaev/solid-octo-meme/config"
	"github.com/khapaev/solid-octo-meme/internal/shorten"
	"github.com/khapaev/solid-octo-meme/pkg/logger"
)

func Run(cfg *config.Config) {
	_ = logger.New(cfg.Log.Level)

	ids := []uint32{0, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536}
	for _, id := range ids {
		fmt.Println(shorten.Shorten(id))
	}
}
