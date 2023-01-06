package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type findChannelResponse struct {
	Title     string
	CustomUrl string
}

func FindChannel(channelId string) findChannelResponse {
	client := &http.Client{}

	url := fmt.Sprintf("%s/channels", os.Getenv("YOUTUBE_API_BASE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("part", "snippet")
	q.Add("id", channelId)
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if err != nil {
		// TODO: handle error
		// return nil
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	bodyJson := string(body)

	var channelsResult struct {
		Items []struct {
			Snippet struct {
				Title     string `json:"title"`
				CustomUrl string `json:"customUrl"`
			} `json:"snippet"`
		} `json:"items"`
	}
	json.Unmarshal([]byte(bodyJson), &channelsResult)

	return findChannelResponse{
		Title:     channelsResult.Items[0].Snippet.Title,
		CustomUrl: channelsResult.Items[0].Snippet.CustomUrl,
	}
}
