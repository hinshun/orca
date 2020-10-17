package command

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var contentCatCommand = &cli.Command{
	Name:      "cat",
	Usage:     "print the content to standard output",
	ArgsUsage: "<digest>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var contentListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List all content",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		contents, err := cln.ImageService().List(ctx)
		if err != nil {
			return err
		}

		for _, content := range contents {
			fmt.Println(content.Name)
		}
		return nil
	},
}

var contentRemoveCommand = &cli.Command{
	Name:      "remove",
	Aliases:   []string{"rm"},
	Usage:     "Remove one or more content by digest",
	ArgsUsage: "<digest> [digest...]",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}
