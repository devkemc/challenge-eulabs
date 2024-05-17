package response

import (
	"github.com/devkemc/challenge-eulabs/pkg/config"
	"github.com/labstack/echo/v4"
)

func Error(c echo.Context, status int, err error, message string) error {
	errorRes := map[string]interface{}{
		"message": message,
	}

	cfg := config.GetConfig()

	if cfg.Environment != config.ProductionEnv {
		errorRes["debug"] = err.Error()
	}

	return c.JSON(status, errorRes)

}
