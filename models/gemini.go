package provider

import (
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

const googleAIProvider = "googleai"

const (
	GoogleAIGemini2o5Pro       = "gemini-2.5-pro"
	GoogleAIGemini2o5Flash     = "gemini-2.5-flash"
	GoogleAIGemini2o5FlashLite = "gemini-2.5-flash-lite"
	GoogleAIGemma3o4b          = "gemma-3-4b-it"
	GoogleAIGemma3o12b         = "gemma-3-12b-it"
	GoogleAIGemma3o27b         = "gemma-3-27b-it"
)

func GoogleAI(g *genkit.Genkit, modelName string) (ai.Model, error) {
	p := genkit.LookupPlugin(g, googleAIProvider)
	if p == nil {
		return nil, fmt.Errorf("googleai plugin not found, make sure to initialize genkit with googleai plugin")
	}
	m := googlegenai.GoogleAIModel(g, modelName)
	if m == nil {
		return nil, fmt.Errorf("model %s not found in googleai plugin", modelName)
	}
	return m, nil
}
