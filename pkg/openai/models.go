package openai

func GetModelsDefaultToken(model string, prompt string) int {
	maxTokens := 512

	switch model {
	case "text-davinci-003":
		maxTokens = 4000
	case "text-curie-001":
		maxTokens = 2048
	case "code-davinci-002":
		maxTokens = 4000
	case "text-babbage-001":
		maxTokens = 2048
	case "code-cushman-001":
		maxTokens = 2048
	}

	maxTokens = maxTokens - len(prompt)

	return maxTokens
}
