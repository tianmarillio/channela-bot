package scripts

import (
	"fmt"
	"time"

	"github.com/tianmarillio/channela-backend/src/api/twitterapi"
)

// FOR TEST PURPOSES
func CreateTimestampTweet() {
	currentTime := time.Now()
	currentTimeStr := currentTime.Format(time.RFC3339)
	tweet := fmt.Sprintf("Time: %s", currentTimeStr)

	fmt.Println("tweeting>", tweet)
	twitterapi.CreateTweet(tweet)
}
