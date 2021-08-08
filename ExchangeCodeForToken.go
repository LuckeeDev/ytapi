package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const googleTokenEndpoint = "https://oauth2.googleapis.com/token"

func ExchangeCodeForToken(code string, clientID string, clientSecret string, redirectURI string) (token string) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectURI)

	encodedData := data.Encode()

	client := &http.Client{}
	r, err := http.NewRequest("POST", googleTokenEndpoint, strings.NewReader(encodedData))

	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var bodyJSON AccessTokenResponse

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(body, &bodyJSON)

	if jsonErr != nil {
		log.Default().Println("JSON err")
		log.Fatal(jsonErr)
	}

	return bodyJSON.AccessToken
}
