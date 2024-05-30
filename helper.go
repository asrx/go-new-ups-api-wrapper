package go_new_ups_api_wrapper

import (
	"strings"

	"github.com/idoubi/goz"
)

func getRequestUrl(targetUrl string, debug bool) string {
	basicUrl := ""
	if debug {
		basicUrl = api_basic_url_sandbox
	} else {
		basicUrl = api_basic_url
	}
	return strings.TrimSuffix(basicUrl, "/") + "/" + strings.TrimPrefix(targetUrl, "/")
}

func NewHttpClient(method, url string) {

	client := goz.NewClient()
	client.Request(method, url, goz.Options{
		Timeout:    3000,
		Headers:    nil,
		FormParams: nil,
	})
}
