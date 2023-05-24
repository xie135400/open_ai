package middleware

import (
	"github.com/kataras/iris/v12"
	"net/http/httputil"
	"net/url"
)

func ProxyData(ctx iris.Context) {
	body, _ := ctx.GetBody()
	ctx.Request().ContentLength = int64(len(body))
	targetUrl, err := url.Parse("https://api.openai.com")
	if err != nil {
		ctx.Next()
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	ctx.Next()
	return
}
