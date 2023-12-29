package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/api/youtube/v3"
)

func SetChannel(ctx context.Context, ytChannel *youtube.ChannelListResponse, db *redis.Client) (*Channel, error) {
	now := time.Now().Unix()
	channel := Channel{
		ChannelID:   ytChannel.Items[0].Id,
		Handle:      ytChannel.Items[0].Snippet.CustomUrl,
		Title:       ytChannel.Items[0].Snippet.Title,
		Description: ytChannel.Items[0].Snippet.Description,
		ThumbnailUrl: ytChannel.Items[0].Snippet.Thumbnails.
			Default.Url,
		ThumbnailWidth: fmt.Sprintf("%d", ytChannel.Items[0].Snippet.Thumbnails.
			Default.Width),
		ThumbnailHeight: fmt.Sprintf("%d", ytChannel.Items[0].Snippet.Thumbnails.
			Default.Height),
		DateAdded: fmt.Sprintf("%d", now),
	}

	key := fmt.Sprintf("channel:%s", channel.ChannelID)
	err := db.HSet(ctx, key, channel).Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &channel, nil
}

func SetVideo(ctx context.Context, videoID string, channel *Channel, db *redis.Client) error {
	key := fmt.Sprintf("video:%s", videoID)
	err := db.HMSet(ctx, key, channel).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
