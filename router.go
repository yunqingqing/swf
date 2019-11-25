package swf

import (
	_ "io"
	"net/http"
)

type Router struct {
	repository map[string]func(ctx *Context)
}

// 给服务注册路由
func (router *Router) Handle(method string, relativePath string, handler func(ctx *Context)) {
	router.repository[relativePath] = handler
}

// http请求入口
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.Path
	ctx := NewContext()
	ctx.r = r
	ctx.w = w

	// TODO: 支持更多的功能：rest api...
	// handler := router.repository[requestURL]
	var handler func(ctx *Context)

	if value, ok := router.repository[requestURL]; ok {
		handler = value
	} else {
		// handler notfound
	}

	// TODO: 可以在这里实现middleware

	// 处理请求
	handler(ctx)
}

func NewAPIBuilder() *Router {
	api := &Router{
		repository: map[string]func(ctx *Context){},
	}
	return api
}
