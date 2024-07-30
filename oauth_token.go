package go_new_ups_api_wrapper

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/idoubi/goz"
)

const grantType = "client_credentials"
const oauthTokenUrl = "/security/v1/oauth/token"

// 签名 clientId:clientSecret
func encodedClientIdAndSecret(clientId, clientSecret string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	return encoded
}

// Header Authorization Basic请求参数
func authorizationBasic(clientId, clientSecret string) string {
	return Basic + encodedClientIdAndSecret(clientId, clientSecret)
}

// 获取认证Token
func OauthToken(clientId, clientSecret string, debug bool) (respAuthToken *ResponseAuthToken, err error) {
	client := goz.NewClient()
	response, err := client.Post(getRequestUrl(oauthTokenUrl, debug), goz.Options{
		Timeout: 3000,
		Headers: getHeaders(authorizationBasic(clientId, clientSecret)),
		FormParams: map[string]interface{}{
			"grant_type": grantType,
		},
	})
	// 请求验证
	if err != nil {
		return nil, err
	}
	// 获取响应体
	body, err := response.GetBody()
	if err != nil {
		return nil, err
	}

	// 判断结果
	if response.GetStatusCode() != http.StatusOK {
		respErr := &ResponseErr{}
		err = json.Unmarshal([]byte(body.String()), respErr)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(respErr.Response.Errors[0].Message)
	}

	// 返回响应结构体
	resp := &ResponseAuthToken{}
	err = json.Unmarshal([]byte(body.String()), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 认证响应体
type ResponseAuthToken struct {
	TokenType   string `json:"token_type" doc:"类型"`
	IssuedAt    string `json:"issued_at" doc:"签发时间"`
	ClientId    string `json:"client_id" doc:"客户标识"`
	AccessToken string `json:"access_token" doc:"Token内容"`
	ExpiresIn   string `json:"expires_in" doc:"过期时间，以秒为单位（预计4小时）"`
	Status      string `json:"status"`
}
