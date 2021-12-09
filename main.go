package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	// ApiKey := os.Getenv("API_KEY")
	// ApiSecret := os.Getenv("API_KEY_SECRET")
	BearerToken := "Bearer" + " " + os.Getenv("BEARER_TOKEN")

	req := getRetweeters("https://api.twitter.com/2/tweets/1341196353582419969/retweeted_by", BearerToken)
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(BearerToken)
	fmt.Println(string(body))

}

func getRetweeters(endpoint string, token string) *http.Response {
	req, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	req.Header.Add("authorization", token)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req
}
