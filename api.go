package main

import (
	"net/http"
	"sync"

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

// Оптимизация 4: Пул для повторного использования слайсов
var resultsPool = sync.Pool{
	New: func() interface{} {
		// Создаем слайс с запасом на 1000 элементов
		return make([]string, 0, 1000)
	},
}

func (a *API) handler(c echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// 1. Получаем "чистую" тарелку из стопки (пула)
	results := resultsPool.Get().([]string)

	// 2. Сбрасываем длину до 0, но оставляем емкость (capacity)
	// Это критически важно: память не выделяется заново!
	results = results[:0]

	for _, name := range req.Names {
		if len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z' {
			results = append(results, "Hello, "+name+"!")
		}
	}

	// Отправляем ответ
	err := c.JSON(http.StatusOK, results)

	// 3. Возвращаем тарелку обратно в стопку (пул)
	// Теперь другой запрос сможет забрать этот же слайс
	resultsPool.Put(results)

	return err
}
