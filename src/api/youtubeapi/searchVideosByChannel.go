package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type video struct {
	VideoID string
	Title   string
}

func SearchVideosByChannel(channelId string, publishedAfter string) []video {
	client := &http.Client{}

	url := fmt.Sprintf("%s/search", os.Getenv("YOUTUBE_API_BASE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("YOUTUBE_API_KEY"))
	q.Add("channelId", channelId)
	q.Add("part", "snippet")
	q.Add("type", "video")
	q.Add("order", "date")
	q.Add("publishedAfter", publishedAfter)
	// q.Add("maxResults", 5)
	req.URL.RawQuery = q.Encode()

	response, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	bodyJson := string(body)

	var searchResult struct {
		Items []struct {
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
			ID struct {
				VideoId string `json:"videoId"`
			} `json:"id"`
		} `json:"items"`
	}
	json.Unmarshal([]byte(bodyJson), &searchResult)

	var videos []video
	for _, item := range searchResult.Items {
		videos = append(videos, video{
			VideoID: item.ID.VideoId,
			Title:   item.Snippet.Title,
		})
	}

	return videos
}
