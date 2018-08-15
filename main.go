package main

import (
	"github.com/PrakharSrivastav/album-service-grpc/internal"
	"github.com/PrakharSrivastav/album-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/album-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create a new application
	app := internal.NewApplication()

	// Create a new db connection
	connection, err := sqlx.Open("sqlite3", "./album.db")
	checkErr(err)

	// Create internal service
	internal := service.New(connection)
	checkErr(internal.CleanupAndInit())

	// Inject internal for rpc interfaces
	albumService := rpc.New(internal)

	// inject album service to grpc server
	app.GrpcServer.Add(albumService)

	// start grpc server
	checkErr(app.GrpcServer.Start())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
