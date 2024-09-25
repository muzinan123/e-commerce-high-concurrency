package main

import (
	"context"
	"imooc-product/backend/web/controllers"
	"imooc-product/common"
	"imooc-product/repositories"
	"imooc-product/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/opentracing/opentracing-go/log"
)

func main() {

	app := iris.New()

	app.Logger().SetLevel("debug")

	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	app.StaticWeb("/assets", "./backend/web/assets")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productRepository := repositories.NewProductManager("product", db)
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
