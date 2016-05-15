package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	Name    = "resolver"
	Usage   = "simple http service to resolve hashed assets"
	Version = "0.1.0"
)

var (
	LoadedStore   *Store
	ResolverStore string
	ResolverUser  string
	ResolverKey   string
)

// preload initializes any global options and configuration
// before the main sub commands are run
func preload(context *cli.Context) error {
	// set loglevel
	if context.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	// set store path
	ResolverStore = context.GlobalString("store")

	// set user and key
	ResolverUser = context.GlobalString("user")
	ResolverKey = context.GlobalString("key")

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = Name
	app.Usage = Usage
	app.Author = "@erroneousboat"
	app.Version = Version

	app.Before = preload

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "run in debug mode",
		},
		cli.StringFlag{
			Name:   "host",
			Value:  "0.0.0.0",
			Usage:  "host on which to serve",
			EnvVar: "RESOLVER_HOST",
		},
		cli.IntFlag{
			Name:   "port, p",
			Value:  4000,
			Usage:  "port on which to serve",
			EnvVar: "RESOLVER_PORT",
		},
		cli.StringFlag{
			Name:   "store, s",
			Value:  "datastore.json",
			Usage:  "path to data store file",
			EnvVar: "RESOLVER_STORE",
		},
		cli.StringFlag{
			Name:   "user, u",
			Value:  "admin",
			Usage:  "user that will be able to access the server",
			EnvVar: "RESOLVER_USER",
		},
		cli.StringFlag{
			Name:   "key, k",
			Value:  "$apr1$rZTtRFsq$BK1GqipOMOTpWuJYuDtJ01",
			Usage:  "key to be used for authentication",
			EnvVar: "RESOLVER_KEY",
		},
	}

	app.Action = func(c *cli.Context) {
		// load data store
		LoadedStore = LoadStore(ResolverStore)

		// create a negroni instance
		n := ServerSetup()

		// run the server
		n.Run(fmt.Sprintf("%s:%d", c.String("host"), c.Int("port")))
	}

	app.Run(os.Args)
}
