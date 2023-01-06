package scripts

import (
	"fmt"
	"time"

	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/api/twitterapi"
	"github.com/tianmarillio/channela-backend/src/api/youtubeapi"
	"github.com/tianmarillio/channela-backend/src/models"
)

func TweetLatestVideos(interval int) {
	publishedAfter := time.Now().
		Add(-time.Minute * time.Duration(interval)).
		Format(time.RFC3339)

	var channels []*models.Channel
	result := config.DB.Find(&channels)
	if result.Error != nil {
		return
	}

	for _, channel := range channels {
		videos := youtubeapi.SearchVideosByChannel(channel.YoutubeChannelId, publishedAfter)
		// fmt.Println(videos)
		for _, video := range videos {
			videoUrl := fmt.Sprintf("youtube.com/watch?v=%s",
				video.VideoID,
			)
			channelUrl := fmt.Sprintf("youtube.com/%s",
				channel.CustomUrl,
			)
			message := fmt.Sprintf("%s | %s %s",
				video.Title,
				channelUrl,
				videoUrl,
			)

			time.AfterFunc(1*time.Second, func() {
				twitterapi.CreateTweet(message)
			})

		}
	}
}
