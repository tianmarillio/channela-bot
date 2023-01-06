package twitterapi

import (
	"bytes"
	"fmt"
	"os"

	"github.com/dghubble/oauth1"
)

func CreateTweet(message string) {
	baseUrl := os.Getenv("TWITTER_API_BASE_URL")
	url := fmt.Sprintf("%s/tweets", baseUrl)

	consumerKey := os.Getenv("TWITTER_API_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_API_CONSUMER_SECRET")
	tokenKey := os.Getenv("TWITTER_API_TOKEN_KEY")
	tokenSecret := os.Getenv("TWITTER_API_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(tokenKey, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	json := fmt.Sprintf(`{"text":"%s"}`, message)
	jsonStr := []byte(json)

	resp, err := httpClient.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("err:>", err)
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("Raw Response Body:\n%v\n", string(body))
}
