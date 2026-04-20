package main

import (
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

// Оптимизация 1: Компилируем регулярку один раз при старте
var nameRegexp = regexp.MustCompile(`^[A-Z]`)

func (a *API) handler(c echo.Context) error {
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// Оптимизация 2: Заранее выделяем память под слайс,
	// чтобы избежать переаллокаций в цикле (пока примерно)
	results := make([]string, 0, len(req.Names))

	for _, name := range req.Names {
		if nameRegexp.MatchString(name) {
			// Оптимизация 3: Простая конкатенация вместо fmt.Sprintf
			results = append(results, "Hello, "+name+"!")
		}
	}

	return c.JSON(http.StatusOK, results)
}
