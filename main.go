package main

import (
	"fmt"
	"os"
	"ytwl/functions"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// const youtubePlaylistsEndpoint = "https://www.googleapis.com/youtube/v3/playlists"
const redirectURI = "http://localhost:3000/redirect"

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Println("There was an error loading the .env file.")
	}

	r := gin.Default()

	r.GET("/login", func(ctx *gin.Context) {
		clientID := os.Getenv("GOOGLE_CLIENT_ID")
		redirectURI := "http://localhost:3000/redirect"
		scopes := []string{"https://www.googleapis.com/auth/youtube"}

		authURL := functions.GetAuthURL(clientID, redirectURI, scopes)

		ctx.Redirect(302, authURL)
	})

	r.GET("/redirect", func(ctx *gin.Context) {
		code := ctx.Query("code")
		clientID := os.Getenv("GOOGLE_CLIENT_ID")
		clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

		token := functions.ExchangeCodeForToken(code, clientID, clientSecret, redirectURI)

		ctx.JSON(200, gin.H{
			"code":  code,
			"token": token,
		})
	})

	// req, err := http.NewRequest("GET", youtubePlaylistsEndpoint, nil)

	// if err != nil {
	// 	fmt.Println("There was an error")
	// }

	// token := "token"

	// req.Header.Set("Authorization", token)

	r.Run(":3000")
}
