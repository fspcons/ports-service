package api

import (
	"context"
	"github.com/fspcons/ports-service/src/app/api/handlers"
	"github.com/fspcons/ports-service/src/config"
	"github.com/fspcons/ports-service/src/usecases/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func StartAPI(_ context.Context, dic *dig.Container) {
	err := dic.Invoke(func(e *echo.Echo, c config.Data, logger *zap.Logger) {
		if err := e.Start(c.RestAPIAddress); err != nil {
			logger.Fatal("The API has stopped", zap.Error(err))
		}
	})
	if err != nil {
		log.Fatalln("Failed to start the API", err)
	}
}

// BuildAPI builds the port rest api server.
func BuildAPI(logger *zap.Logger, portUC port.UseCase) *echo.Echo {
	e := newServer()
	e.HTTPErrorHandler = NewHTTPErrorHandler(e, logger)
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "I'm alive!")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	handlers.AddPortHandlers(e, portUC)

	return e
}

// newServer creates a new configured REST api router
func newServer() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//TODO here I could configure additional middlewares (such as auth, CSRF, CORS configs, throttling configs etc
	// to improve server security but for time’s sake I'll just leave this comment :)

	return e
}
