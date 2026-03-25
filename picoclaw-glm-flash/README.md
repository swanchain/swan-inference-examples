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

Binary size: ~6MB (no CGO, no runtime dependencies).

### Run

```bash
export SWAN_API_KEY=sk-swan-YOUR-API-KEY
./swan-chat "What is decentralized AI?"
./swan-chat "Translate 'hello world' to Chinese"
```

### Deploy to Edge Device

```bash
# Copy to Raspberry Pi
scp swan-chat-arm64 pi@raspberrypi:~/
ssh pi@raspberrypi 'SWAN_API_KEY=sk-swan-xxx ./swan-chat "Hello from edge"'
```

## PicoClaw Agent Setup

For the full PicoClaw AI agent experience (interactive terminal, tool use, multi-model):

### 1. Install PicoClaw

```bash
# Download binary
curl -LO https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw-linux-amd64
chmod +x picoclaw-linux-amd64
sudo mv picoclaw-linux-amd64 /usr/local/bin/picoclaw

# Or ARM64
curl -LO https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw-linux-arm64
```

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

## Available Models

| Model | Type | Best For |
|-------|------|----------|
| `zai-org/GLM-4.7-Flash` | LLM (30B MoE) | Fast reasoning, code, 131K context |
| `Qwen/Qwen2.5-7B-Instruct` | LLM (7B) | General chat, fastest responses |
| `TheDrummer/Cydonia-24B-v4.1` | LLM (24B) | Creative writing, roleplay |

Full list: [inference.swanchain.io/models](https://inference.swanchain.io/models)
