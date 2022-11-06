package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ttacon/chalk"
)

const APIBaseURL string = "https://official-joke-api.appspot.com"

var client *http.Client

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetJoke() {
	url := APIBaseURL + "/random_joke"

	var joke Joke

	err := GetJson(url, &joke)

	if err != nil {
		fmt.Printf("error getting BitCoint rate: %s\n", err.Error())
		return
	}

	fmt.Println(chalk.Blue.Color("Time:"), joke.Setup)
	fmt.Println(chalk.Cyan.Color("Time:"), joke.Punchline)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetJoke()
}
