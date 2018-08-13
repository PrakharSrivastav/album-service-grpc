package main

import (
	"fmt"

	"github.com/PrakharSrivastav/album-service-grpc/internal"
	"github.com/PrakharSrivastav/album-service-grpc/internal/rpc"
)

func main() {
	app := internal.NewApplication()

	albumService := rpc.New(nil)

	app.GrpcServer.Add(albumService)
	err := app.GrpcServer.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}
