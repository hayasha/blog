package main

import (
	"github.com/hayasha/blog/controllers"
	"github.com/hayasha/blog/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToMysql()
}

func main() {
	routers := controllers.IndexRouters()
	routers.Run()
}
