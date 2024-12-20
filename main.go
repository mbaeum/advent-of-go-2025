package main

import (
	"context"
	"log/slog"

	"github.com/mbaeum/advent-of-go-2025/cmd"
)

func main() {

	cli := cmd.NewCLI()
	if err := cli.RunContext(context.Background()); err != nil {
		slog.Error("can't execute command", "error", err)
	}
}
