package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
	"github.com/firebase/genkit/go/plugins/server"
	provider "github.com/snowmerak/useful-genkit/models"
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
		genkit.WithPlugins(o),
	)

	if _, err := provider.OllamaGptOss20b(g); err != nil {
		log.Fatalf("Failed to get model: %v", err)
	}

	_ = tools.GetCurrentTime(g)

	mux := http.NewServeMux()

	log.Println("Starting server on http://localhost:3400")

	// Start the server
	if err := server.Start(ctx, ":3400", mux); err != nil {
		log.Printf("Server error: %v\n", err)
	}

	log.Println("Shutting down gracefully...")
}
