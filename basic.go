package go_new_ups_api_wrapper

import (
	"strings"

	"github.com/idoubi/goz"
)

const api_basic_url = "https://onlinetools.ups.com/"
const api_basic_url_sandbox = "https://wwwcie.ups.com/"

const ContentType = "application/x-www-form-urlencoded"
const Basic = "Basic "
const Bearer = "Bearer "

func getHeaders(authorization string) map[string]interface{} {
	return map[string]interface{}{
		"Content-Type":  ContentType,
		"Authorization": authorization,
	}
}

// 错误响应
type ResponseErr struct {
	Response struct {
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"response"`
}

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
