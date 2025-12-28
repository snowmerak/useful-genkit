package models

import (
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	oai "github.com/firebase/genkit/go/plugins/compat_oai"
)

const OpenrouterProvider = "openai"

func OpenRouterDevstral2512Free(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	o, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := o.DefineModel(OpenrouterProvider, "mistralai/devstral-2512:free", ai.ModelOptions{
		Supports: &ai.ModelSupports{
			ToolChoice: true,
			Tools:      true,
		},
	})
	if m == nil {
		return nil, fmt.Errorf("model mistralai/devstral-2512:free not found in openrouter plugin")
	}

	return m, nil
}

func GetOpenRouterDevstral2512Free(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	_, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/mistralai/devstral-2512:free", OpenrouterProvider))
	if m == nil {
		return nil, fmt.Errorf("model mistralai/devstral-2512:free not found, make sure to define it first")
	}
	return m, nil
}

func OpenRouterDevstral2512(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	o, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := o.DefineModel(OpenrouterProvider, "mistralai/devstral-2512", ai.ModelOptions{
		Supports: &ai.ModelSupports{
			ToolChoice: true,
			Tools:      true,
		},
	})
	if m == nil {
		return nil, fmt.Errorf("model mistralai/devstral-2512 not found in openrouter plugin")
	}

	return m, nil
}

func GetOpenRouterDevstral2512(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	_, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/mistralai/devstral-2512", OpenrouterProvider))
	if m == nil {
		return nil, fmt.Errorf("model mistralai/devstral-2512 not found, make sure to define it first")
	}
	return m, nil
}

func OpenRouterQwen3CoderFree(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	o, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := o.DefineModel(OpenrouterProvider, "qwen/qwen3-coder:free", ai.ModelOptions{
		Supports: &ai.ModelSupports{
			ToolChoice: true,
			Tools:      true,
		},
	})
	if m == nil {
		return nil, fmt.Errorf("model qwen/qwen3-coder:free not found in openrouter plugin")
	}

	return m, nil
}

func GetOpenRouterQwen3CoderFree(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	_, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/qwen/qwen3-coder:free", OpenrouterProvider))
	if m == nil {
		return nil, fmt.Errorf("model qwen/qwen3-coder:free not found, make sure to define it first")
	}
	return m, nil
}

func OpenRouterQwen3Coder(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	o, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := o.DefineModel(OpenrouterProvider, "qwen/qwen3-coder", ai.ModelOptions{
		Supports: &ai.ModelSupports{
			ToolChoice: true,
			Tools:      true,
		},
	})
	if m == nil {
		return nil, fmt.Errorf("model qwen/qwen3-coder not found in openrouter plugin")
	}

	return m, nil
}

func GetOpenRouterQwen3Coder(g *genkit.Genkit) (ai.Model, error) {
	p := genkit.LookupPlugin(g, OpenrouterProvider)
	if p == nil {
		return nil, fmt.Errorf("openrouter plugin not found, make sure to initialize genkit with openrouter plugin")
	}

	_, ok := p.(*oai.OpenAICompatible)
	if !ok {
		return nil, fmt.Errorf("plugin is not of type compat_oai.CompatOAI")
	}

	m := genkit.LookupModel(g, fmt.Sprintf("%s/qwen/qwen3-coder", OpenrouterProvider))
	if m == nil {
		return nil, fmt.Errorf("model qwen/qwen3-coder not found, make sure to define it first")
	}
	return m, nil
}
