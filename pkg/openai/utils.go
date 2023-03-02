package openai

import (
	"encoding/json"
	"log"

	"github.com/samber/lo"
)

func dictTuple(tuples []lo.Tuple2[[]byte, []byte]) map[lo.Tuple2[string, string]]int {
	i := -1
	return lo.SliceToMap(tuples, func(item lo.Tuple2[[]byte, []byte]) (lo.Tuple2[string, string], int) {
		i++
		return lo.T2(string(item.A), string(item.B)), i
	})
}

// makeQuery constructs a JSON object for the POST request to the OpenAI API
func OpenAIQuery(text string, maxTokens int, temperature float32, frequencyPenalty float32, presencePenalty float32, n int, model string) []byte {

	// GPTConfig contains the default settings for the GPT API request.
	type GPTConfig struct {
		Model            string  `json:"model"`
		Prompt           string  `json:"prompt"`
		MaxTokens        int     `json:"max_tokens"`
		Temperature      float32 `json:"temperature"`
		FrequencyPenalty float32 `json:"frequency_penalty"`
		PresencePenalty  float32 `json:"presence_penalty"`
		N                int     `json:"n"`
		Stream           bool    `json:"stream"`
	}

	//ChatGPT specific field
	type Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	//Chat GPT specific query
	type ChatGPTConfig struct {
		Model    string    `json:"model"`
		Messages []Message `json:"messages"`
	}

	//var jsonData := nil
	// Chat GPT as API
	if model == "gpt-3.5-turbo" {
		query := &ChatGPTConfig{
			Model: model,
			Messages: []Message{
				{Role: "user", Content: text},
			},
		}
		jsonData, err := json.Marshal(query)
		if err == nil {
			return jsonData
		} else {
			log.Fatal("Could not build Chat GPT query", err)
		}
	} else { // Marshal the JSON object into a byte array (generic)
		query := &GPTConfig{
			Model:            model,
			Prompt:           text,
			MaxTokens:        maxTokens,
			Temperature:      temperature,
			FrequencyPenalty: frequencyPenalty,
			PresencePenalty:  presencePenalty,
			N:                n,
			Stream:           false,
		}
		jsonData, err := json.Marshal(query)
		if err == nil {
			return jsonData
		} else {
			log.Fatal("Could not build GPT query", err)
		}
	}
	return nil
}
