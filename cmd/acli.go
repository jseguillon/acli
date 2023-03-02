package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	o "pkg/openai"
	s "pkg/scripts"
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
	var model string
	var script string
	var jsonData []byte

	// Define the root command
	var rootCmd = &cobra.Command{
		Use: "acli",
		Short: `Send a message to Open AI's model and get a response. Needs ACLI_OPENAI_KEY env var to be defined. 

# Sample usage:

* use 'acli' for discussions or complex task solving: 
	> acli "can GPT help me for daily command line tasks ?"
	> acli "[complex description of feature request for bash/javascript/python/etc...]"
* use 'howto' function for quick one liner answers and interactive mode: 
	> howto openssl test SSL expiracy of github.com
	> howto "find all files more than 30Mb "
* use 'fix' for quick fixing typos: 
	[run typo command like 'rrm', 'lls', 'cd..', etc..]
	then type 'fix' and get fixed command ready to run`,

		Run: func(cmd *cobra.Command, args []string) {
			// Get the API key from the ACLI_OPENAI_KEY environment variable
			apiKey = os.Getenv("ACLI_OPENAI_KEY")
			if apiKey == "" {
				log.Fatal("Please set the ACLI_OPENAI_KEY environment variable")
			}

			// Get the string argument to send to GPT chat
			if len(args) < 1 {
				log.Fatal("Please provide a string argument to send")
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

			if script != "" {
				if script == "fixCmd" {
					text = s.GetScriptFixCmdQueryPrompt(text, args[0], args[1])
				} else if script == "howCmd" {
					text = s.GetScriptHowCmdQueryPrompt(text, args[0])
				}
				model = "gpt-3.5-turbo"
				jsonData = o.OpenAIQuery(text, maxTokens, temperature, frequencyPenalty, presencePenalty, n, model)
			} else {
				if !cmd.Flags().Changed("max-tokens") {
					maxTokens = o.GetModelsDefaultToken(model, text)
				}
				if maxTokens < 0 {
					log.Fatal("Please give a shorted promt or override token estimation using '-m' flag. model max - estimated < 0")
				}
				jsonData = o.OpenAIQuery(text, maxTokens, temperature, frequencyPenalty, presencePenalty, n, model)
			}

			// Create a new HTTP client
			client := &http.Client{}

			endpoint := "https://api.openai.com/v1/chat/completions"
			if model != "gpt-3.5-turbo" {
				endpoint = "https://api.openai.com/v1/chat/completions"
			}
			// Build the request
			req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
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

			// Define a struct for non gpt3.5 response
			type jsonObject struct {
				Choices []struct {
					Text string `json:"text"`
				} `json:"choices"`
			}

			//gpt3.5 message
			type ChatGPTMessage struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			}

			//gpt3.5 choice
			type ChatGPTChoice struct {
				Index        int            `json:"index"`
				FinishReason string         `json:"finish_reason"`
				Message      ChatGPTMessage `json:"message"`
			}

			//gpt3.5 response
			type ChatGPTMessageResponse struct {
				ID      string          `json:"id"`
				Object  string          `json:"object"`
				Created int64           `json:"created"`
				Model   string          `json:"model"`
				Choices []ChatGPTChoice `json:"choices"`
			}

			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode != 200 {
				log.Fatal("Error ", resp.StatusCode, string(body[:]))
			}

			var respObj interface{}
			if model == "gpt-3.5-turbo" {
				respObj = &ChatGPTMessageResponse{}
			} else {
				respObj = &jsonObject{}
			}
			err = json.Unmarshal(body, &respObj)

			if err != nil {
				log.Fatal("Error unmarshalling ", err)
			}

			//Case of gpt-turbo aka ChatGPT
			if turrboAnswer, ok := respObj.(*ChatGPTMessageResponse); ok {
				if !ok {
					log.Fatal("Error ", string(body[:]))
				}

				if len(turrboAnswer.Choices) == 0 {
					log.Fatal("Error ", "empty answer")
				}

				if script != "" {
					s.RunScript(turrboAnswer.Choices[0].Message.Content)
				} else {
					fmt.Print(strings.TrimPrefix(turrboAnswer.Choices[0].Message.Content, "\n"))
				}
			} else {
				answer, ok := respObj.(*jsonObject)
				if !ok {
					log.Fatal("Error ", string(body[:]))
				}

				if len(answer.Choices) == 0 {
					log.Fatal("Error ", "empty answer")
				}

				// Print the response
				for _, c := range answer.Choices {
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
	rootCmd.Flags().IntVarP(&maxTokens, "max-tokens", "m", -1, `The maximum number of tokens to generate in the completion. 
Defaults to model max-tokens minus prompt lenght.
Models max: text-davinci-003=4000, text-curie-001=2048, code-davinci-002=8000, text-babbage-001=2048, code-cushman-001=2048"`)
	rootCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.1, `What sampling temperature to use. 
Higher values means the model will take more risks. 
Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer.`)
	rootCmd.Flags().Float32VarP(&frequencyPenalty, "frequency-penalty", "f", 0, `Number between -2.0 and 2.0. 
Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().Float32VarP(&presencePenalty, "presence-penalty", "p", 0, `Number between -2.0 and 2.0. 
Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.`)
	rootCmd.Flags().IntVarP(&n, "choices", "n", 1, `How many completions to generate for each prompt. 
Note: Because this parameter generates many completions, it can quickly consume your token quota. 
Use carefully and ensure that you have reasonable settings for max_tokens and stop.`)
	rootCmd.Flags().StringVarP(&model, "model", "", "gpt-3.5-turbo", `Open AI model to use. Some examples:
- gpt-3.5-turbo: Chat GPT model, (no other flags apply),
- text-davinci-003: previously most capable GPT-3 model,
- code-davinci-002: most capable Codex model. Particularly good at translating natural language to code,
- text-curie-001: very capable, but faster and lower cost than Davinci. 
(See https://beta.openai.com/docs/models/ for more)`)
	rootCmd.Flags().StringVarP(&script, "script", "", "", `Run embedded script.`)

	// Parse the command-line arguments
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
