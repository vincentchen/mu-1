package cli

import (
	"fmt"
	"github.com/stelligent/mu/common"
	"github.com/urfave/cli"
)

func newServicesCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:    "service",
		Aliases: []string{"svc"},
		Usage:   "options for managing services",
		Subcommands: []cli.Command{
			*newServicesShowCommand(ctx),
			*newServicesDeployCommand(ctx),
			*newServicesSetenvCommand(ctx),
			*newServicesUndeployCommand(ctx),
		},
	}

	return cmd
}

func newServicesShowCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:  "show",
		Usage: "show service details",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to show",
			},
		},
		Action: func(c *cli.Context) error {
			service := c.String("service")
			fmt.Printf("showing service: %s\n", service)
			return nil
		},
	}

	return cmd
}

func newServicesDeployCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:      "deploy",
		Usage:     "deploy service to environment",
		ArgsUsage: "<environment>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to deploy",
			},
		},
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			serviceName := c.String("service")
			fmt.Printf("deploying service: %s to environment: %s\n", serviceName, environmentName)
			return nil
		},
	}

	return cmd
}

func newServicesSetenvCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:      "setenv",
		Usage:     "set environment variable",
		ArgsUsage: "<environment> <key1>=<value1>...",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to deploy",
			},
		},
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			serviceName := c.String("service")
			keyvals := c.Args().Tail()
			fmt.Printf("setenv service: %s to environment: %s with vals: %s\n", serviceName, environmentName, keyvals)
			return nil
		},
	}

	return cmd
}

func newServicesUndeployCommand(ctx *common.Context) *cli.Command {
	cmd := &cli.Command{
		Name:      "undeploy",
		Usage:     "undeploy service from environment",
		ArgsUsage: "<environment>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "service, s",
				Usage: "service to undeploy",
			},
		},
		Action: func(c *cli.Context) error {
			environmentName := c.Args().First()
			serviceName := c.String("service")
			fmt.Printf("undeploying service: %s to environment: %s\n", serviceName, environmentName)
			return nil
		},
	}

	return cmd
}
