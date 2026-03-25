# PicoClaw + GLM-4.7-Flash on Swan Inference

Run AI inference on **edge devices** (ARM, RISC-V, x86) using [PicoClaw](https://github.com/sipeed/picoclaw) and [GLM-4.7-Flash](https://huggingface.co/zai-org/GLM-4.7-Flash) served by Swan Inference's decentralized GPU network.

No local GPU needed — the heavy model runs on Swan Chain's provider network. Your edge device just sends HTTP requests.

## What's in This Example

| File | Description |
|------|-------------|
| `main.go` | Standalone Go client (~80 lines, zero dependencies) — builds for any platform |
| `config.json` | PicoClaw agent config pointing to Swan Inference |
| `Makefile` | Cross-compile for ARM, RISC-V, x86 |

## Standalone Go Client (No PicoClaw Required)

A minimal Go binary that calls Swan Inference. Cross-compiles to ARM/RISC-V for edge deployment.

### Build

```bash
# Current platform
go build -o swan-chat .

# Cross-compile for edge devices
GOOS=linux GOARCH=arm64 go build -o swan-chat-arm64 .        # Raspberry Pi 4/5, Jetson
GOOS=linux GOARCH=arm GOARM=7 go build -o swan-chat-armv7 .  # Raspberry Pi 3, older ARM
GOOS=linux GOARCH=riscv64 go build -o swan-chat-riscv64 .    # RISC-V boards (Sipeed, StarFive)
```

Binary size: ~7MB (no CGO, no runtime dependencies).

### Configure

Edit `config.json` with your API key:

```json
{
  "base_url": "https://inference.swanchain.io/v1",
  "api_key": "sk-swan-YOUR-API-KEY",
  "model": "zai-org/GLM-4.7-Flash",
  "max_tokens": 200
}
```

Config is loaded from (in priority order):
1. `./config.json` (current directory)
2. `~/.swan-chat/config.json` (home directory)
3. Environment variables (`SWAN_API_KEY`, `SWAN_BASE_URL`, `SWAN_MODEL`)

### Run

```bash
./swan-chat "What is decentralized AI?"
./swan-chat "Translate 'hello world' to Chinese"
```

Or use env vars instead of config file:

```bash
export SWAN_API_KEY=sk-swan-YOUR-API-KEY
./swan-chat "Hello"
```

### Deploy to Edge Device

```bash
# Build for ARM64
GOOS=linux GOARCH=arm64 go build -o swan-chat-arm64 .

# Copy binary + config to Raspberry Pi
scp swan-chat-arm64 config.json pi@raspberrypi:~/

# Run on the Pi
ssh pi@raspberrypi './swan-chat-arm64 "Hello from edge"'
```

## PicoClaw Agent Setup

For the full PicoClaw AI agent experience (interactive terminal, tool use, multi-model):

### 1. Install PicoClaw (v0.2.3+)

```bash
# Linux x86_64
curl -LO https://github.com/sipeed/picoclaw/releases/download/v0.2.3/picoclaw_Linux_x86_64.tar.gz
tar xzf picoclaw_Linux_x86_64.tar.gz && sudo mv picoclaw /usr/local/bin/

# Linux ARM64 (Raspberry Pi 4/5, Jetson)
curl -LO https://github.com/sipeed/picoclaw/releases/download/v0.2.3/picoclaw_Linux_arm64.tar.gz
tar xzf picoclaw_Linux_arm64.tar.gz && sudo mv picoclaw /usr/local/bin/

# Linux RISC-V
curl -LO https://github.com/sipeed/picoclaw/releases/download/v0.2.3/picoclaw_Linux_riscv64.tar.gz

# macOS (Apple Silicon)
curl -LO https://github.com/sipeed/picoclaw/releases/download/v0.2.3/picoclaw_Darwin_arm64.tar.gz

# Debian/Ubuntu packages also available (.deb)
curl -LO https://github.com/sipeed/picoclaw/releases/download/v0.2.3/picoclaw_aarch64.deb
sudo dpkg -i picoclaw_aarch64.deb
```

All platforms: [github.com/sipeed/picoclaw/releases/tag/v0.2.3](https://github.com/sipeed/picoclaw/releases/tag/v0.2.3)

### 2. Configure

```bash
mkdir -p ~/.picoclaw
cp config.json ~/.picoclaw/config.json
# Edit and replace sk-swan-YOUR-API-KEY with your actual key
```

### 3. Run

```bash
picoclaw
```

## Configuration

The `config.json` connects PicoClaw to Swan Inference:

```json
{
  "model_list": [
    {
      "model_name": "glm-4.7-flash",
      "model": "openai/zai-org/GLM-4.7-Flash",
      "api_base": "https://inference.swanchain.io/v1",
      "api_key": "sk-swan-YOUR-API-KEY"
    }
  ]
}
```

- `openai/` prefix = OpenAI-compatible provider
- `api_base` points to Swan Inference
- Model ID `zai-org/GLM-4.7-Flash` matches the Swan Inference catalog

## Why This Combination?

| Component | Role |
|-----------|------|
| **PicoClaw** | 8MB Go binary, runs on any edge device |
| **GLM-4.7-Flash** | 30B MoE model (3B active), 131K context, MIT license |
| **Swan Inference** | Decentralized GPU providers serve the model — no local GPU needed |

Your edge device sends a lightweight HTTP request. Swan Chain's GPU providers handle the heavy inference. You get GPT-class AI on a $35 Raspberry Pi.

## Verified Test Results

Tested with PicoClaw v0.2.3 + GLM-4.7-Flash on Swan Inference (March 2026):

```
$ echo "What is Swan Chain in 2 sentences?" | picoclaw agent

🦞 Interactive mode (Ctrl+C to exit)
INF  Agent initialized  tools_count=13
INF  LLM requested tool calls  tools=["web_search"]
INF  Tool call: web_search({"count":5,"query":"Swan Chain blockchain"})
INF  Tool execution completed  duration_ms=789
INF  LLM response without tool calls (direct answer)  iterations=2

🦞 Swan Chain is a comprehensive AI blockchain infrastructure launched
in 2021 that provides decentralized storage, computing, bandwidth, and
payment solutions, positioning itself at the forefront of bridging
blockchain technology with the growing demand for accessible AI
computing resources.
```

Key observations:
- GLM-4.7-Flash autonomously used **tool calling** (web search) before answering
- 2 LLM iterations: tool call + final answer
- PicoClaw has 13 built-in tools available to the model
- Total response time: ~2 seconds including web search

### Go Client Test

```
$ ./swan-chat "Translate hello world to Chinese"
你好世界

[10 in + 3 out tokens, 176ms, model: zai-org/GLM-4.7-Flash]
```

## Available Models

| Model | Type | Best For |
|-------|------|----------|
| `zai-org/GLM-4.7-Flash` | LLM (30B MoE) | Fast reasoning, code, 131K context |
| `Qwen/Qwen2.5-7B-Instruct` | LLM (7B) | General chat, fastest responses |
| `TheDrummer/Cydonia-24B-v4.1` | LLM (24B) | Creative writing, roleplay |

Full list: [inference.swanchain.io/models](https://inference.swanchain.io/models)
