package services

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/kyokomi/emoji/v2"
	"go-chatgpt-cli/logger"
)

var client gpt3.Client

func InitClient(key string) {
	client = gpt3.NewClient(key)
}

func GetClient() gpt3.Client {
	if client == nil {
		logger.LanLogger.Panic("client is nil")
	}
	return client
}

const (
	maxTokens   = 100
	temperature = 0.8
	engine      = gpt3.TextDavinci003Engine
)
const DefaultBotName = ":unicorn_face: GPT: "

func GetResponse(request string) (response string, err error) {
	fmt.Println()
	fmt.Print(emoji.Sprint(DefaultBotName))
	i := 0
	ctx := context.Background()
	err = client.CompletionStreamWithEngine(
		ctx,
		engine,
		gpt3.CompletionRequest{
			Prompt:      []string{request},
			MaxTokens:   gpt3.IntPtr(maxTokens),
			Temperature: gpt3.Float32Ptr(temperature),
		}, func(resp *gpt3.CompletionResponse) {
			fmt.Print(resp.Choices[0].Text)
			i++
		},
	)
	fmt.Println("\n")
	if err != nil {

		logger.LanLogger.Errorf("error: %v", err)
		return "", err
	}

	return "", nil
}
