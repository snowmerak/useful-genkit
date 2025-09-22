package tools

import (
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

type CurrentTimeInput struct {
}

type CurrentTimeOutput struct {
	Time time.Time `json:"time"`
}

func CurrentTime(g *genkit.Genkit) ai.Tool {
	t := genkit.DefineTool(g, "GetCurrentTime", "A tool to get the current time.", func(ctx *ai.ToolContext, _ CurrentTimeInput) (CurrentTimeOutput, error) {
		return CurrentTimeOutput{
			Time: time.Now(),
		}, nil
	})

	return t
}
