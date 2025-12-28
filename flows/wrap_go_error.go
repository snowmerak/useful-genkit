package flows

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/snowmerak/useful-genkit/logic"
	"github.com/snowmerak/useful-genkit/models"
	"github.com/snowmerak/useful-genkit/prompts"
	"github.com/snowmerak/useful-genkit/tools"
)

type WrapGoErrorInput struct {
	Path string `json:"path"`
}

type WrapGoErrorOutput struct {
	ProcessedFiles []string `json:"processed_files"`
}

func WrapGoErrorFlow(g *genkit.Genkit) {
	genkit.DefineFlow(g, "WrapGoErrorFlow", func(ctx context.Context, input WrapGoErrorInput) (WrapGoErrorOutput, error) {
		var processedFiles []string

		// 1. List all files in the directory recursively (Inline implementation)
		var files []string
		err := filepath.Walk(input.Path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return WrapGoErrorOutput{}, fmt.Errorf("failed to walk directory: %w", err)
		}

		// 2. Process each Go file
		for _, file := range files {
			if !strings.HasSuffix(file, ".go") {
				continue
			}

			// Read file content (Inline implementation)
			contentBytes, err := os.ReadFile(file)
			if err != nil {
				return WrapGoErrorOutput{}, fmt.Errorf("failed to read file %s: %w", file, err)
			}
			content := string(contentBytes)

			// Check if it needs wrapping (simple heuristic to save tokens)
			// if !strings.Contains(content, "return err") && !strings.Contains(content, "return nil, err") {
			// 	continue
			// }

			// Generate modified code
			prompt := genkit.LookupPrompt(g, prompts.WrapErrorPromptName)
			if prompt == nil {
				return WrapGoErrorOutput{}, fmt.Errorf("prompt %s not found", prompts.WrapErrorPromptName)
			}

			// Use a model to generate the response
			model, err := models.GetGoogleAI(g, models.GoogleAIGemini2o5FlashLite)
			if err != nil {
				return WrapGoErrorOutput{}, fmt.Errorf("failed to get model: %w", err)
			}

			req, err := prompt.Render(ctx, prompts.WrapErrorInput{
				Code:     content,
				BasePath: input.Path,
				FilePath: file,
			})
			if err != nil {
				return WrapGoErrorOutput{}, fmt.Errorf("failed to render prompt: %w", err)
			}

			// Add tools to the request
			findDefTool := genkit.LookupTool(g, tools.FindDefinitionTool)
			findUsageTool := genkit.LookupTool(g, tools.FindUsageTool)
			findStructsTool := genkit.LookupTool(g, tools.FindStructsTool)

			var toolRefs []ai.ToolRef
			if findDefTool != nil {
				toolRefs = append(toolRefs, findDefTool)
			}
			if findUsageTool != nil {
				toolRefs = append(toolRefs, findUsageTool)
			}
			if findStructsTool != nil {
				toolRefs = append(toolRefs, findStructsTool)
			}

			result, err := logic.GenerateDataWithTool[prompts.WrapErrorOutput](
				ctx,
				g,
				ai.WithTools(toolRefs...),
				req.Messages,
				ai.WithModel(model),
				// ai.WithConfig(req.Config),
			)
			if err != nil {
				return WrapGoErrorOutput{}, fmt.Errorf("failed to generate code for %s: %w", file, err)
			}

			newCode := result.Code
			// Clean up markdown code blocks if present
			newCode = strings.TrimPrefix(newCode, "```go")
			newCode = strings.TrimPrefix(newCode, "```")
			newCode = strings.TrimSuffix(newCode, "```")
			newCode = strings.TrimSpace(newCode)

			// Write back to file (Inline implementation)
			if err := os.WriteFile(file, []byte(newCode), 0644); err != nil {
				return WrapGoErrorOutput{}, fmt.Errorf("failed to write file %s: %w", file, err)
			}

			processedFiles = append(processedFiles, file)
		}

		return WrapGoErrorOutput{ProcessedFiles: processedFiles}, nil
	})
}
