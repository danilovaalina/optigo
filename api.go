package main

import (
	"net/http"

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
	// Мы не выделяем заранее 1000, чтобы не пугать память

	for _, name := range req.Names {
		// Оптимизация 1: Вместо regexp используем быструю проверку символа
		if len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z' {
			// Оптимизация 2: Простая склейка (в Go для 2-3 элементов она ок,
			// но так мы убираем лишнюю логику regexp)
			results = append(results, "Hello, "+name+"!")
		}
	}

	return c.JSON(http.StatusOK, results)
}
