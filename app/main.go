package main

import (
	"DockerController/app/internal/controllers"
	"DockerController/app/internal/initializers"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/moby/moby/client"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	cli := initializers.InitDockerClient()
	defer func(c *client.Client) {
		if err := c.Close(); err != nil {
			log.Printf("Error closing docker client: %s", err.Error())
		}
	}(cli)

	server := &controllers.Server{}
	g.Go(func() error {
		return initializers.ConnectGRPC(gCtx, server.RegisterFunc())
	})

	if err := g.Wait(); err != nil {
		log.Fatal("Program terminated: " + err.Error())
	}
}
