package go_new_ups_api_wrapper

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
