package openai

import (
	"log"
)

func GetModelsDefaultToken(model string, prompt string) int {
	maxTokens := 512
	encoder, err := NewEncoder()
	if err != nil {
		log.Fatal(err)
	}

	encoded, err := encoder.Encode(prompt)
	if err != nil {
		log.Fatal(err)
	}

	ratio := 1.15
	switch model {
	case "gpt-3.5-turbo":
		maxTokens = 4000
	case "text-davinci-003":
		maxTokens = 4000 - int((ratio * float64(len(encoded))))
	case "text-curie-001":
		maxTokens = 2048 - len(prompt)/2
	case "code-davinci-002":
		maxTokens = 4000 - int((ratio * float64(len(encoded))))
	case "text-babbage-001":
		maxTokens = 2048 - len(prompt)/2
	case "code-cushman-001":
		maxTokens = 2048 - len(prompt)/2
	}

	return maxTokens
}
