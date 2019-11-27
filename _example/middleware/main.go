package main

import (
	"fmt"

	"github.com/yunqingqing/swf"
)

func main() {
	app := swf.New()

	app.Handle("GET", "/helloworld", func(ctx *swf.Context) {
		ctx.WriteString("hello, word")
	}, func(ctx *swf.Context) {
		fmt.Printf("prev request")
	})

	app.Run()
}
