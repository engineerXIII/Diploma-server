package jobs

import "github.com/labstack/echo/v4"

type Handlers interface {
	Create() echo.HandlerFunc
	Status() echo.HandlerFunc
}
