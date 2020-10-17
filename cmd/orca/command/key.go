package command

import (
	"github.com/hinshun/orca/contentd"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

var keyAddCommand = &cli.Command{
	Name:      "add",
	Usage:     "Adds a public key",
	ArgsUsage: "<name> <pubkey>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var keyGenCommand = &cli.Command{
	Name:      "gen",
	Usage:     "Create a new keypair",
	ArgsUsage: "<name>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		_, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		return nil
	},
}

var keyListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List keys",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var keyRemoveCommand = &cli.Command{
	Name:      "remove",
	Aliases:   []string{"rm"},
	Usage:     "Remove one or more keys",
	ArgsUsage: "<key> [key...]",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var keyRenameCommand = &cli.Command{
	Name:      "rename",
	Usage:     "Rename a key",
	ArgsUsage: "<name> <newName>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}
