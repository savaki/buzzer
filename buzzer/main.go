package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/savaki/buzzer/app"
)

type Options struct {
	Port string
}

func Opts(c *cli.Context) Options {
	return Options{
		Port: c.String("port"),
	}
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{"port", "8080", "port to run on", "PORT"},
	}
	app.Action = Run
	app.Run(os.Args)
}

func Run(c *cli.Context) {
	opts := Opts(c)

	router := app.Routes()
	http.ListenAndServe(":"+opts.Port, router)
}
