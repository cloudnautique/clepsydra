package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/clepsydra/events"
	"github.com/urfave/cli"
)

// VERSION of the application
var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "clepsydra"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = start

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
