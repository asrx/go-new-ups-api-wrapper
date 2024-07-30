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
	tk, err := OauthToken(clientId, clientSecret, debug)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tk.AccessToken)
}
