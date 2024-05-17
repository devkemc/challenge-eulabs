package response

import "github.com/labstack/echo/v4"

// Response schema
type Response struct {
	Data interface{} `json:"data"`
}

func JSON(c echo.Context, status int, data interface{}) error {
	if data == nil {
		c.Response().WriteHeader(status)
		return nil
	}
	return c.JSON(status, Response{
		Data: data,
	})
}
