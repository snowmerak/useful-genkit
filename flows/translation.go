package flows

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/snowmerak/useful-genkit/models"
	"github.com/snowmerak/useful-genkit/prompts"
)

type TranslationInput struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
	Domain string `json:"domain"`
}

type TranslationOutput struct {
	Translated string `json:"translated"`
}

const TranslationFlowName = "TranslationFlow"

func TranslationFlow(g *genkit.Genkit) {
	genkit.DefineFlow(g, TranslationFlowName, func(ctx context.Context, input *TranslationInput) (*TranslationOutput, error) {
		m, err := models.GetOllamaGptOss20b(g)
		if err != nil {
			return nil, fmt.Errorf("failed to get model: %w", err)
		}

		messages, err := prompts.RenderTranslationPrompt(ctx, g, input.Text, input.Source, input.Target, input.Domain)
		if err != nil {
			return nil, fmt.Errorf("failed to render translation prompt: %w", err)
		}

		resp, _, err := genkit.GenerateData[TranslationOutput](ctx, g, ai.WithMessages(messages...), ai.WithModel(m))
		if err != nil {
			return nil, fmt.Errorf("failed to generate translation: %w", err)
		}

		return resp, nil
	})
}
