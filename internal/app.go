package internal

import (
	"github.com/PrakharSrivastav/album-service-grpc/internal/backend"
)

// Application is a wrapper for all the backend infrastructure
type Application struct {
	GrpcServer *backend.Server
}

// NewApplication returns an instance of Application
func NewApplication() *Application {
	app := Application{}
	app.GrpcServer = backend.NewGrpcServer()
	return &app
}
