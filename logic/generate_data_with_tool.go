package logic

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const MaxTurns = 100

func GenerateDataWithTool[out any](ctx context.Context, g *genkit.Genkit, tools ai.CommonGenOption, messages []*ai.Message, opts ...ai.GenerateOption) (*out, error) {
	toolOpts := append([]ai.GenerateOption{tools, ai.WithMessages(messages...), ai.WithMaxTurns(MaxTurns)}, opts...)
	resp, err := genkit.Generate(ctx, g, toolOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to generate: %w", err)
	}

	messages = resp.History()
	messages = append(messages, &ai.Message{
		Role: ai.RoleUser,
		Content: []*ai.Part{
			{
				Text: "Above is the preprocessed conversation history. Please provide the final answer based on the conversation.",
			},
		},
	})

	result, _, err := genkit.GenerateData[out](ctx, g, append(opts, ai.WithMessages(messages...))...)
	if err != nil {
		return nil, fmt.Errorf("failed to generate data: %w", err)
	}

	return result, nil
}
