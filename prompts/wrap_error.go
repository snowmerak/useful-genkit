package prompts

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const WrapErrorPromptName = "WrapErrorPrompt"

type WrapErrorInput struct {
	Code     string `json:"code"`
	BasePath string `json:"base_path"`
	FilePath string `json:"file_path"`
}

type WrapErrorOutput struct {
	Code string `json:"code"`
}

func WrapErrorPrompt(g *genkit.Genkit) ai.Prompt {
	return genkit.DefinePrompt(g, WrapErrorPromptName, ai.WithPrompt(`You are a Go expert. Your task is to refactor the given Go code.
Find all occurrences where an error is returned directly (e.g., "return err", "return nil, err").
Replace them with "fmt.Errorf" to wrap the error with meaningful context based on the function name and operation being performed.
Use the "return fmt.Errorf(\"context message: %%w\", err)" pattern.
Do NOT change any other logic.
Do NOT wrap errors that are already wrapped or created with fmt.Errorf or errors.New.
Only wrap raw "err" variables being returned.

Base Path: {{base_path}}
File Path: {{file_path}}

Here is the code:

{{code}}

Return the FULL source code with the modifications applied. Do not omit any parts of the code.`), ai.WithInputType(WrapErrorInput{}), ai.WithConfig(&ai.GenerationCommonConfig{
		Temperature: 0.1,
	}))
}
