package main

import (
	"context"

	"github.com/Southclaws/sampctl/print"
	"github.com/pkg/errors"
	"gopkg.in/urfave/cli.v1"

	"github.com/Southclaws/sampctl/download"
	"github.com/Southclaws/sampctl/runtime"
	"github.com/Southclaws/sampctl/types"
	"github.com/Southclaws/sampctl/util"
)

var serverRunFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "dir",
		Value: ".",
		Usage: "working directory for the server - by default, uses the current directory",
	},
	cli.BoolFlag{
		Name:  "container",
		Usage: "starts the server as a Linux container instead of running it in the current directory",
	},
	cli.BoolFlag{
		Name:  "mountCache",
		Usage: "if `--container` is set, mounts the local cache directory inside the container",
	},
	cli.BoolFlag{
		Name:  "ensurePlugins",
		Usage: "forces plugin binaries ensure before run",
	},
	cli.BoolFlag{
		Name:  "noCache",
		Usage: "forces download of plugins",
	},
}

func serverRun(c *cli.Context) error {
	if c.Bool("verbose") {
		print.SetVerbose()
	}

	dir := util.FullPath(c.String("dir"))
	container := c.Bool("container")
	mountCache := c.Bool("mountCache")
	ensurePlugins := c.Bool("ensurePlugins")
	noCache := c.Bool("noCache")

	cfg, err := runtime.NewConfigFromEnvironment(dir)
	if err != nil {
		return errors.Wrap(err, "failed to interpret directory as server runtime environment")
	}

	err = runtime.Ensure(context.Background(), gh, &cfg, noCache, ensurePlugins)
	if err != nil {
		return err
	}

	if container {
		cfg.Container = &types.ContainerConfig{MountCache: mountCache}
		cfg.AppVersion = c.App.Version
	}

	cacheDir, err := download.GetCacheDir()
	if err != nil {
		return errors.Wrap(err, "failed to get or create cache directory")
	}

	err = runtime.Run(context.Background(), cfg, cacheDir)

	return err
}
