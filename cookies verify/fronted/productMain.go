package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.StaticWeb("/public", "./fronted/web/public")
	app.StaticWeb("/html", "./fronted/web/htmlProductShow")

	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
