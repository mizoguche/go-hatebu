package hatebu

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/garyburd/go-oauth/oauth"
)

func Add(url string, comment string) {
	// http://nasust.hatenablog.com/entry/2016/12/17/111125
	oauthClient := &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  "XXXXXXXXXXXXXX",
			Secret: "XXXXXXXXXXXXXX",
		},
		TemporaryCredentialRequestURI: "https://www.hatena.com/oauth/initiate",
		ResourceOwnerAuthorizationURI: "https://www.hatena.com/oauth/authorize",
		TokenRequestURI:               "https://www.hatena.com/oauth/token",
	}

	accessToken := oauth.Credentials{
		Token:  "XXXXXXXXXXXXXXX",
		Secret: "XXXXXXXXXXXXXXX",
	}

	response, err := oauthClient.Get(nil, &accessToken, "https://blog.hatena.ne.jp/nasust/nasust.hatenablog.com/atom", nil)

	if err != nil {
		log.Fatal("Get Err: ", err)
		panic(-1)
	}

	fmt.Println("Status: ", response.Status)

	bodyData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Read Err:", err)
		panic(-1)
	}

	bodyStr := string(bodyData)

	fmt.Println(bodyStr)
}
