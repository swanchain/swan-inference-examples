// Minimal example: call Swan Inference from Go (same language as PicoClaw).
// Works on edge devices, ARM, RISC-V — anywhere Go compiles.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	baseURL = "https://inference.swanchain.io/v1/chat/completions"
	model   = "zai-org/GLM-4.7-Flash"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model     string    `json:"model"`
	Messages  []message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
}

func main() {
	apiKey := os.Getenv("SWAN_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Set SWAN_API_KEY environment variable")
		os.Exit(1)
	}

	prompt := "What is decentralized AI inference? Reply in 2 sentences."
	if len(os.Args) > 1 {
		prompt = os.Args[1]
	}

	body, _ := json.Marshal(request{
		Model:     model,
		Messages:  []message{{Role: "user", Content: prompt}},
		MaxTokens: 200,
	})

	req, _ := http.NewRequest("POST", baseURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "HTTP %d: %s\n", resp.StatusCode, string(data))
		os.Exit(1)
	}

	var res response
	json.Unmarshal(data, &res)

	if len(res.Choices) > 0 {
		fmt.Println(res.Choices[0].Message.Content)
		fmt.Printf("\n[%d in + %d out tokens, %dms]\n",
			res.Usage.PromptTokens, res.Usage.CompletionTokens,
			time.Since(start).Milliseconds())
	}
}
