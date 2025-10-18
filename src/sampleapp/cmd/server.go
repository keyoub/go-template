package cmd

import (
	"context"
	"sampleapp/api"

	"github.com/urfave/cli/v3"
)

func RunServer(ctx context.Context, cmd *cli.Command) error {
	port := cmd.String("port")
	return api.StartServer(port)
}
