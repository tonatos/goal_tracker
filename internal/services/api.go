package services

import "net/url"

type Api interface {
	BuildURL(path string) *url.URL
	Post(path *url.URL, data []byte) ([]byte, error)
	Get(path *url.URL) ([]byte, error)
}
