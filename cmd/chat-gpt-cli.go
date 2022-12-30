package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	o "pkg/openai"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	// Declare variables for command-line arguments
	var apiKey string
	var text string
	var maxTokens int
	var temperature float32
	var frequencyPenalty float32
	var presencePenalty float32
	var n int
	var prompt string
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "gpt-chat",
		Short: "Send a message to GPT chat and get a response. Needs CHAT_GPT_API_KEY env var to be defined.",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the API key from the CHAT_GPT_API_KEY environment variable
			apiKey = os.Getenv("CHAT_GPT_API_KEY")
			if apiKey == "" {
				log.Fatal("Please set the CHAT_GPT_API_KEY environment variable")
			}

			// Get the string argument to send to GPT chat
			if len(args) < 1 {
				log.Fatal("Please provide a string argument to send to GPT chat")
			}
			text = args[0]

			// Validate the temperature argument
			if temperature < 0 || temperature > 1 {
				log.Fatal("Please provide a temperature between 0 and 1")
			}
			// Validate the frequency penalty argument
			if frequencyPenalty < -2.0 || frequencyPenalty > 2.0 {
				log.Fatal("Please provide a frequency penalty between -2.0 and 2.0")
			}
			// Validate the presence penalty argument
			if presencePenalty < -2.0 || presencePenalty > 2.0 {
				log.Fatal("Please provide a presence penalty between -2.0 and 2.0")
			}
			// Promt named 'fix' has some specific requirements
			if prompt != "" && prompt == "fix" {
				text = fmt.Sprintf("Fix given command with error. Answer shell command that can be piped. \\n\\nCommand: \"%s\" with error %s \\nFixed command, no leading #: ", args[0], args[1])
				maxTokens = 256
				temperature = 0.1
			} else { // Ensure prompt max_tokens is not more than 4096
				maxTokensWithPrompt := 4096 - len(text)
				if maxTokensWithPrompt < 4096 {
					maxTokens = maxTokensWithPrompt
					if maxTokens < 0 {
						log.Fatal("Error prompt is too long. Model max token is 4096 but you provided a prompt of length: ", len(text))
					}
				}
			}

			// Create a new HTTP client
			client := &http.Client{}
			jsonData := o.OpenAIQuery(text, maxTokens, temperature, frequencyPenalty, presencePenalty, n)

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

			if resp.StatusCode != 200 {
				log.Fatal("Error ", resp.StatusCode, string(body[:]))
			}

			if prompt != "" && prompt == "fix" {
				fmt.Fprint(os.Stderr, strings.TrimLeft(obj.Choices[0].Text, "\n"))
				fmt.Fprintln(os.Stderr, " [^C to escape or Enter to run ]")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				fmt.Println(strings.TrimLeft(obj.Choices[0].Text, "\n"))
			} else {
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
			}
		},
	}

	// Define command-line flags
	rootCmd.Flags().IntVarP(&maxTokens, "max-tokens", "m", 2048, `The maximum number of tokens to generate in the completion. 
The token count of your prompt plus max_tokens cannot exceed the model's context length. Max 4096.`)
	rootCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.1, `What sampling temperature to use. 
Higher values means the model will take more risks. 
Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.`)
	rootCmd.Flags().Float32VarP(&frequencyPenalty, "frequency-penalty", "f", 0, `Number between -2.0 and 2.0. 
Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().Float32VarP(&presencePenalty, "presence-penalty", "p", 0, `Number between -2.0 and 2.0. 
Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().IntVarP(&n, "", "n", 1, `How many completions to generate for each prompt. 
Note: Because this parameter generates many completions, it can quickly consume your token quota. 
Use carefully and ensure that you have reasonable settings for max_tokens and stop.`)
	rootCmd.Flags().StringVarP(&prompt, "prompt", "", "", `Run pre-recorded prompt. Currently, only 'fix' prompt is available`)

	// Parse the command-line arguments
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
