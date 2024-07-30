package go_new_ups_api_wrapper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/idoubi/goz"
)

const api_basic_url = "https://onlinetools.ups.com/"
const api_basic_url_sandbox = "https://wwwcie.ups.com/"

const ContentType = "application/x-www-form-urlencoded"
const Basic = "Basic "
const Bearer = "Bearer "
const SubVersion = "1707" // "2205"

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

// Get请求
func HttpGet(url string, headerMap, params map[string]interface{}) (string, error) {
	client := goz.NewClient()
	response, err := client.Get(url, goz.Options{
		Timeout: 3000,
		// Query:        nil,
		Headers: headerMap,
		// Cookies:      nil,
		FormParams: params,
	})

	// 请求验证
	if err != nil {
		return "", err
	}
	// 获取响应体
	body, err := response.GetBody()
	if err != nil {
		return "", err
	}

	// 判断结果
	if response.GetStatusCode() != http.StatusOK {
		respErr := &ResponseErr{}
		err = json.Unmarshal([]byte(body.String()), respErr)
		if err != nil {
			return "", err
		}
		return "", errors.New(respErr.Response.Errors[0].Message)
	}
	// 返回响应结构体
	// result := make(map[string]interface{})
	// err = json.Unmarshal([]byte(body.String()), &result)
	// if err != nil {
	// 	return nil, err
	// }
	return body.String(), nil
}

// Post 请求
func HttpPost(url string, headerMap map[string]interface{}, params interface{}) (string, error) {
	client := goz.NewClient()
	response, err := client.Post(url, goz.Options{
		Timeout: 3000,
		// Query:        nil,
		Headers: headerMap,
		// Cookies:      nil,
		// FormParams:   nil,
		JSON: params,
		// XML:          nil,
		// Multipart:    nil,
		// Proxy:        "",
		// Certificates: nil,
	})
	// 请求验证
	if err != nil {
		return "", err
	}
	// 获取响应体
	body, err := response.GetBody()
	if err != nil {
		return "", err
	}

	// 判断结果
	if response.GetStatusCode() != http.StatusOK {
		respErr := &ResponseErr{}
		err = json.Unmarshal([]byte(body.String()), respErr)
		if err != nil {
			return "", err
		}
		return "", errors.New(respErr.Response.Errors[0].Message)
	}
	// 返回响应结构体
	// result := make(map[string]interface{})
	// err = json.Unmarshal([]byte(body.String()), &result)
	// if err != nil {
	// 	return nil, err
	// }
	return body.String(), nil
}

// Delete 请求
func HttpDelete(url string, headerMap map[string]interface{}, params interface{}) (string, error) {
	client := goz.NewClient()
	response, err := client.Delete(url, goz.Options{
		Timeout: 3000,
		// Query:        nil,
		Headers: headerMap,
		// Cookies:      nil,
		// FormParams:   nil,
		JSON: params,
		// XML:          nil,
		// Multipart:    nil,
		// Proxy:        "",
		// Certificates: nil,
	})
	// 请求验证
	if err != nil {
		return "", err
	}
	// 获取响应体
	body, err := response.GetBody()
	if err != nil {
		return "", err
	}

	// 判断结果
	if response.GetStatusCode() != http.StatusOK {
		respErr := &ResponseErr{}
		err = json.Unmarshal([]byte(body.String()), respErr)
		if err != nil {
			return "", err
		}
		return "", errors.New(respErr.Response.Errors[0].Message)
	}
	// 返回响应结构体
	// result := make(map[string]interface{})
	// err = json.Unmarshal([]byte(body.String()), &result)
	// if err != nil {
	// 	return nil, err
	// }
	return body.String(), nil
}
