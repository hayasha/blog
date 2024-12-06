package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hayasha/blog/controllers"
	"github.com/hayasha/blog/initializers"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToMysql()
	initializers.ConfigureOpenTelemetry()
}

func main() {
	defer initializers.ShutdownOpenTelemetry()

	ginEngine := gin.Default()

	ginEngine.Use(otelgin.Middleware("hayasha-blog"))

	controllers.IndexRouters(ginEngine)

	ginEngine.Run()
}
