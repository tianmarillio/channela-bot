package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type channel struct {
	YoutubeChannelID string `json:"youtubeChannelId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
}

func SearchChannel(keyword string) []channel {
	client := &http.Client{}
	url := fmt.Sprintf("%s/search", os.Getenv("YOUTUBE_API_BASE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("part", "snippet")
	q.Add("type", "channel")
	q.Add("q", keyword)
	q.Add("maxResults", strconv.Itoa(10))
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if err != nil {
		// TODO: handle error
		// return nil
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	bodyJson := string(body)

	var searchResult struct {
		Items []struct {
			ID struct {
				ChannelID string `json:"channelId"`
			} `json:"id"`
			Snippet struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"snippet"`
		} `json:"items"`
	}
	json.Unmarshal([]byte(bodyJson), &searchResult)

	var channels []channel
	for _, item := range searchResult.Items {
		channels = append(channels, channel{
			YoutubeChannelID: item.ID.ChannelID,
			Title:            item.Snippet.Title,
			Description:      item.Snippet.Description,
		})
	}

	return channels
}
