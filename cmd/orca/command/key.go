package command

import (
	"context"
	"fmt"

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
		ctx := context.Background()
		cln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		name, pubKey := c.Args().Get(0), c.Args().Get(1)
		err = cln.Keystore().Add(ctx, name, pubKey)
		if err != nil {
			return err
		}

		fmt.Printf("Added key %s\n", pubKey)
		return nil
	},
}

var keyGenCommand = &cli.Command{
	Name:      "gen",
	Usage:     "Create a new keypair",
	ArgsUsage: "<name>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		cln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		name := c.Args().First()
		key, err := cln.Keystore().Generate(ctx, name, "ed25519", 0)
		if err != nil {
			return err
		}

		pubKey, err := contentd.Libp2pCidFromPubKey(key.PublicKey.Data)
		if err != nil {
			return err
		}
		fmt.Printf("Generated key %s %s\n", key.Name, pubKey)
		return nil
	},
}

var keyListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List keys",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		cln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		keys, err := cln.Keystore().List(ctx)
		if err != nil {
			return err
		}

		for _, key := range keys {
			pubKey, err := contentd.Libp2pCidFromPubKey(key.PublicKey.Data)
			if err != nil {
				return err
			}
			fmt.Printf("Key: %s, Public Key: %s\n", key.Name, pubKey)
		}
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
		ctx := context.Background()
		cln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		names := c.Args().Slice()
		deleted, err := cln.Keystore().Remove(ctx, names...)
		if err != nil {
			return err
		}

		for _, name := range deleted {
			fmt.Printf("Deleted %s\n", name)
		}
		return nil
	},
}

var keyRenameCommand = &cli.Command{
	Name:      "rename",
	Usage:     "Rename a key",
	ArgsUsage: "<oldName> <newName>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		ctx := context.Background()
		cln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		oldName, newName := c.Args().Get(0), c.Args().Get(1)
		err = cln.Keystore().Rename(ctx, oldName, newName)
		if err != nil {
			return err
		}

		fmt.Printf("Renamed %q to %q\n", oldName, newName)
		return nil
	},
}
