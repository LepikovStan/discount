package crawler

import (
	"io"
	"net/http"
	"errors"
	"fmt"
)

type crawler struct {

}

func (c *crawler) Crawl(url string) (io.Reader, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return resp.Body, nil
	}

	err = errors.New(fmt.Sprintf("Wrong response status: %d for %s", resp.StatusCode, url))
	return nil, err
}

func New() *crawler {
	return new(crawler)
}
