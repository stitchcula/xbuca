package main

import (
	"context"
	"flag"
	"m7s.live/engine/v4"
)

var (
	c = flag.String("c", "etc/config.yaml", "config file")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	if err := engine.Run(ctx, *c); err != nil {
		panic(err)
	}

	cancel()
}
