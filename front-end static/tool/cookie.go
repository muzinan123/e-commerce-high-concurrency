package tool

import (
	"net/http"

	"github.com/kataras/iris"
)

func GlobalCookie(ctx iris.Context, name string, value string) {
	ctx.SetCookie(&http.Cookie{Name: name, Value: value, Path: "/"})
}
