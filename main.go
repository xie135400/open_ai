package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"open_ai/middleware"
)

func main() {
	app := iris.New()
	app.OnErrorCode(http.StatusNotFound, middleware.ProxyData)
	err := app.Run(
		iris.Addr(":8089"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithoutBodyConsumptionOnUnmarshal,
	)
	if err != nil {
		panic(err)
	}
	return
}
