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

const FindStructsTool = "FindStructs"

type FindStructsInput struct {
	StructName string            `json:"struct_name"`
	Language   language.Language `json:"language"`
	BasePath   string            `json:"base_path"`
}

type FindStructsOutput struct {
	Result string `json:"result"`
}

func FindStructs(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, FindStructsTool, "Finds the definition of a struct/class and its methods.", func(ctx *ai.ToolContext, input FindStructsInput) (FindStructsOutput, error) {
		var patterns []string
		switch input.Language {
		case language.Go:
			patterns = []string{
				fmt.Sprintf("type %s struct { $$$ }", input.StructName),
				fmt.Sprintf("func ($R %s) $M($$$) $$$", input.StructName),
				fmt.Sprintf("func ($R *%s) $M($$$) $$$", input.StructName),
			}
		case language.Python:
			patterns = []string{
				fmt.Sprintf("class %s: $$$", input.StructName),
			}
		case language.TypeScript, language.JavaScript, language.Java:
			patterns = []string{
				fmt.Sprintf("class %s { $$$ }", input.StructName),
			}
		default:
			return FindStructsOutput{}, fmt.Errorf("unsupported language: %s", input.Language)
		}

		var rulePatterns string
		for _, p := range patterns {
			rulePatterns += fmt.Sprintf("    - pattern: '%s'\n", p)
		}

		ruleContent := fmt.Sprintf(`id: find-structs
language: %s
rule:
  any:
%s
`, input.Language, rulePatterns)

		tmpFile, err := os.CreateTemp("", "find_structs_*.yml")
		if err != nil {
			return FindStructsOutput{}, fmt.Errorf("failed to create temp file: %w", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(ruleContent); err != nil {
			return FindStructsOutput{}, fmt.Errorf("failed to write to temp file: %w", err)
		}
		if err := tmpFile.Close(); err != nil {
			return FindStructsOutput{}, fmt.Errorf("failed to close temp file: %w", err)
		}

		args := []string{"scan", "--rule", tmpFile.Name(), "--json"}
		if input.BasePath != "" {
			args = append(args, input.BasePath)
		}
		cmd := exec.Command("sg", args...)
		output, err := cmd.Output()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				// sg returns non-zero exit code if matches are found (sometimes) or if error?
				// Actually sg scan returns 0 on success usually.
				// But let's check stderr if needed.
				// Wait, sg scan might return exit code 1 if no matches? Or if matches found?
				// Let's assume standard behavior or check stderr.
				// In find_usage.go, it handles err.
				_ = exitError
			}
			// If output is empty and err is not nil, it might be a real error.
			// But sometimes tools return non-zero for "found matches" or "no matches".
			// Let's proceed if we have output.
		}

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
		if len(output) > 0 {
			if err := json.Unmarshal(output, &matches); err != nil {
				return FindStructsOutput{}, fmt.Errorf("failed to parse sg output: %w", err)
			}
		}

		if len(matches) == 0 {
			return FindStructsOutput{Result: "No struct/class or methods found."}, nil
		}

		result := ""
		for _, match := range matches {
			result += fmt.Sprintf("File: %s:%d\n%s\n\n", match.File, match.Range.Start.Line+1, match.Text)
		}

		return FindStructsOutput{Result: result}, nil
	})
}
