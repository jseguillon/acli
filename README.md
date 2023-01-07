# chat-gpt-cli

A command line interface for interacting with openAI's AI models.

## Requirements

### Get open API key

Sign up for an open API account on their website: https://openai.com/api/. After signing in, create an API key at this URL: https://beta.openai.com/account/api-keys. 

## Usage

Use `chat-gpt-cli` for discussions or complex task solving. Examples: 
* `   chat-gpt-cli "can GPT help me for daily command line tasks ?" `
* `   chat-gpt-cli "[complex description of feature request for bash/javascript/python/etc...]" `

Use `howto` function for quick one liner answers and interactive mode. Examples:
* `   howto openssl test SSL expiracy of github.com`
* `   howto "find all files more than 30Mb "`

Use `fix` for quick fixing typos. Examples:
* [run typo command like 'rrm', 'lls', 'cd..', etc..]
* then type `fix` and get fixed command ready to run

## Install

### Script install

Run:
```
curl -sSLO https://raw.githubusercontent.com/jseguillon/chat-gpt-cli/main/get.sh && \
bash get.sh
```

### Or manual install

Go to [releases page](https://github.com/jseguillon/chat-gpt-cli/releases), find appropriate binary for your system. Download, install where you want and `chmod +x` it. Example: 

```
sudo curl -SL [release_url] -o /usr/local/bin/chat-gpt-cli
sudo chmod +x /usr/local/bin/chat-gpt-cli
```

Add configuration in any `.rc` file you want:

```
CHAT_GPT_API_KEY="XXXXX"

alias fix='eval $(chat-gpt-cli --script fixCmd "$(fc -nl -1)" $?)'
howto() { h="$@"; eval $(chat-gpt-cli --script howCmd "$h") ; }
```

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

- allow to continue chat via a `-c` flag when openAI offers this feature on its API
- add tests 

## License

This program is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
