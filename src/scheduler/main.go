package scheduler

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/tianmarillio/channela-backend/src/scripts"
)

func Start() {
	if os.Getenv("RUN_SCHEDULER") != "true" {
		return
	}

	s := gocron.NewScheduler(time.UTC)

	// SCHEDULED SCRIPTS
	tweetIntervalMinuteStr := os.Getenv("TWEET_INTERVAL_MINUTE")
	tweetIntervalMinuteInt, _ := strconv.Atoi(tweetIntervalMinuteStr)
	// s.Cron(fmt.Sprintf("*/%s * * * *", tweetIntervalMinuteStr)).Do(func() {
	s.Cron("0 */4 * * *").Do(func() {
		fmt.Println("START: Tweeting latest videos")
		scripts.TweetLatestVideos(tweetIntervalMinuteInt)
		fmt.Println("DONE: Tweeting latest videos")
	})

	s.StartAsync()
}
