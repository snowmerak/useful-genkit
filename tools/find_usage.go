package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/snowmerak/useful-genkit/utils/language"
)

const FindUsageTool = "FindUsage"

type FindUsageInput struct {
	Query    string            `json:"query"`
	Language language.Language `json:"language"`
	BasePath string            `json:"base_path"`
}

type FindUsageOutput struct {
	Result string `json:"result"`
}

func FindUsage(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, FindUsageTool, "Finds usages of a symbol (function, method, or type) in the codebase using ast-grep.", func(ctx *ai.ToolContext, input FindUsageInput) (FindUsageOutput, error) {
		// Create a temporary rule file
		ruleContent := fmt.Sprintf(`id: find-usage
language: %s
rule:
  any:
    - kind: function_declaration
    - kind: method_declaration
    - kind: type_declaration
  has:
    stopBy: end
    regex: %s
`, input.Language, input.Query)

		tmpFile, err := os.CreateTemp("", "find_usage_*.yml")
		if err != nil {
			return FindUsageOutput{}, fmt.Errorf("failed to create temp file: %w", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(ruleContent); err != nil {
			return FindUsageOutput{}, fmt.Errorf("failed to write to temp file: %w", err)
		}
		if err := tmpFile.Close(); err != nil {
			return FindUsageOutput{}, fmt.Errorf("failed to close temp file: %w", err)
		}

		// Run ast-grep
		args := []string{"scan", "--json", "-r", tmpFile.Name()}
		if input.BasePath != "" {
			args = append(args, input.BasePath)
		}
		cmd := exec.Command("sg", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			// If output is empty, it's a real error. If not, it might just be exit code 1 for no matches or similar (though sg usually returns 0).
			if len(output) == 0 {
				return FindUsageOutput{Result: fmt.Sprintf("Error running ast-grep: %v", err)}, nil
			}
		}

		// Parse JSON output to make it more readable or just return as string
		type SgMatch struct {
			Text  string `json:"text"`
			File  string `json:"file"`
			Range struct {
				Start struct {
					Line int `json:"line"`
				} `json:"start"`
				End struct {
					Line int `json:"line"`
				} `json:"end"`
			} `json:"range"`
		}
		var matches []SgMatch
		if err := json.Unmarshal(output, &matches); err != nil {
			// If not JSON, return raw output (it might be an error message)
			return FindUsageOutput{Result: string(output)}, nil
		}

		if len(matches) == 0 {
			return FindUsageOutput{Result: "No usages found."}, nil
		}

		// Format the output
		var result string
		for _, match := range matches {
			result += fmt.Sprintf("File: %s (Line %d:%d)\n```%s\n%s\n```\n\n", match.File, match.Range.Start.Line+1, match.Range.End.Line+1, input.Language, match.Text)
		}

		return FindUsageOutput{
			Result: result,
		}, nil
	})
}
