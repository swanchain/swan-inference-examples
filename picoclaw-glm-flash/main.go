// Minimal example: call Swan Inference from Go (same language as PicoClaw).
// Works on edge devices, ARM, RISC-V — anywhere Go compiles.
//
// Config priority: config.json > environment variables > defaults
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Config loaded from config.json or environment
type Config struct {
	BaseURL   string `json:"base_url"`
	APIKey    string `json:"api_key"`
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
}

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
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func loadConfig() Config {
	cfg := Config{
		BaseURL:   "https://inference.swanchain.io/v1",
		Model:     "zai-org/GLM-4.7-Flash",
		MaxTokens: 200,
	}

	// Try loading config.json from current dir, then ~/.swan-chat/
	for _, path := range []string{
		"config.json",
		filepath.Join(homeDir(), ".swan-chat", "config.json"),
	} {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var fileCfg Config
		if json.Unmarshal(data, &fileCfg) == nil {
			if fileCfg.BaseURL != "" {
				cfg.BaseURL = fileCfg.BaseURL
			}
			if fileCfg.APIKey != "" {
				cfg.APIKey = fileCfg.APIKey
			}
			if fileCfg.Model != "" {
				cfg.Model = fileCfg.Model
			}
			if fileCfg.MaxTokens > 0 {
				cfg.MaxTokens = fileCfg.MaxTokens
			}
			break
		}
	}

	// Environment variables override config file
	if v := os.Getenv("SWAN_BASE_URL"); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv("SWAN_API_KEY"); v != "" {
		cfg.APIKey = v
	}
	if v := os.Getenv("SWAN_MODEL"); v != "" {
		cfg.Model = v
	}

	return cfg
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return "."
}

func main() {
	cfg := loadConfig()

	if cfg.APIKey == "" {
		fmt.Fprintln(os.Stderr, "No API key found. Set it via:")
		fmt.Fprintln(os.Stderr, "  1. config.json:  {\"api_key\": \"sk-swan-xxx\"}")
		fmt.Fprintln(os.Stderr, "  2. Environment:  export SWAN_API_KEY=sk-swan-xxx")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Get a key at https://inference.swanchain.io/signup")
		os.Exit(1)
	}

	prompt := "What is decentralized AI inference? Reply in 2 sentences."
	if len(os.Args) > 1 {
		prompt = os.Args[1]
	}

	body, _ := json.Marshal(request{
		Model:     cfg.Model,
		Messages:  []message{{Role: "user", Content: prompt}},
		MaxTokens: cfg.MaxTokens,
	})

	url := cfg.BaseURL + "/chat/completions"
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)

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

	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", res.Error.Message)
		os.Exit(1)
	}

	if len(res.Choices) > 0 {
		fmt.Println(res.Choices[0].Message.Content)
		fmt.Printf("\n[%d in + %d out tokens, %dms, model: %s]\n",
			res.Usage.PromptTokens, res.Usage.CompletionTokens,
			time.Since(start).Milliseconds(), cfg.Model)
	}
}
