package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/Netflix/p2plab/errdefs"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/platforms"
	"github.com/containerd/containerd/remotes"
	"github.com/hinshun/orca/contentd"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

var imageListCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List images",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		images, err := cln.ImageService().List(ctx)
		if err != nil {
			return err
		}

		for _, image := range images {
			fmt.Println(image.Name)
		}
		return nil
	},
}

var imagePullCommand = &cli.Command{
	Name:      "pull",
	Usage:     "Pull an image from your peers.",
	ArgsUsage: "<ref>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return errors.Errorf("must specify a ref")
		}

		ref := c.Args().First()

		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		contentdCln, err := contentd.NewClient(c.String("contentd-address"))
		if err != nil {
			return errors.Wrap(err, "failed to create contentd client")
		}

		image, err := PullImage(ctx, contentdCln, cln, ref)
		if err != nil {
			return err
		}

		fmt.Printf("Pulled image %s@%s\n", image.Name, image.Target.Digest)
		return nil
	},
}

var imageRemoveCommand = &cli.Command{
	Name:      "remove",
	Aliases:   []string{"rm"},
	Usage:     "Remove one or more images",
	ArgsUsage: "<image> [image...]",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		names := c.Args().Slice()

		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		for _, name := range names {
			err = cln.ImageService().Delete(ctx, name)
			if err != nil {
				return err
			}
			fmt.Printf("Deleted image %q\n", name)
		}

		return nil
	},
}

func PullImage(ctx context.Context, contentdCln *contentd.Client, containerdCln *containerd.Client, ref string) (img images.Image, err error) {
	var (
		store    = containerdCln.ContentStore()
		resolver = contentdCln.Resolver()
	)

	name, desc, err := resolver.Resolve(ctx, ref)
	if err != nil {
		return
	}

	fetcher, err := resolver.Fetcher(ctx, name)
	if err != nil {
		return
	}

	// Get all the children for a descriptor
	childrenHandler := images.ChildrenHandler(store)

	// Set any children labels for that content
	childrenHandler = images.SetChildrenMappedLabels(store, childrenHandler, nil)

	// Filter children by platforms if specified.
	childrenHandler = images.FilterPlatforms(childrenHandler, platforms.Default())

	handlers := []images.Handler{
		remotes.FetchHandler(store, fetcher),
		childrenHandler,
	}

	handler := images.Handlers(handlers...)
	err = images.Dispatch(ctx, handler, nil, desc)
	if err != nil {
		return
	}

	img = images.Image{
		Name:   name,
		Target: desc,
	}

	is := containerdCln.ImageService()
	for {
		created, err := is.Create(ctx, img)
		if err != nil {
			if !strings.Contains(err.Error(), errdefs.ErrAlreadyExists.Error()) {
				return img, err
			}

			updated, err := is.Update(ctx, img)
			if err != nil {
				// if image was removed, try create again
				if errdefs.IsNotFound(err) {
					continue
				}
				return img, err
			}

			return updated, nil
		}
		return created, nil
	}
}
