package provider

import (
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

const ollamaProvider = "ollama"

func OllamaGptOss20b(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	return o.DefineModel(g, ollama.ModelDefinition{
		Name: "gpt-oss:20b",
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn:  true,
			ToolChoice: true,
			Tools:      true,
		},
	}), nil
}

func OllamaGemma3(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	modelName := fmt.Sprintf("gemma3:%db", bits)
	return o.DefineModel(g, ollama.ModelDefinition{
		Name: modelName,
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn: true,
		},
	}), nil
}

func OllamaQwen3(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	modelName := fmt.Sprintf("qwen3:%db", bits)
	return o.DefineModel(g, ollama.ModelDefinition{
		Name: modelName,
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn:  true,
			ToolChoice: true,
			Tools:      true,
		},
	}), nil
}
