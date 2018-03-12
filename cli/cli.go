package cli

import (
	builtinLogger "log"
	"os"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/logrus"
	"github.com/urfave/cli"
)

var (
	version = "development"

	// log is a application-wide logger.
	log loggers.Logger
)

// Prerun configures application before running and executing from urfave/cli.
func Prerun(c *cli.Context) error {
	var err error

	err = initLogger(c)
	if err != nil {
		return err
	}

	return nil
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	app := cli.NewApp()
	app.Name = "wscp"
	app.Usage = "WebSocket Consumer & Producer"
	app.Action = RootAction
	app.Flags = RootFlags
	app.Commands = RootCommands
	app.Before = Prerun
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		builtinLogger.Fatal(err)
	}
}

// initLogger inits logger component.
func initLogger(c *cli.Context) error {
	var (
		lc  = logrus.Config{Level: "info"}
		err error
	)

	if c.Bool("debug") {
		lc.Level = "debug"
	}

	log, err = logrus.NewFromConfig(lc)
	if err != nil {
		return err
	}

	return nil
}
