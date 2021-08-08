package functions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ListPlaylists(token string) (playlists map[string]interface{}) {
	const youtubePlaylistsEndpoint = "https://www.googleapis.com/youtube/v3/playlists?part=snippet&mine=true&pageToken=CAUQAA"

	client := &http.Client{}
	r, err := http.NewRequest("GET", youtubePlaylistsEndpoint, nil)

	r.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(r)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(err)
	}

	var jsonBody map[string]interface{}

	json.Unmarshal(body, &jsonBody)

	return jsonBody
}
