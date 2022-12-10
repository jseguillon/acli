# chat-gpt-cli

A command line interface for chatting with the OpenAI GPT-3 language model.

## Requirements

- An OpenAI API key with access to the GPT-3 model, set as the `CHAT_GPT_API_KEY` environment variable

## Install

Go to [releases page](releases), find appropriate binary for your system. Download, install where you want and `chmod +x` it. Example: 

```
curl -SL -o [release_url] /usr/local/bin/chat-gpt-cli
chmod +x /usr/local/bin/chat-gpt-cli
```


## Usage

### Get open API key

Sign up for an open API account on their website: https://beta.openai.com/docs/getting-started/application-setup. After creating an account and signing in, you can create an API key by clicking on your user name in the top right corner, then selecting "API Keys" from the dropdown menu. Click on the "New API Key" button to create a new API key and copy it to your clipboard. 


### Run

Set your API key via the `CHAT_GPT_API_KEY` environnement variable, then run `chat-gpt-cli`:

```
chat-gpt-cli "enter here your question"
```

The program will send the provided string argument to GPT-3 and print the response to the command line.

## TODOs

- allow to continue chat via a `-c` flag
- add tests 

## License

This program is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
