package prompts

import (
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const TranslationPromptName = "Translation"

type TranslationInput struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
	Domain string `json:"domain"`
}

func TranslationPrompt(g *genkit.Genkit) ai.Prompt {
	return genkit.DefinePrompt(g, TranslationPromptName, ai.WithPrompt(`Act as an expert translator specializing in {{domain}}. Your task is to translate the following text from {{source}} to {{target}}, ensuring the highest level of accuracy and preserving the original nuance.

{{text}}`), ai.WithInputType(TranslationInput{}))
}

func RenderTranslationPrompt(ctx context.Context, g *genkit.Genkit, text, sourceLang, targetLang, domain string) ([]*ai.Message, error) {
	p := genkit.LookupPrompt(g, TranslationPromptName)
	if p == nil {
		return nil, fmt.Errorf("prompt %s not found", TranslationPromptName)
	}
	opt, err := p.Render(ctx, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to render prompt: %w", err)
	}
	return opt.Messages, nil
}
