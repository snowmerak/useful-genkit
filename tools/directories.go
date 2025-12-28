package tools

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const (
	ListFilesTool           = "ListFiles"
	CreateDirectoryTool     = "CreateDirectory"
	DeleteDirectoryTool     = "DeleteDirectory"
	GetCurrentDirectoryTool = "GetCurrentDirectory"
	WalkDirectoryTool       = "WalkDirectory"
)

// ListFilesInput defines the input for the ListFiles tool.
type ListFilesInput struct {
	Path string `json:"path"`
}

// ListFilesOutput defines the output for the ListFiles tool.
type ListFilesOutput struct {
	Files []string `json:"files"`
}

// ListFiles creates a tool to list files in a directory.
func ListFiles(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, ListFilesTool, "Lists files and directories in the specified path.", func(ctx *ai.ToolContext, input ListFilesInput) (ListFilesOutput, error) {
		entries, err := os.ReadDir(input.Path)
		if err != nil {
			return ListFilesOutput{}, fmt.Errorf("failed to read directory: %w", err)
		}

		files := make([]string, 0, len(entries))
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() {
				name += "/"
			}
			files = append(files, name)
		}

		return ListFilesOutput{Files: files}, nil
	})
}

// CreateDirectoryInput defines the input for the CreateDirectory tool.
type CreateDirectoryInput struct {
	Path string `json:"path"`
}

// CreateDirectoryOutput defines the output for the CreateDirectory tool.
type CreateDirectoryOutput struct {
	Success bool `json:"success"`
}

// CreateDirectory creates a tool to create a new directory.
func CreateDirectory(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, CreateDirectoryTool, "Creates a new directory at the specified path.", func(ctx *ai.ToolContext, input CreateDirectoryInput) (CreateDirectoryOutput, error) {
		if err := os.MkdirAll(input.Path, 0755); err != nil {
			return CreateDirectoryOutput{Success: false}, fmt.Errorf("failed to create directory: %w", err)
		}
		return CreateDirectoryOutput{Success: true}, nil
	})
}

// DeleteDirectoryInput defines the input for the DeleteDirectory tool.
type DeleteDirectoryInput struct {
	Path string `json:"path"`
}

// DeleteDirectoryOutput defines the output for the DeleteDirectory tool.
type DeleteDirectoryOutput struct {
	Success bool `json:"success"`
}

// DeleteDirectory creates a tool to delete a directory.
func DeleteDirectory(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, DeleteDirectoryTool, "Deletes the directory at the specified path. Use with caution.", func(ctx *ai.ToolContext, input DeleteDirectoryInput) (DeleteDirectoryOutput, error) {
		// Basic safety check: prevent deleting root or empty path
		if input.Path == "/" || input.Path == "" || input.Path == "." {
			return DeleteDirectoryOutput{Success: false}, fmt.Errorf("cannot delete root or current directory")
		}

		// Check if it is a directory
		info, err := os.Stat(input.Path)
		if err != nil {
			return DeleteDirectoryOutput{Success: false}, fmt.Errorf("failed to stat path: %w", err)
		}
		if !info.IsDir() {
			return DeleteDirectoryOutput{Success: false}, fmt.Errorf("path is not a directory")
		}

		if err := os.RemoveAll(input.Path); err != nil {
			return DeleteDirectoryOutput{Success: false}, fmt.Errorf("failed to remove directory: %w", err)
		}
		return DeleteDirectoryOutput{Success: true}, nil
	})
}

// GetCurrentDirectoryInput defines the input for the GetCurrentDirectory tool.
type GetCurrentDirectoryInput struct{}

// GetCurrentDirectoryOutput defines the output for the GetCurrentDirectory tool.
type GetCurrentDirectoryOutput struct {
	Path string `json:"path"`
}

// GetCurrentDirectory creates a tool to get the current working directory.
func GetCurrentDirectory(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, GetCurrentDirectoryTool, "Gets the current working directory.", func(ctx *ai.ToolContext, _ GetCurrentDirectoryInput) (GetCurrentDirectoryOutput, error) {
		dir, err := os.Getwd()
		if err != nil {
			return GetCurrentDirectoryOutput{}, fmt.Errorf("failed to get current directory: %w", err)
		}
		return GetCurrentDirectoryOutput{Path: dir}, nil
	})
}

// WalkDirectoryInput defines the input for the WalkDirectory tool.
type WalkDirectoryInput struct {
	Path string `json:"path"`
}

// WalkDirectoryOutput defines the output for the WalkDirectory tool.
type WalkDirectoryOutput struct {
	Files []string `json:"files"`
}

// WalkDirectory creates a tool to recursively list all files in a directory.
func WalkDirectory(g *genkit.Genkit) ai.Tool {
	return genkit.DefineTool(g, WalkDirectoryTool, "Recursively lists all files in the specified directory.", func(ctx *ai.ToolContext, input WalkDirectoryInput) (WalkDirectoryOutput, error) {
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
			return WalkDirectoryOutput{}, fmt.Errorf("failed to walk directory: %w", err)
		}
		return WalkDirectoryOutput{Files: files}, nil
	})
}
