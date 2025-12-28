package models

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

func GetOllamaGptOss20b(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	m := genkit.LookupModel(g, fmt.Sprintf("%s/gpt-oss:20b", ollamaProvider))
	if m == nil {
		return nil, fmt.Errorf("model gpt-oss:20b not found, make sure to define it first")
	}
	return m, nil
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

func GetOllamaGemma3(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	modelName := fmt.Sprintf("%s/gemma3:%db", ollamaProvider, bits)
	m := genkit.LookupModel(g, modelName)
	if m == nil {
		return nil, fmt.Errorf("model %s not found, make sure to define it first", modelName)
	}
	return m, nil
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
	modelName := fmt.Sprintf("%s/qwen3:%db", ollamaProvider, bits)
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

func GetOllamaQwen3(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}
	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}
	modelName := fmt.Sprintf("qwen3:%db", bits)
	m := genkit.LookupModel(g, modelName)
	if m == nil {
		return nil, fmt.Errorf("model %s not found, make sure to define it first", modelName)
	}
	return m, nil
}

func OllamaDevstralSmall2(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	return o.DefineModel(g, ollama.ModelDefinition{
		Name: "devstral-small-2:24b",
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn:  true,
			ToolChoice: true,
			Tools:      true,
		},
	}), nil
}

func GetOllamaDevstralSmall2(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/devstral-small-2:24b", ollamaProvider))
	if m == nil {
		return nil, fmt.Errorf("model devstral-small-2:24b not found, make sure to define it first")
	}

	return m, nil
}

func OllamaQwen3Coder(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	modelName := fmt.Sprintf("Nehc/qwen3-coder:%db", bits)
	return o.DefineModel(g, ollama.ModelDefinition{
		Name: modelName,
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn: true,
		},
	}), nil
}

func GetOllamaQwen3Coder(g *genkit.Genkit, bits int) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	modelName := fmt.Sprintf("%s/Nehc/qwen3-coder:%db", ollamaProvider, bits)
	m := genkit.LookupModel(g, modelName)
	if m == nil {
		return nil, fmt.Errorf("model %s not found, make sure to define it first", modelName)
	}

	return m, nil
}

func OllamaMinistral3o14B(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	o, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	return o.DefineModel(g, ollama.ModelDefinition{
		Name: "ministral-3:14b",
		Type: "chat",
	}, &ai.ModelOptions{
		Supports: &ai.ModelSupports{
			Multiturn: true,
		},
	}), nil
}

func GetOllamaMinistral3o14B(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, ollamaProvider)
	if p == nil {
		return nil, fmt.Errorf("ollama plugin not found, make sure to initialize genkit with ollama plugin")
	}

	_, ok := p.(*ollama.Ollama)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type ollama.Ollama")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/ministral-3:14b", ollamaProvider))
	if m == nil {
		return nil, fmt.Errorf("model ministral-3:14b not found, make sure to define it first")
	}

	return m, nil
}
