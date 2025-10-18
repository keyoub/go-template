package main

import (
	"context"
	"log"
	"os"

	"sampleapp/cmd"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "sampleapp",
		Usage: "A sample Go web service",
		Commands: []*cli.Command{
			{
				Name:  "apiserver",
				Usage: "Start the API server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "8080",
						Usage: "Port to run the server on",
					},
				},
				Action: cmd.RunServer,
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
