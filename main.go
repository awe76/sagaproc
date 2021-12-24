package main

import (
	"github.com/awe76/sagaproc/handler"
	pb "github.com/awe76/sagaproc/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "sagaproc"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterSagaprocHandler(srv.Server(), new(handler.Sagaproc))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
