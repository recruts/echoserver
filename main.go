package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

// Handler обрабатывает все GET и POST запросы и выводит информацию о пользователе и переданные переменные
func Handler(c echo.Context) error {
	params := make(map[string]interface{})
	headers := c.Request().Header
	for key, vals := range c.QueryParams() {
		if len(vals) > 1 {
			params[key] = vals
		} else {
			params[key] = vals[0]
		}
	}
	formParams, err := c.FormParams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	for key, vals := range formParams {
		if len(vals) > 1 {
			params[key] = vals
		} else {
			params[key] = vals[0]
		}
	}
	responseData := map[string]interface{}{
		"ip":         c.RealIP(),
		"user_agent": headers.Get("User-Agent"),
		"header":     headers,
		"params":     params,
	}
	return c.JSON(http.StatusOK, responseData)
}
func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Any("/echo", Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Start(":" + port)
}
