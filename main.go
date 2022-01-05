package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	//loads env file that has key and secret
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	ApiKey := os.Getenv("API_KEY")
	ApiSecret := os.Getenv("API_KEY_SECRET")

	//function that makes an auth request to retieve vearer token
	tokenReq := getToken("https://api.twitter.com/oauth2/token?grant_type=client_credentials", ApiKey, ApiSecret)

	var client http.Client
	tokenRes, err := client.Do(tokenReq)
	if err != nil {
		log.Fatal(err)
	}
	defer tokenRes.Body.Close()

	var token oauth2.Token
	dec := json.NewDecoder(tokenRes.Body)
	err = dec.Decode(&token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conf := &oauth2.Config{}

	twitterClient := conf.Client(ctx, &token)
	twitterRes, err := twitterClient.Get("https://api.twitter.com/2/tweets/1341196353582419969/retweeted_by")
	if err != nil {
		log.Fatal(err)
	}
	defer twitterRes.Body.Close()
	//outputs response body within the terminal of all users who retweeted the tweet associated with the tweet_id
	io.Copy(os.Stdout, twitterRes.Body)

}

func getToken(endpoint string, key string, secret string) *http.Request {
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req
}
