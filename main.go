package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/firebase/genkit/go/plugins/ollama"
	"github.com/firebase/genkit/go/plugins/server"
	provider "github.com/snowmerak/useful-genkit/models"
	"github.com/snowmerak/useful-genkit/prompts"
	"github.com/snowmerak/useful-genkit/tools"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	o := &ollama.Ollama{
		ServerAddress: "http://localhost:11434",
		Timeout:       300,
	}

	g := genkit.Init(ctx,
		genkit.WithPlugins(o, &googlegenai.GoogleAI{
			APIKey: os.Getenv("GEMINI_API_KEY"),
		}),
	)

	if _, err := provider.OllamaGptOss20b(g); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.OllamaGemma3(g, 12); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.OllamaQwen3(g, 14); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}

	if _, err := provider.GoogleAI(g, provider.GoogleAIGemma3o4b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.GoogleAI(g, provider.GoogleAIGemma3o12b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.GoogleAI(g, provider.GoogleAIGemma3o27b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.GoogleAI(g, provider.GoogleAIGemini2o5Pro); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.GoogleAI(g, provider.GoogleAIGemini2o5Flash); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := provider.GoogleAI(g, provider.GoogleAIGemini2o5FlashLite); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}

	_ = prompts.TranslationPrompt(g)

	_ = tools.GetCurrentTime(g)

	mux := http.NewServeMux()

	log.Println("Starting server on http://localhost:3400")

	// Start the server
	if err := server.Start(ctx, ":3400", mux); err != nil {
		log.Printf("Server error: %v\n", err)
	}

	log.Println("Shutting down gracefully...")
}
