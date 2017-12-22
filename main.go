package main

import (
	"ecologyServer/router"
	"ecologyServer/utils/mongo"
	"github.com/henrylee2cn/faygo"
)

func main() {
	mongo.Init()
	router.Route(faygo.New("ecologyServer"))
	faygo.Run()
}
