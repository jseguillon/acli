# chat-gpt-cli

A command line interface for chatting with the OpenAI GPT-3 language model.

## Requirements

- An OpenAI API key with access to the GPT-3 model, set as the `CHAT_GPT_API_KEY` environment variable

## Install

Go to [releases page](https://github.com/jseguillon/chat-gpt-cli/releases), find appropriate binary for your system. Download, install where you want and `chmod +x` it. Example: 

```
sudo curl -SL [release_url] -o /usr/local/bin/chat-gpt-cli
sudo chmod +x /usr/local/bin/chat-gpt-cli
```


## Usage

### Get open API key

Sign up for an open API account on their website: https://openai.com/api/. After creating an account and signing in, you can create an API key by clicking on your user name in the top right corner, then selecting "API Keys" from the dropdown menu. Click on the "New API Key" button to create a new API key and copy it to your clipboard. 


### Run

Set your API key via the `CHAT_GPT_API_KEY` environnement variable, then run `chat-gpt-cli`:

```
CHAT_GPT_API_KEY="XXXXX" chat-gpt-cli "enter here your question"
```

The program will send the provided string argument to GPT-3 and print the response to the command line.

### Flags 

| Short   |     Long      |  Description |
|:----------|:-------------|:------|
|  -n | --choices          |  How many completions to generate for each prompt. Note: Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for max_tokens and stop. (default 1) |
|  -t | --temperature       |  What sampling temperature to use. Higher values means the model will take more risks.  Try 0.9 for more creative applications, and 0 (argmax sampling) for ones with a well-defined answer. (default 0.1) |
|  -f | --frequency-penalty |  Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim. |
|  -m | --max-tokens        |  Defaults to model max-tokens minus prompt lenght. <br/> Models max: text-davinci-003=4000, text-curie-001=2048, code-davinci-002=8000, text-babbage-001=2048, code-cushman-001=2048  |
|  -p | --presence-penalty   |  Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim. |
|     | --model |  Open AI model to use. Some examples:<br/> - text-davinci-003: most capable GPT-3 model, <br/>- code-davinci-002: most capable Codex model. Particularly good at translating natural language to code, <br/>- text-curie-001: very capable, but faster and lower cost than Davinci. <br/> (See https://beta.openai.com/docs/models/ for more) (default "text-davinci-003") | 
## TODOs

- allow to continue chat via a `-c` flag
- add tests 

## License

This program is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
