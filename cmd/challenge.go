package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

func newChallengeCmd(l *slog.Logger) *cobra.Command {
	l.Debug("Running cmd", "cmd", "challenge")

	cfg, err := util.NewConfig("config.json")
	if err != nil {
		return nil
	}
	l.Debug("Got config", "config", cfg)

	c := &cobra.Command{
		Use:   "challenge",
		Short: "Manage challenges",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	c.AddCommand(newNewChallengeCmd(l, cfg))
	c.AddCommand(newRunChallengeCmd(l, cfg))
	return c
}

func newNewChallengeCmd(l *slog.Logger, cfg *util.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Create a new challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			l = l.With("subcmd", "new")
			path, err := util.GetTargetPath("challenges")
			if err != nil {
				l.Error("Error getting target path", "error", err)
				return err
			}
			fm, err := util.NewFileManger(path, cfg)
			if err != nil {
				l.Error("Error creating file manager", "error", err)
				return err
			}

			if len(args) == 0 {
				l.Error("Need an id as arg")
				return errors.New("no input args provided")
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				l.Error("Error parsing argument", "error", err)
				return err
			}
			err = fm.NewChallenge(id)
			if err != nil {
				l.Error("Could not create new challenge", "error", err)
				return err
			}

			return nil
		},
	}
}

func newRunChallengeCmd(l *slog.Logger, cfg *util.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run an existing challenge",
		RunE: func(cmd *cobra.Command, args []string) error {
			l = l.With("subcmd", "run")
			id, err := strconv.Atoi(args[0])
			if err != nil {
				l.Error("Error parsing argument", "error", err)
				return err
			}
			fm, err := util.NewChallengeFileManager(id, cfg)
			if err != nil {
				l.Error("Error creating file manager", "error", err)
				return err
			}
			contents, err := fm.ReadFile("data_test.txt")
			if err != nil {
				return err
			}
			fmt.Printf("contents: %s", contents)

			return nil
		},
	}
}
