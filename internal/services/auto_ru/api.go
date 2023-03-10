package auto_ru

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/stretchr/testify/mock"
)

type AutoRUApi struct {
	mock.Mock
	BaseURL string
	Headers map[string]string
}

func (api *AutoRUApi) Request(method string, url *url.URL, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url.String(), data)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Printf("Client: could not create request: %s\n", err)
		return nil, err
	}

	for k, v := range api.Headers {
		req.Header.Set(k, v)
	}

	csrf, _ := req.Cookie("_csrf_token")
	req.Header.Set("x-csrf-token", csrf.Value)

	log.SetPrefix("[INFO] ")
	log.Printf("Request: url: %s\n", req.URL)
	log.Printf("Request: method: %s\n", req.Method)
	log.Printf("Request: content-type: %s\n", req.Header.Get("content-type"))
	log.Printf("Request: content-length: %s\n", strconv.FormatInt(req.ContentLength, 10))
	log.Printf("Request: headers: %s\n", req.Header)

	http.DefaultClient.Transport = &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Printf("Client: error making http request: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.SetPrefix("[ERROR] ")
		log.Printf("Client: could not read request body: %s\n", err)
		return nil, err
	}

	return body, nil
}

func (api *AutoRUApi) BuildURL(path string) *url.URL {
	baseUrl, _ := url.Parse(fmt.Sprintf("%s%s", api.BaseURL, path))
	return baseUrl
}

func (api *AutoRUApi) Post(path *url.URL, data []byte) ([]byte, error) {
	return api.Request(http.MethodPost, path, bytes.NewBuffer(data))
}

func (api *AutoRUApi) Get(path *url.URL) ([]byte, error) {
	return api.Request(http.MethodGet, path, nil)
}
