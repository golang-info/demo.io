package postData

import (
	"net/http"
	"github.com/labstack/echo/v4"

)

func PostData(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, m)
}