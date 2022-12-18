package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Get the API key from the CHAT_GPT_API_KEY environment variable
	apiKey := os.Getenv("CHAT_GPT_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the CHAT_GPT_API_KEY environment variable")
	}

	// Get the string argument to send to GPT chat
	if len(os.Args) < 2 {
		log.Fatal("Please provide a string argument to send to GPT chat")
	}
	text := os.Args[1]

	// Create a new HTTP client
	client := &http.Client{}

	// Ensure prompt max_tokens is not more than 4096
	max_tokens := 2048
	max_tokens_with_prompt := 4096 - len(text)
	if max_tokens_with_prompt < 2048 {
		max_tokens = max_tokens_with_prompt
	}

	// Create JSON data to send in the request body
	var jsonData = []byte(`
	{
		"model": "text-davinci-003",
		"prompt": "` + text + `",
		"max_tokens": ` + strconv.Itoa(max_tokens) + `,
		"temperature": 0.1,
		"frequency_penalty": 0,
		"presence_penalty": 0,
		"stream": false
	}
	`)

	// Build the request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	// Set the API key in the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Define a struct for the JSON response
	type jsonObject struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON object into a struct
	var obj jsonObject
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response
	for _, c := range obj.Choices {
		fmt.Print(c.Text)
	}
	fmt.Println("")
}
