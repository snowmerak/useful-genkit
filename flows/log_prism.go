package flows

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/snowmerak/useful-genkit/logic"
	"github.com/snowmerak/useful-genkit/models"
	"github.com/snowmerak/useful-genkit/prompts"
	"github.com/snowmerak/useful-genkit/tools"
)

type LogPrismFlowInput struct {
	Path string `json:"path"`
}

type LogPrismFlowOutput struct {
	ProcessedFiles []string `json:"processed_files"`
}

func LogPrismFlow(g *genkit.Genkit) {
	genkit.DefineFlow(g, "LogPrismFlow", func(ctx context.Context, input LogPrismFlowInput) (LogPrismFlowOutput, error) {
		var processedFiles []string

		// 1. List all files in the directory recursively
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
			return LogPrismFlowOutput{}, fmt.Errorf("failed to walk directory: %w", err)
		}

		// 2. Process each Go file
		for _, file := range files {
			// Read file content
			contentBytes, err := os.ReadFile(file)
			if err != nil {
				return LogPrismFlowOutput{}, fmt.Errorf("failed to read file %s: %w", file, err)
			}
			content := string(contentBytes)

			// Generate modified code
			prompt := genkit.LookupPrompt(g, prompts.LogPrismPromptName)
			if prompt == nil {
				return LogPrismFlowOutput{}, fmt.Errorf("prompt %s not found", prompts.LogPrismPromptName)
			}

			// Use a model to generate the response
			model, err := models.GetOpenRouterQwen3Coder(g)
			if err != nil {
				return LogPrismFlowOutput{}, fmt.Errorf("failed to get model: %w", err)
			}

			req, err := prompt.Render(ctx, prompts.LogPrismInput{
				Code:     content,
				BasePath: input.Path,
				FilePath: file,
			})
			if err != nil {
				return LogPrismFlowOutput{}, fmt.Errorf("failed to render prompt: %w", err)
			}

			// Add tools to the request (optional, but good for context if needed)
			findDefTool := genkit.LookupTool(g, tools.FindDefinitionTool)
			findUsageTool := genkit.LookupTool(g, tools.FindUsageTool)
			findStructsTool := genkit.LookupTool(g, tools.FindStructsTool)
			readFileTool := genkit.LookupTool(g, tools.ReadFileTool)
			writeFileTool := genkit.LookupTool(g, tools.WriteFileTool)
			listFilesTool := genkit.LookupTool(g, tools.ListFilesTool)
			createDirectoryTool := genkit.LookupTool(g, tools.CreateDirectoryTool)
			deleteDirectoryTool := genkit.LookupTool(g, tools.DeleteDirectoryTool)
			walkDirectoryTool := genkit.LookupTool(g, tools.WalkDirectoryTool)

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
			if readFileTool != nil {
				toolRefs = append(toolRefs, readFileTool)
			}
			if writeFileTool != nil {
				toolRefs = append(toolRefs, writeFileTool)
			}
			if listFilesTool != nil {
				toolRefs = append(toolRefs, listFilesTool)
			}
			if createDirectoryTool != nil {
				toolRefs = append(toolRefs, createDirectoryTool)
			}
			if deleteDirectoryTool != nil {
				toolRefs = append(toolRefs, deleteDirectoryTool)
			}
			if walkDirectoryTool != nil {
				toolRefs = append(toolRefs, walkDirectoryTool)
			}

			result, err := logic.GenerateDataWithTool[prompts.LogPrismOutput](
				ctx,
				g,
				ai.WithTools(toolRefs...),
				req.Messages,
				ai.WithModel(model),
			)
			if err != nil {
				// Log error but continue processing other files? Or fail?
				// For now, let's log and continue or return error.
				// wrap_go_error.go seems to return error inside the loop in my reading, but let's check logic.GenerateDataWithTool usage.
				// It returns result and error.
				return LogPrismFlowOutput{}, fmt.Errorf("failed to generate code for %s: %w", file, err)
			}

			// Write back to file
			if result.Code != "" && result.Code != content {
				err = os.WriteFile(file, []byte(result.Code), 0644)
				if err != nil {
					return LogPrismFlowOutput{}, fmt.Errorf("failed to write file %s: %w", file, err)
				}
				processedFiles = append(processedFiles, file)
			}
		}

		return LogPrismFlowOutput{
			ProcessedFiles: processedFiles,
		}, nil
	})
}
