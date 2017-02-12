package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/clepsydra/events"
	"github.com/urfave/cli"
)

// VERSION of the application
var VERSION = "v0.0.0-dev"

func beforeApp(c *cli.Context) error {
    if c.GlobalBool("debug") {
        logrus.SetLevel(logrus.DebugLevel)
    }
    return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "clepsydra"
	app.Version = VERSION
	app.Usage = "clepsydra"
	app.Action = start
    app.Before = beforeApp
    app.Flags = []cli.Flag{
            cli.BoolFlag{
                    Name: "debug,d",
            },
    }

	app.Run(os.Args)
}

func start(c *cli.Context) error {
	router, err := events.NewEventRouter()
	if err != nil {
		logrus.Fatal(err)
	}

	events.StartRouter(router)

	return nil
}
