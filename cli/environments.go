package cli

import (
	"errors"
	"github.com/stelligent/mu/common"
	"github.com/stelligent/mu/workflows"
	"github.com/urfave/cli"
	"os"
	"strings"
	"time"
)

func newEnvironmentsCommand(ctx *common.Context) *cli.Command {

	cmd := &cli.Command{
		Name:    "environment",
		Aliases: []string{"env"},
		Usage:   "options for managing environments",
		Subcommands: []cli.Command{
			*newEnvironmentsListCommand(ctx),
			*newEnvironmentsShowCommand(ctx),
			*newEnvironmentsUpsertCommand(ctx),
			*newEnvironmentsTerminateCommand(ctx),
			*newEnvironmentsLogsCommand(ctx),
		},
	}

	return cmd
}

func newEnvironmentsUpsertCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:      "upsert",
		Aliases:   []string{"up"},
		Usage:     "create/update an environment",
		ArgsUsage: "<environment>",
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			if len(environmentName) == 0 {
				cli.ShowCommandHelp(c, "upsert")
				return errors.New("environment must be provided")
			}

			workflow := workflows.NewEnvironmentUpserter(ctx, environmentName)
			return workflow()
		},
	}

	return cmd
}

func newEnvironmentsListCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "list environments",
		Action: func(c *cli.Context) error {
			workflow := workflows.NewEnvironmentLister(ctx, os.Stdout)
			return workflow()
		},
	}

	return cmd
}

func newEnvironmentsShowCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "show",
		Usage: "show environment details",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "format, f",
				Usage: "output format, either 'json' or 'cli' (default: cli)",
				Value: "cli",
			},
		},
		ArgsUsage: "<environment>",
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			if len(environmentName) == 0 {
				cli.ShowCommandHelp(c, "show")
				return errors.New("environment must be provided")
			}
			workflow := workflows.NewEnvironmentViewer(ctx, c.String("format"), environmentName, os.Stdout)
			return workflow()
		},
	}

	return cmd
}
func newEnvironmentsTerminateCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:      "terminate",
		Aliases:   []string{"term"},
		Usage:     "terminate an environment",
		ArgsUsage: "<environment>",
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			if len(environmentName) == 0 {
				cli.ShowCommandHelp(c, "terminate")
				return errors.New("environment must be provided")
			}
			workflow := workflows.NewEnvironmentTerminator(ctx, environmentName)
			return workflow()
		},
	}

	return cmd
}
func newEnvironmentsLogsCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "logs",
		Usage: "show environment logs",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "follow, f",
				Usage: "follow logs for latest changes",
			},
			cli.DurationFlag{
				Name:  "search-duration, t",
				Usage: "duration to go into the past for searching (e.g. 5m for 5 minutes)",
				Value: 1 * time.Minute,
			},
		},
		ArgsUsage: "<environment> [<filter>...]",
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			if len(environmentName) == 0 {
				cli.ShowCommandHelp(c, "logs")
				return errors.New("environment must be provided")
			}

			workflow := workflows.NewEnvironmentLogViewer(ctx, c.Duration("search-duration"), c.Bool("follow"), environmentName, os.Stdout, strings.Join(c.Args().Tail(), " "))
			return workflow()
		},
	}

	return cmd
}
