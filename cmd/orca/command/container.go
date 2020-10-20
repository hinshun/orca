package command

import (
	"fmt"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/containerd/console"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/cmd/ctr/commands"
	"github.com/containerd/containerd/cmd/ctr/commands/tasks"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/oci"
	"github.com/containerd/containerd/platforms"
	"github.com/hinshun/orca/contentd"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

var containerExecCommand = &cli.Command{
	Name:      "exec",
	Usage:     "Run a command in a running container",
	ArgsUsage: "<container> <command> [arg...]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "tty",
			Aliases: []string{"t"},
			Usage:   "allocate a tty",
		},
		&cli.BoolFlag{
			Name:    "detach",
			Aliases: []string{"d"},
			Usage:   "detach from the process after it has started execution",
		},
	},
	Action: func(c *cli.Context) error {
		var (
			id       = c.Args().First()
			args     = c.Args().Tail()
			tty      = c.Bool("tty")
			detach   = c.Bool("detach")
			rootless = c.Bool("rootless")
		)

		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		container, err := cln.LoadContainer(ctx, id)
		if err != nil {
			return err
		}

		spec, err := container.Spec(ctx)
		if err != nil {
			return err
		}

		task, err := container.Task(ctx, nil)
		if err != nil {
			return err
		}

		pspec := spec.Process
		pspec.Terminal = tty
		pspec.Args = args

		var (
			ioCreator cio.Creator
			stdinC    = &stdinCloser{
				stdin: os.Stdin,
			}
		)

		cioOpts := []cio.Opt{
			cio.WithStreams(stdinC, os.Stdout, os.Stderr),
		}
		if tty {
			cioOpts = append(cioOpts, cio.WithTerminal)
		}
		if rootless {
			cioOpts = append(cioOpts, cio.WithFIFODir("/tmp/orca/fifo-dir"))
		}
		ioCreator = cio.NewCreator(cioOpts...)

		execId := "deadbeef"
		process, err := task.Exec(ctx, execId, pspec, ioCreator)
		if err != nil {
			return err
		}
		stdinC.closer = func() {
			process.CloseIO(ctx, containerd.WithStdinCloser)
		}
		if !detach {
			defer process.Delete(ctx)
		}

		statusC, err := process.Wait(ctx)
		if err != nil {
			return err
		}

		var con console.Console
		if tty {
			con = console.Current()
			defer con.Reset()
			if err := con.SetRaw(); err != nil {
				return err
			}
		}
		if !detach {
			if tty {
				err = tasks.HandleConsoleResize(ctx, process, con)
				if err != nil {
					log.Printf("console resize err: %s", err)
				}
			} else {
				sigc := commands.ForwardAllSignals(ctx, process)
				defer commands.StopCatch(sigc)
			}
		}

		err = process.Start(ctx)
		if err != nil {
			return err
		}
		if detach {
			return nil
		}

		status := <-statusC
		code, _, err := status.Result()
		if err != nil {
			return err
		}
		if code != 0 {
			return cli.NewExitError("", int(code))
		}
		return nil
	},
}

var containerLogsCommand = &cli.Command{
	Name:      "logs",
	Usage:     "Fetch the logs of a container",
	ArgsUsage: "<container>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		return nil
	},
}

var containerListCommand = &cli.Command{
	Name:      "list",
	Usage:     "List containers",
	Aliases:   []string{"ls"},
	ArgsUsage: "<container>",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		containers, err := cln.ContainerService().List(ctx)
		if err != nil {
			return err
		}

		for _, container := range containers {
			fmt.Println(container.ID)
		}
		return nil
	},
}

var containerRemoveCommand = &cli.Command{
	Name:      "remove",
	Usage:     "Remove one or more containers",
	Aliases:   []string{"rm"},
	ArgsUsage: "<container> [container...]",
	Flags:     []cli.Flag{},
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			return nil
		}

		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		for _, id := range c.Args().Slice() {
			container, err := cln.LoadContainer(ctx, id)
			if err != nil {
				return err
			}

			task, err := container.Task(ctx, cio.Load)
			if err != nil {
				return err
			}

			err = task.Kill(ctx, syscall.SIGKILL, containerd.WithKillAll)
			if err != nil {
				return err
			}

			fmt.Printf("Deleting task with pid %d\n", task.Pid())
			_, err = task.Delete(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("Deleting container %s\n", id)
			err = container.Delete(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("Removed container %q\n", id)
		}

		return nil
	},
}

var containerRunCommand = &cli.Command{
	Name:      "run",
	Usage:     "Run a command in a new container",
	ArgsUsage: "<image> <command> <arg...>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "rm",
			Usage: "automatically remove the container when it exits",
		},
		&cli.BoolFlag{
			Name:    "tty",
			Aliases: []string{"t"},
			Usage:   "allocate a tty",
		},
		&cli.BoolFlag{
			Name:    "detach",
			Aliases: []string{"d"},
			Usage:   "run container in background and print container ID",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Assign a name to the container",
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			return errors.Errorf("must specify an image")
		}

		var (
			ref      = c.Args().First()
			args     = c.Args().Tail()
			rm       = c.Bool("rm")
			tty      = c.Bool("tty")
			detach   = c.Bool("detach")
			rootless = c.Bool("rootless")
			id       = c.String("name")
		)

		if id == "" {
			// todo: generate id
			id = "hello"
		}

		cln, ctx, cancel, err := NewClient(c)
		if err != nil {
			return err
		}
		defer cancel()

		img, err := cln.GetImage(ctx, ref)
		if err != nil {
			if !errdefs.IsNotFound(err) {
				return errors.Wrap(err, "failed to get image")
			}

			contentdCln, err := contentd.NewClient(c.String("contentd-address"))
			if err != nil {
				return errors.Wrap(err, "failed to create contentd client")
			}

			image, err := PullImage(ctx, contentdCln, cln, ref)
			if err != nil {
				return err
			}

			img = containerd.NewImageWithPlatform(cln, image, platforms.Default())
		}

		snapshotter := containerd.DefaultSnapshotter
		unpacked, err := img.IsUnpacked(ctx, snapshotter)
		if err != nil {
			return err
		}
		if !unpacked {
			if err := img.Unpack(ctx, snapshotter); err != nil {
				return err
			}
		}

		var (
			opts  []oci.SpecOpts
			cOpts []containerd.NewContainerOpts
			s     specs.Spec
		)

		opts = append(opts,
			oci.WithDefaultSpec(),
			oci.WithDefaultUnixDevices,
			oci.WithImageConfigArgs(img, args),
		)
		if tty {
			opts = append(opts, oci.WithTTY)
		}
		if rootless {
			opts = append(opts, oci.WithCgroup(""))
		}

		cOpts = append(cOpts,
			containerd.WithImage(img),
			containerd.WithSnapshotter(snapshotter),
			containerd.WithNewSnapshot(id, img),
			containerd.WithImageStopSignal(img, "SIGTERM"),
			containerd.WithSpec(&s, opts...),
		)

		container, err := cln.NewContainer(ctx, id, cOpts...)
		if err != nil {
			return errors.Wrap(err, "failed to create container")
		}

		if rm {
			defer container.Delete(ctx)
		}

		var con console.Console
		if tty {
			con = console.Current()
			defer con.Reset()

			err = con.SetRaw()
			if err != nil {
				return err
			}
		}

		var cioOpts []cio.Opt
		if rootless {
			cioOpts = append(cioOpts, cio.WithFIFODir("/tmp/orca/fifo-dir"))
		}

		task, err := tasks.NewTask(ctx, cln, container, "", con, false, "", cioOpts)
		if err != nil {
			return errors.Wrap(err, "failed to create task")
		}

		var statusC <-chan containerd.ExitStatus
		if !detach {
			defer task.Delete(ctx)
			statusC, err = task.Wait(ctx)
			if err != nil {
				return err
			}
		}

		err = task.Start(ctx)
		if err != nil {
			return err
		}
		if detach {
			fmt.Println(id)
			return nil
		}

		if tty {
			err = tasks.HandleConsoleResize(ctx, task, con)
			if err != nil {
				log.Printf("console resize err: %s", err)
			}
		} else {
			sigc := commands.ForwardAllSignals(ctx, task)
			defer commands.StopCatch(sigc)
		}

		status := <-statusC
		code, _, err := status.Result()
		if err != nil {
			return err
		}

		_, err = task.Delete(ctx)
		if err != nil {
			return err
		}
		if code != 0 {
			return cli.NewExitError("", int(code))
		}
		return nil
	},
}

type stdinCloser struct {
	stdin  *os.File
	closer func()
}

func (s *stdinCloser) Read(p []byte) (int, error) {
	n, err := s.stdin.Read(p)
	if err == io.EOF {
		if s.closer != nil {
			s.closer()
		}
	}
	return n, err
}
