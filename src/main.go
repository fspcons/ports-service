package main

import (
	"context"
	"github.com/fspcons/ports-service/src/app"
	"github.com/fspcons/ports-service/src/app/api"
	"github.com/fspcons/ports-service/src/config"
	_ "github.com/fspcons/ports-service/src/docs"
	"github.com/fspcons/ports-service/src/gateway/file"
	"github.com/fspcons/ports-service/src/gateway/ports"
	"github.com/fspcons/ports-service/src/usecases/port"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

// @title        Ports REST API
// @version      1.0
// @description  API to manage Ports data.
func main() {
	ctx := context.Background()
	di := buildDIC(ctx)
	apps := []app.Start{api.StartAPI}

	container := app.NewContainer(ctx, di, apps)
	container.Run()
}

// mustProvide ensure dependencies can be provided, otherwise panic
func mustProvide(di *dig.Container, constructor any, opts ...dig.ProvideOption) {
	if err := di.Provide(constructor, opts...); err != nil {
		panic(err)
	}
}

// buildDIC build dependency injection container
func buildDIC(ctx context.Context) *dig.Container {
	di := dig.New()
	logger := config.MustNewLogger("ports-service")

	//config
	mustProvide(di, func() context.Context { return ctx })
	mustProvide(di, func() *zap.Logger { return logger })
	mustProvide(di, config.ReadFromEnv)

	//gateways
	mustProvide(di, ports.NewInMemoryGateway)
	mustProvide(di, file.NewFileGateway)
	//useCases
	mustProvide(di, port.NewUseCase)
	//API
	mustProvide(di, api.BuildAPI)

	return di
}
