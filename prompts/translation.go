package prompts

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const Translation = "Translation"

func RegisterTranslationPrompt(g *genkit.Genkit) ai.Prompt {
	return genkit.DefinePrompt(g, Translation, ai.WithPrompt(`Act as an expert translator specializing in {{Domain}}. Your task is to translate the following text from {{Source}} to {{Target}}, ensuring the highest level of accuracy and preserving the original nuance.

{{Text}}`))
}

func ParamsOfTranslationPrompt(text, sourceLang, targetLang, domain string) map[string]any {
	return map[string]any{
		"Text":   text,
		"Source": sourceLang,
		"Target": targetLang,
		"Domain": domain,
	}
}
