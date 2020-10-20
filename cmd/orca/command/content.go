package command

import (
	"fmt"
	"io"
	"os"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/errdefs"
	digest "github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	cli "github.com/urfave/cli/v2"
)

var contentCatCommand = &cli.Command{
	Name:      "cat",
	Usage:     "print the content to standard output",
	ArgsUsage: "<digest>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		dgst, err := digest.Parse(c.Args().First())
		if err != nil {
			return err
		}

		ra, err := cln.ContentStore().ReaderAt(ctx, ocispec.Descriptor{Digest: dgst})
		if err != nil {
			return err
		}
		defer ra.Close()

		_, err = io.Copy(os.Stdout, content.NewReader(ra))
		return err
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

		return cln.ContentStore().Walk(ctx, func(info content.Info) error {
			fmt.Println(info.Digest)
			return nil
		})
	},
}

var contentRemoveCommand = &cli.Command{
	Name:      "remove",
	Aliases:   []string{"rm"},
	Usage:     "Remove one or more content by digest",
	ArgsUsage: "<digest> [digest...]",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		var dgsts []digest.Digest
		for _, arg := range c.Args().Slice() {
			dgst, err := digest.Parse(arg)
			if err != nil {
				return err
			}
			dgsts = append(dgsts, dgst)
		}

		var exitError error
		for _, dgst := range dgsts {
			err = cln.ContentStore().Delete(ctx, dgst)
			if err != nil {
				if !errdefs.IsNotFound(err) {
					if exitError == nil {
						exitError = err
					}
				}
				continue
			}
			fmt.Println(dgst)
		}

		return nil
	},
}
