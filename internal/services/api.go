package services

import "net/url"

type Api interface {
	BuildURL(urlName string) *url.URL
	Post(url *url.URL, data []byte) ([]byte, error)
	Get(url *url.URL) ([]byte, error)
}
