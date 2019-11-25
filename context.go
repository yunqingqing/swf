package swf

import (
	"net/http"
	"html/template"
	"io"
	"log"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

//返回字符串body
func (ctx *Context)WriteString(str string){
	io.WriteString(ctx.w, str)
}

//返回html的body
func (ctx *Context)WriteHTML(path string){
	t, _ := template.ParseFiles(path)
	log.Println(t.Execute(ctx.w, nil))
}

//返回json的body
func (ctx *Context)WritejSON(str string){
}

func NewContext() *Context {
	return &Context{}
}