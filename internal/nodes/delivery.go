package nodes

import "github.com/labstack/echo/v4"

// Nodes HTTP Handlers interface
type Handlers interface {
	GetList() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	Status() echo.HandlerFunc
}
