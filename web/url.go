package web

import (
	"net/url"
)

type URLData struct {
	Host    string
	Path    string
	VideoID string
}

// https://gobyexample.com/url-parsing
func ParseURLData(s string) (*URLData, error) {
	var v string
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	q, _ := url.ParseQuery(u.RawQuery)
	if len(q["v"]) > 0 {
		v = q["v"][0]
	}

	return &URLData{
		Host:    u.Host,
		Path:    u.Path,
		VideoID: v,
	}, nil
}
