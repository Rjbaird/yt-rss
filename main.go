package main

import (
	"fmt"

	"github.com/bairrya/yt-rss/config"
	"github.com/bairrya/yt-rss/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.ENV()
	if err != nil {
		panic(err)
	}
	port := fmt.Sprintf(":%s", config.PORT)
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*.html")
	r.GET("/", routes.GetHome)
	r.POST("/video", routes.PostFeed)
	r.POST("/playlist", routes.PostPlaylist)
	r.POST("/search", routes.PostSearch)

	r.Run(port)
}
