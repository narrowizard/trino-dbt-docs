package services

import (
	"crypto/tls"
	"net/http"
)

var CustomerHttpClient *http.Client

func init() {
	CustomerHttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
