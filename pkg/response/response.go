package response

import "github.com/labstack/echo/v4"

// Response schema
type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func JSON(c echo.Context, status int, data interface{}) error {
	return c.JSON(status, Response{
		Data: data,
	})
}
