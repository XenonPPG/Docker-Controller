package main

import (
	"DockerController/app/internal/controllers"
	"DockerController/app/internal/initializers"
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	// error group
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// gRPC
	server := &controllers.Server{}
	g.Go(func() error {
		return initializers.ConnectGRPC(server.RegisterFunc())
	})

	if err := g.Wait(); err != nil {
		log.Fatal("Program terminated: " + err.Error())
	}
}
