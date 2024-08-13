package langchain

import (
	"context"
	"log"

	"github.com/euklid-dev/deep_chained/internal/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	// llm *anthropic.LLM
	llm *openai.LLM
)

func Initialize() {
	var err error

	// llm, err = openai.New(openai.WithToken(config.GlobalAppConfig.OPEN_AI_KEY), openai.WithModel("gpt-4o-2024-08-06"))
	llm, err = openai.New(openai.WithToken(config.GlobalAppConfig.OPEN_AI_KEY), openai.WithModel("gpt-4o-mini-2024-07-18"))

	if err != nil {
		log.Fatal(err)
	}
}

func ConversateWithHistory(ctx context.Context, history []llms.MessageContent) (string, error) {
	res, err := llm.GenerateContent(ctx, history, llms.WithJSONMode(), llms.WithCandidateCount(1), llms.WithMaxTokens(16384))

	if err != nil {
		return "", err
	}

	return res.Choices[0].Content, nil
}

func ConversateWithLM(ctx context.Context, systemPrompt string, userPrompt string) (string, error) {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, systemPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, userPrompt),
	}

	res, err := llm.GenerateContent(ctx, messages)

	if err != nil {
		return "", err
	}

	return res.Choices[0].Content, nil
}
