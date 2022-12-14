package main

import (
	"fmt"
	"net/http"

	"monorepo_test/libs/hello"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/two/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = e.Start(":8082")
	fmt.Println("request recieved to service two")
}
