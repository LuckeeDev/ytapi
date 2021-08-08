package functions

import (
	"fmt"
	"strings"
)

func GetAuthURL(clientID string, redirectURI string, scopes []string) (authURL string) {
	const googleAuthEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	
	scope := strings.Join(scopes, " ")

	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s", googleAuthEndpoint, clientID, redirectURI, scope)
}
