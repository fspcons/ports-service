package api

import (
	"errors"
	"github.com/fspcons/ports-service/src/domain"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func NewHTTPErrorHandler(e *echo.Echo, logger *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if err != nil {
			var status int
			switch {
			case errors.Is(err, domain.ErrNoRecords):
				status = http.StatusNotFound
			case errors.Is(err, domain.ErrRecordAlreadyExists):
				status = http.StatusUnprocessableEntity
			case errors.Is(err, domain.ErrInvalidPort):
				status = http.StatusBadRequest
			default:
				logger.Sugar().Error(err)
				err = echo.NewHTTPError(http.StatusInternalServerError, "internal error")
				status = http.StatusInternalServerError
			}
			c.Response().WriteHeader(status)
		}
		e.DefaultHTTPErrorHandler(err, c)
	}
}
