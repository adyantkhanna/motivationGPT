package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Adyant's Open AI GPT Test API"))
}

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
	fmt.Println(resp.Choices[0].Text)
	w.Write([]byte(resp.Choices[0].Text))
}
