package handlers

import (
	"github.com/fspcons/ports-service/src/domain"
	"github.com/fspcons/ports-service/src/usecases/port"
	"github.com/fspcons/ports-service/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	MsgInvalidFormat = "invalid request body format"
	MsgInvalidData   = "invalid request data"
)

func RegisterPortHandlers(e *echo.Echo, uc port.UseCase) {
	e.POST("/v1/port", PostPort(uc))
	//using Patch here because on the useCase I'm checking for fields to update the record
	e.PATCH("/v1/port/:id", UpdatePort(uc))
}

// PostPort
// @Router   /v1/port [post]
// @Summary  Creates a new port
// @Tags     port
// @Accept   json
// @Param    data  body      domain.Port  true  "Port data"
// @Success  201   {object}  domain.Port
// @Produce  json
// @Failure  400  "invalid request body format"
// @Failure  422  "informed record already exists"
func PostPort(uc port.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(domain.Port)
		if err := c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, MsgInvalidFormat)
		}
		err := uc.Create(c.Request().Context(), p)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, p)
	}
}

// UpdatePort
// @Router   /v1/port/{id} [patch]
// @Summary  Updates a port by id
// @Tags     port
// @Accept   json
// @Param    id  path      string       true  "ID of the port to update"
// @Param    data    body      port.Update  true  "Port data"
// @Success  200     {object}  domain.Port
// @Produce  json
// @Failure  400  "invalid request data"
// @Failure  404  "no records were found"
func UpdatePort(uc port.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if utils.IsEmpty(id) {
			return echo.NewHTTPError(http.StatusBadRequest, MsgInvalidData)
		}
		u := new(port.Update)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		p, err := uc.Update(c.Request().Context(), id, *u)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &p)
	}
}
