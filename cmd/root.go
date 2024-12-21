package cmd

import (
	"context"
	"log/slog"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
	"github.com/spf13/cobra"
)

type CLI struct {
	l   *slog.Logger
	cmd *cobra.Command
}

func (c CLI) RunContext(ctx context.Context) error {
	err := c.cmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c CLI) registerCommands() {
	c.cmd.AddCommand(newHelloCmd(c.l))
}

func NewCLI() CLI {
	l := util.NewLogger()
	cmd := &cobra.Command{
		Use: "aoc",
	}
	cli := CLI{l: l, cmd: cmd}
	cli.registerCommands()
	return cli
}
