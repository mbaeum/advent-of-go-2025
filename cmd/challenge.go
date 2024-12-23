package cmd

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/mbaeum/advent-of-go-2025/challenges"
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

			factory := challenges.NewChallengeFactory()

			challenge, err := factory.GetChallenge(id)
			if err != nil {
				return err
			}

			challenge.SetSessionCookie(cfg.SessionCookie)
			res, err := challenge.RunPartOne(challenges.MainMode)
			if err != nil {
				return err
			}

			l.Info("Challenge returned", "challenge", 1, "result", res)

			res, err = challenge.RunPartTwo(challenges.MainMode)
			if err != nil {
				return err
			}

			l.Info("Challenge returned", "challenge", 2, "result", res)

			return nil

		},
	}
}
