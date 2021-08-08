package main

import (
	"fmt"
	"os"
	"ytwl/functions"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const redirectURI = "http://localhost:3000/redirect"

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Println("There was an error loading the .env file.")
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/login")
	})

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

		playlists := functions.ListPlaylists(token)

		ctx.JSON(200, gin.H{
			"code":      code,
			"token":     token,
			"playlists": playlists,
		})
	})

	r.Run(":3000")
}
