package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	fakeRuntimePlugin := cli.NewApp()
	fakeRuntimePlugin.Name = "fakeRuntimePlugin"
	fakeRuntimePlugin.Usage = "I am FakeRuntimePlugin!"

	fakeRuntimePlugin.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "debug",
		},
		cli.StringFlag{
			Name: "log",
		},
		cli.StringFlag{
			Name: "newuidmap",
		},
		cli.StringFlag{
			Name: "newgidmap",
		},
	}

	fakeRuntimePlugin.Commands = []cli.Command{
		CreateCommand,
		StateCommand,
		EventsCommand,
	}

	_ = fakeRuntimePlugin.Run(os.Args)
}

var CreateCommand = cli.Command{
	Name: "create",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "no-new-keyring",
		},
		cli.StringFlag{
			Name: "bundle",
		},
		cli.StringFlag{
			Name: "pid-file",
		},
	},

	Action: func(ctx *cli.Context) error {
		err := ioutil.WriteFile(filepath.Join(os.TempDir(), "args"), []byte(strings.Join(os.Args, " ")), 0777)
		if err != nil {
			panic(err)
		}

		return nil
	},
}

var StateCommand = cli.Command{
	Name:  "state",
	Flags: []cli.Flag{},

	Action: func(ctx *cli.Context) error {
		fmt.Printf("{}")
		return nil
	},
}

var EventsCommand = cli.Command{
	Name:  "events",
	Flags: []cli.Flag{},

	Action: func(ctx *cli.Context) error {
		fmt.Printf("{}")
		return nil
	},
}
