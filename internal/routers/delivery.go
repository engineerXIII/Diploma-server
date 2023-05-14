package routers

import "github.com/labstack/echo/v4"

// Routers HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	//Update() echo.HandlerFunc
	//Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Status() echo.HandlerFunc
}
