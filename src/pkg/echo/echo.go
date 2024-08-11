package echo

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	return e
}
