package go_new_ups_api_wrapper

import (
	"fmt"
	"testing"
)

func TestGetURL(t *testing.T) {
	url := getRequestUrl("/api/ok", true)
	fmt.Println(url)
}

func Test_OAuthToken(t *testing.T) {
	clientId := "yourClientID"
	clientSecret := "yourSecret"
	token, err := OauthToken(clientId, clientSecret, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", token)
}
