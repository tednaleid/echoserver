package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))
	e.Use(middleware.Recover())

	e.GET("/:id", getResult)
	e.POST("/:id", postResult)

	e.Logger.Fatal(e.Start(":1323"))
}

func getResult(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func postResult(c echo.Context) error {
	if c.Request().Body != nil {
		reqBody, _ := ioutil.ReadAll(c.Request().Body)
		reqString := string(reqBody)

		if len(reqString) > 0 {
			println(reqString)
			return c.String(http.StatusOK, reqString)
		}
	}

	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
