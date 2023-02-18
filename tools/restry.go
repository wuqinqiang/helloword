package tools

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	Resty *resty.Client

	defaultClient = &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			MaxIdleConns:      15,
			IdleConnTimeout:   90 * time.Second,
		},
	}
)

func init() {
	Resty = resty.NewWithClient(defaultClient)
}
