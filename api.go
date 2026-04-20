package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

type API struct {
	*echo.Echo
}

func NewAPI() *API {
	a := &API{
		Echo: echo.New(),
	}

	pprof.Register(a.Echo)

	a.POST("/process", a.handler)

	return a
}

type userRequest struct {
	Names []string `json:"names"`
}

func (a *API) handler(c echo.Context) error {
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	var results []string
	// Компиляция регулярки внутри хендлера (тратит CPU)
	re := regexp.MustCompile(`^[A-Z]`)

	for _, name := range req.Names {
		if re.MatchString(name) {
			// fmt.Sprintf создает лишние аллокации в цикле
			results = append(results, fmt.Sprintf("Hello, %s!", name))
		}
	}

	return c.JSON(http.StatusOK, results)
}
