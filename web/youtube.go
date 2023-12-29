package web

import (
	"github.com/bairrya/yt-rss/config"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YouTubeClient() (*youtube.Service, error) {
	ctx := context.Background()
	config, err := config.ENV()
	if err != nil {
		return nil, err
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(config.YOUTUBE_API_KEY))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func FetchVideoData(id string, yt *youtube.Service) (*youtube.VideoListResponse, error) {
	// https://developers.google.com/youtube/v3/docs/videos/list
	v, _ := yt.Videos.List([]string{"snippet"}).Id(id).Do()
	// TODO: check for errors
	return v, nil
}

func FetchChannelDataByID(id string, yt *youtube.Service) (*youtube.ChannelListResponse, error) {
	// https://developers.google.com/youtube/v3/docs/channels/list
	c, _ := yt.Channels.List([]string{"snippet"}).Id(id).Do()
	// TODO: check for errors
	return c, nil
}

func FetchPlaylistVideosByID(id string, yt *youtube.Service) (*youtube.PlaylistItemListResponse, error) {
	// https://developers.google.com/youtube/v3/docs/playlists/list
	p, _ := yt.PlaylistItems.List([]string{"snippet"}).PlaylistId(id).Do()
	// TODO: check for errors
	return p, nil
}

func FetchBulkVideoData(ids []string, yt *youtube.Service) (*youtube.VideoListResponse, error) {
	// https://developers.google.com/youtube/v3/docs/videos/list
	v, _ := yt.Videos.List([]string{"snippet"}).Id(ids...).Do()
	// TODO: check for errors
	return v, nil
}