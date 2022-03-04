package main

import (
	"dtm-log-view/config"
	"dtm-log-view/controller"
	"flag"
	"github.com/gin-gonic/gin"
)

var confFile = flag.String("c", "./config.yml", "Path to the server configuration file.")

func main() {
	flag.Parse()

	config.MustLoadConfig(*confFile)

	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// home page
	r.StaticFile("/", "./templates/index.html")
	r.StaticFile("/index.html", "./templates/index.html")

	// file list api
	r.GET("/file/get_list", func(c *gin.Context) {
		controller.GetList(c)
	})

	// file content api
	r.POST("/file/get_content", func(c *gin.Context) {
		controller.GetContent(c)
	})
	return r
}
