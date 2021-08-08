package main

import (
	"fmt"
	"strings"
)

const googleAuthEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"

func GetAuthURL(clientID string, redirectURI string, scopes []string) (authURL string) {
	scope := strings.Join(scopes, " ")

	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s", googleAuthEndpoint, clientID, redirectURI, scope)
	// return googleAuthEndpoint + "?client_id=" + clientID + "&redirect_uri=" + redirectURI + "&response_type=code&scope=" + scope + "&access_type=offline"
}
