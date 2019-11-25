package swf

import (
	"fmt"
	"log"
	"net/http"
)

// If S contains an embedded field T, the method sets of S and *S both include
// promoted methods with receiver T. The method set of *S also includes promoted methods with receiver *T.
// If S contains an embedded field *T, the method sets of S and *S both include
// promoted methods with receiver T or *T.
type Application struct {
	*Router
	// context *Context
}

func (app *Application) Run() {

	log.Fatal(http.ListenAndServe(":12345", app.Router))
	fmt.Printf("test.....")
}

func New() *Application {
	return &Application{
		Router: NewAPIBuilder(),
		// Router: Router{
		// 	repository: map[string]func(ctx *Context){},
		// },
	}
}
