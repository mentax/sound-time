package main

import (
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

type params struct {
	input string
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func initCommandLine(args []string) error {
	app := cli.NewApp()
	app.Suggest = true
	app.EnableBashCompletion = true
	app.Name = "sound-time"
	app.Usage = "Get length (time) of mp3 file. \n\n" +
		"Example: \n" +
		"   sound-time 1.mp3 \n" +
		"   sound-time.exe data/2.mp3 "

	app.Version = version + " built in " + date + " from commit: [" + commit + "] by " + builtBy
	app.ArgsUsage = "[mp3 file]"

	app.Action = func(c *cli.Context) error {

		p, err := checkAndReturnParams(c)
		if err != nil {
			return err
		}

		if err = displayDuration(p); err != nil {
			return cli.Exit(err.Error(), 99)
		}
		return nil
	}

	return app.Run(args)
}

func checkAndReturnParams(c *cli.Context) (*params, error) {
	p := &params{}

	if c.NArg() != 1 {
		return nil, cli.Exit("Please provide exactly one mp3 file", 2)
	}

	f := c.Args().First()
	filename, err := filepath.Abs(f)
	if err != nil {
		return nil, cli.Exit("Wrong path to mp3 file "+filename, 3)
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, cli.Exit("mp3 file does not exist ( "+filename+" )", 4)
	}

	p.input = filename

	return p, nil
}
