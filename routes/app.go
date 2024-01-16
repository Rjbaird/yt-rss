package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/bairrya/yt-rss/database"
	"github.com/bairrya/yt-rss/web"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "base", nil)
}

func PostFeed(c *gin.Context) {
	hxr := c.GetHeader("HX-Request")
	if hxr != "true" {
		c.Redirect(http.StatusSeeOther, "/")
	}

	// get the url from the form
	formUrl := c.PostForm("video-url")
	// parse the url
	url, err := web.ParseURLData(formUrl)
	// if not a youtube url, return error
	if err != nil || !strings.Contains(url.Host, "youtube.com") {
		c.String(http.StatusOK, "Not a youtube url")
		return
	}
	// if not a video url, return error
	if url.VideoID == "" {
		c.String(http.StatusOK, "Not a video url")
		return
	}

	// check if video id and channel data is already in the cache
	ctx := context.Background()
	db, err := database.RedisConnect()
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusOK, err.Error())
		return
	}
	defer db.Close()

	// TODO: check redis cache for video id

	// get channel data from youtube api
	yt, err := web.YouTubeClient()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	v, err := web.FetchVideoData(url.VideoID, yt)
	// if no data, return error
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	channel, err := web.FetchChannelDataByID(v.Items[0].Snippet.ChannelId, yt)

	// if no data, return error
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	// set channel data in database
	data, err := database.SetChannel(ctx, channel, db)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	// cache video id
	// set video data in database
	_ = database.SetVideo(ctx, url.VideoID, data, db)
	// return channel data
	c.HTML(http.StatusOK, "video-card", data)
}

func PostPlaylist(c *gin.Context) {
	c.HTML(http.StatusOK, "playlist", nil)
}

func PostSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "search", nil)
}
