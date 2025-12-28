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
	"github.com/snowmerak/useful-genkit/flows"
	"github.com/snowmerak/useful-genkit/models"
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

	if _, err := models.OllamaGptOss20b(g); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.OllamaQwen3Coder(g, 30); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.OllamaDevstralSmall2(g); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.OllamaMinistral3o14B(g); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}

	if _, err := models.GoogleAI(g, models.GoogleAIGemma3o4b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.GoogleAI(g, models.GoogleAIGemma3o12b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.GoogleAI(g, models.GoogleAIGemma3o27b); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.GoogleAI(g, models.GoogleAIGemini2o5Pro); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.GoogleAI(g, models.GoogleAIGemini2o5Flash); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}
	if _, err := models.GoogleAI(g, models.GoogleAIGemini2o5FlashLite); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}

	_ = prompts.TranslationPrompt(g)
	_ = prompts.WrapErrorPrompt(g)
	_ = prompts.LogPrismPrompt(g)

	_ = tools.GetCurrentTime(g)
	_ = tools.FindUsage(g)
	_ = tools.FindDefinition(g)
	_ = tools.FindStructs(g)
	// _ = tools.FindInterfaceImplementations(g)

	_ = tools.ListFiles(g)
	_ = tools.CreateDirectory(g)
	_ = tools.DeleteDirectory(g)
	_ = tools.GetCurrentDirectory(g)
	_ = tools.WalkDirectory(g)
	_ = tools.ReadFile(g)
	_ = tools.WriteFile(g)

	flows.TranslationFlow(g)
	flows.WrapGoErrorFlow(g)
	flows.LogPrismFlow(g)

	mux := http.NewServeMux()

	log.Println("Starting server on http://localhost:3400")

	// Start the server
	if err := server.Start(ctx, ":3400", mux); err != nil {
		log.Printf("Server error: %v\n", err)
	}

	log.Println("Shutting down gracefully...")
}
