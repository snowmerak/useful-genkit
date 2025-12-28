package tools

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const (
	ReadFileTool  = "ReadFile"
	WriteFileTool = "WriteFile"
)

// ReadFileInput defines the input for the ReadFile tool.
type ReadFileInput struct {
	Path string `json:"path"`
}

// ReadFileOutput defines the output for the ReadFile tool.
type ReadFileOutput struct {
	Content string `json:"content"`
}

// ReadFile creates a tool to read the content of a file.
func ReadFile(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, ReadFileTool, "Reads the content of a file at the specified path.", func(ctx *ai.ToolContext, input ReadFileInput) (ReadFileOutput, error) {
		content, err := os.ReadFile(input.Path)
		if err != nil {
			return ReadFileOutput{}, fmt.Errorf("failed to read file: %w", err)
		}
		return ReadFileOutput{Content: string(content)}, nil
	})
}

// WriteFileInput defines the input for the WriteFile tool.
type WriteFileInput struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// WriteFileOutput defines the output for the WriteFile tool.
type WriteFileOutput struct {
	Success bool `json:"success"`
}

// WriteFile creates a tool to write content to a file.
func WriteFile(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, WriteFileTool, "Writes content to a file at the specified path. Overwrites existing content.", func(ctx *ai.ToolContext, input WriteFileInput) (WriteFileOutput, error) {
		if err := os.MkdirAll(filepath.Dir(input.Path), 0755); err != nil {
			return WriteFileOutput{Success: false}, fmt.Errorf("failed to create directories: %w", err)
		}

		if err := os.WriteFile(input.Path, []byte(input.Content), 0644); err != nil {
			return WriteFileOutput{Success: false}, fmt.Errorf("failed to write file: %w", err)
		}
		return WriteFileOutput{Success: true}, nil
	})
}
