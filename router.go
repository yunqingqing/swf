package swf

import (
	"net/http"
)

type HttpHandler func(ctx *Context)

type Router struct {
	repository map[string]Route
}

type Route struct {
	handler     HttpHandler
	middlewares []HttpHandler
}

// 给服务注册路由
func (router *Router) Handle(method string, relativePath string, handler HttpHandler, middlewares ...HttpHandler) {
	route := Route{
		handler:     handler,
		middlewares: middlewares,
	}
	router.repository[relativePath] = route
}

// http请求入口
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.Path
	ctx := NewContext()
	ctx.r = r
	ctx.w = w

	// TODO: 支持更多的功能：rest api...
	// handler := router.repository[requestURL]
	var handler HttpHandler
	var middlewares []HttpHandler

	if route, ok := router.repository[requestURL]; ok {
		handler = route.handler
		middlewares = route.middlewares
	} else {
		// handler notfound
	}

	// TODO: 可以在这里实现middleware
	for _, middleware := range middlewares {
		middleware(ctx)
	}

	// 处理请求
	handler(ctx)
}

func NewAPIBuilder() *Router {
	api := &Router{
		repository: map[string]Route{},
	}
	return api
}
