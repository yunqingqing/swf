package main

import(
	"github.com/yunqingqing/swf"
)

func main() {
	app := swf.New()

	app.Handle("GET", "/helloworld", func(ctx *swf.Context) {
		ctx.WriteString("hello, word")
	})
	
	app.Run()
}