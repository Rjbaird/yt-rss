package database

type Channel struct {
	ChannelID       string `json:"channel_id" redis:"channel_id"`
	Handle          string `json:"handle" redis:"handle"`
	Title           string `json:"title" redis:"title"`
	Description     string `json:"description" redis:"description"`
	ThumbnailUrl    string `json:"thumbnail_url" redis:"thumbnail_url"`
	ThumbnailWidth  string `json:"thumbnail_width" redis:"thumbnail_width"`
	ThumbnailHeight string `json:"thumbnail_height" redis:"thumbnail_height"`
	DateAdded       string `json:"date_added" redis:"date_added"` // unix timestamp
}
