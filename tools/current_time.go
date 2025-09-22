package tools

import (
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

type GetCurrentTimeInput struct {
}

type GetCurrentTimeOutput struct {
	Time time.Time `json:"time"`
}

func GetCurrentTime(g *genkit.Genkit) ai.Tool {
	t := genkit.DefineTool(g, "GetCurrentTime", "A tool to get the current time.", func(ctx *ai.ToolContext, _ GetCurrentTimeInput) (GetCurrentTimeOutput, error) {
		return GetCurrentTimeOutput{
			Time: time.Now(),
		}, nil
	})

	return t
}
