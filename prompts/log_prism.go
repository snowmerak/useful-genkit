package prompts

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const LogPrismPromptName = "LogPrismPrompt"

type LogPrismInput struct {
	Code string `json:"code"`
}

type LogPrismOutput struct {
	Code string `json:"code"`
}

func LogPrismPrompt(g *genkit.Genkit) ai.Prompt {
	return genkit.DefinePrompt(g, LogPrismPromptName, ai.WithPrompt(`You are a Go expert. Your task is to add "Prism" style logging (Span and State) to the given Go code based on the concept below.

`+logPrismConcept+`

## Instructions

1.  **Analyze the Code**: Understand the flow of the provided Go code.
2.  **Add Span Logging**:
    *   Identify functions representing units of work.
    *   Initialize `+"`Span`"+` at the start.
    *   Handle errors with `+"`span.Fail(err)`"+`.
3.  **Add State Logging**:
    *   Log state changes with `+"`StateLogger.Transition`"+`.
    *   Log data snapshots with `+"`StateLogger.Snapshot`"+`.
4.  **Context**: Ensure `+"`context.Context`"+` is used to propagate RequestID.

## Input Code

{{code}}

Return the FULL source code with logging added.`, ai.WithInputType(LogPrismInput{}), ai.WithConfig(&ai.GenerationCommonConfig{
		Temperature: 0.1,
	})))
}
