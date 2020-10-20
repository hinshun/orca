package command

import (
	cli "github.com/urfave/cli/v2"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "orca"
	app.Usage = "cli for container management"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "rootless",
			EnvVars: []string{"ORCA_CONTAINERD_ROOTLESS"},
			Usage:   "containerd is running rootless",
		},
		&cli.StringFlag{
			Name:    "containerd-address",
			Aliases: []string{"ctrd-addr"},
			Usage:   "containerd address",
			Value:   "/run/user/1001/containerd/containerd.sock",
		},
		&cli.StringFlag{
			Name:    "contentd-address",
			Aliases: []string{"contentd-addr"},
			Usage:   "contentd address",
			Value:   "/run/user/1001/contentd/contentd.sock",
		},
		&cli.StringFlag{
			Name:  "namespace",
			Usage: "namespace to use with commands",
			Value: "default",
		},
		&cli.DurationFlag{
			Name:  "timeout",
			Usage: "total timeout for commands",
		},
	}

	runCommand := *containerRunCommand
	runCommand.Category = "shortcut"

	execCommand := *containerExecCommand
	execCommand.Category = "shortcut"

	pullCommand := *imagePullCommand
	pullCommand.Category = "shortcut"

	app.Commands = []*cli.Command{
		// podCommand,
		containerCommand,
		imageCommand,
		contentCommand,
		keyCommand,

		&runCommand,
		&execCommand,
		&pullCommand,
	}

	return app
}

var podCommand = &cli.Command{
	Name:        "pod",
	Usage:       "Manage pods",
	Category:    "manage",
	Subcommands: []*cli.Command{},
}

var containerCommand = &cli.Command{
	Name:     "container",
	Aliases:  []string{"ctr"},
	Usage:    "Manage containers",
	Category: "manage",
	Subcommands: []*cli.Command{
		containerExecCommand,
		containerLogsCommand,
		containerListCommand,
		containerRemoveCommand,
		containerRunCommand,
	},
}

var imageCommand = &cli.Command{
	Name:     "image",
	Aliases:  []string{"img"},
	Usage:    "Manage images",
	Category: "manage",
	Subcommands: []*cli.Command{
		imageListCommand,
		imagePullCommand,
		imageRemoveCommand,
	},
}

var contentCommand = &cli.Command{
	Name:     "content",
	Usage:    "Manage content",
	Category: "manage",
	Subcommands: []*cli.Command{
		contentCatCommand,
		contentListCommand,
		contentRemoveCommand,
	},
}

var keyCommand = &cli.Command{
	Name:     "key",
	Usage:    "Manage keys",
	Category: "manage",
	Subcommands: []*cli.Command{
		keyAddCommand,
		keyGenCommand,
		keyListCommand,
		keyRemoveCommand,
		keyRenameCommand,
	},
}
