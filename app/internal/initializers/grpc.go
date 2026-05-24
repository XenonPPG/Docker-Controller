package initializers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ConnectGRPC(ctx context.Context, registerFunc func(server *grpc.Server)) error {
	lis, err := net.Listen(os.Getenv("GRPC_NETWORK"), ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	registerFunc(s)

	serverError := make(chan error, 1)

	go func() {
		log.Printf("Server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			serverError <- err
		}
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("gRPC server error: %w", err)
	case <-ctx.Done():
		log.Println("Received cancel signal. Shutting down gRPC server...")
		s.GracefulStop()
		log.Println("gRPC server gracefully stopped")
		return nil
	}
}
