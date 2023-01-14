package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// health is a simple health check handler.
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Adyant's Open AI GPT Test API"))
	p := "Hello from Adyant's Open AI GPT Test API"
	json.NewEncoder(w).Encode(p)
}

// openAI is a simple handler that calls the Open AI API
func openAI(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	prompt := "Tell me a different inspirational quote"

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   500,
		Prompt:      prompt,
		Temperature: 1,
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	out := GPTResponse{Response: resp.Choices[0].Text[2:], Good: true}
	// fmt.Println(resp.Choices[0].Text[2:])
	// w.Write([]byte(resp.Choices[0].Text[2:]))
	json.NewEncoder(w).Encode(out)
}
