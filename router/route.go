package router

import (
	"ecologyServer/handler"
	"ecologyServer/middleware"
	"fmt"
	"github.com/henrylee2cn/faygo"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	fmt.Printf("start to add route")
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
		frame.NewNamedAPI("Save result", "POST", "/save/scanner/result", handler.SaveResult),
	)
	fmt.Printf("add route success %v\n", &handler.Test{})
}
