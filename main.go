package main

import (
	"github.com/PrakharSrivastav/album-service-grpc/internal"
	"github.com/PrakharSrivastav/album-service-grpc/internal/rpc"
	"github.com/PrakharSrivastav/album-service-grpc/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	checkErr(err)
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(sameSpan),
	)

	// Create a new application
	app := internal.NewApplication(tracer)

	// Create a new db connection
	connection, err := sqlx.Open("sqlite3", "./album.db")
	checkErr(err)

	// Create internal service
	internal := service.New(connection, tracer)
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

const (
	// Our service name.
	serviceName = "album-service"

	// Host + port of our service.
	hostPort = "127.0.0.1:6564"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://zipkin:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false
)
