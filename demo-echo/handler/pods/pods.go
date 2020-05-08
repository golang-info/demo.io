package pods

import (
	"github.com/labstack/echo/v4"

	"demo.io/demo-echo/k8sClient"
	"demo.io/demo-echo/models"
	"demo.io/demo-echo/utils"
)

func GetPods(c echo.Context) error {

	n := new(models.PodsSum)
	n.Num = k8sClient.GetPodsNum()
	if err := c.Bind(n); err != nil {
		return err
	}

	//return c.JSON(http.StatusOK, n)
	return utils.NewResponse(n)
}
