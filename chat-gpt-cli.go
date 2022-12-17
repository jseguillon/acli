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
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var apiKey string
	var text string
	var maxTokens int
	var temperature float32
	var frequencyPenalty float32
	var presencePenalty float32
	var n int

	var rootCmd = &cobra.Command{
		Use:   "gpt-chat",
		Short: "Send a message to GPT chat and get a response",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the string argument to send to GPT chat
			if len(args) < 1 {
				log.Fatal("Please provide a string argument to send to GPT chat")
			}
			text = args[0]

			if temperature < 0 || temperature > 1 {
				log.Fatal("Please provide a temperature between 0 and 1")
			}
			if frequencyPenalty < -2.0 || frequencyPenalty > 2.0 {
				log.Fatal("Please provide a frequency penalty between -2.0 and 2.0")
			}
			if presencePenalty < -2.0 || presencePenalty > 2.0 {
				log.Fatal("Please provide a presence penalty between -2.0 and 2.0")
			}
			// Ensure prompt max_tokens is not more than 4096
			maxTokensWithPrompt := 4096 - len(text)
			if maxTokensWithPrompt < 4096 {
				maxTokens = maxTokensWithPrompt
			}

			// Create a new HTTP client
			client := &http.Client{}

			// Create JSON data to send in the request body
			var jsonData = []byte(`
			{
				"model": "text-davinci-003",
				"prompt": "` + text + `",
				"max_tokens": ` + strconv.Itoa(maxTokens) + `,
				"temperature": ` + fmt.Sprintf("%f", temperature) + `,
				"frequency_penalty": ` + fmt.Sprintf("%f", frequencyPenalty) + `,
				"presence_penalty": ` + fmt.Sprintf("%f", presencePenalty) + `,
				"n": ` + strconv.Itoa(n) + `,
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
				fmt.Print(strings.TrimPrefix(c.Text, "\n"))
				if n > 1 {
					fmt.Println("")
					fmt.Println("\n---------------------")
				} else {
					fmt.Println("")
				}
			}
		},
	}

	rootCmd.Flags().IntVarP(&maxTokens, "max-tokens", "m", 2048, `The maximum number of tokens to generate in the completion. The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length of 2048 tokens (except for the newest models, which support 4096).`)
	rootCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.1, `What sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.`)
	rootCmd.Flags().Float32VarP(&frequencyPenalty, "frequency-penalty", "f", 0, `Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().Float32VarP(&presencePenalty, "presence-penalty", "p", 0, `Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().IntVarP(&n, "", "n", 1, `How many completions to generate for each prompt. Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop.`)
	// rootCmd.MarkFlagRequired("api-key")
	// Get the API key from the CHAT_GPT_API_KEY environment variable
	apiKey = os.Getenv("CHAT_GPT_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the CHAT_GPT_API_KEY environment variable")
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
