package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func newHelloCmd(l *slog.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Say Hello",
		RunE: func(cmd *cobra.Command, args []string) error {
			// ctx := cmd.Context()
			l = l.With("cmd", "hello")
			l.Debug("Running hello")
			name := "World"
			if len(args) > 0 {
				name = args[0]
			}
			fmt.Fprintf(cmd.OutOrStderr(), "Hello %s!", name)
			return nil
		},
	}
}
