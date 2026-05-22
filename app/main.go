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
	// error group
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// TODO: review if I need this
	g, ctx := errgroup.WithContext(ctx)

	// docker
	cli := initializers.InitDockerClient()
	defer func(c *client.Client) {
		if err := c.Close(); err != nil {
			log.Printf("Error closing docker client: %s", err.Error())
		}
	}(cli)

	// gRPC
	server := &controllers.Server{}
	g.Go(func() error {
		return initializers.ConnectGRPC(server.RegisterFunc())
	})

	if err := g.Wait(); err != nil {
		log.Fatal("Program terminated: " + err.Error())
	}
}
