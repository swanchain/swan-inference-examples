# PicoClaw + GLM-4.7-Flash on Swan Inference

Run [PicoClaw](https://github.com/sipeed/picoclaw) (lightweight Go AI agent) with [GLM-4.7-Flash](https://huggingface.co/zai-org/GLM-4.7-Flash) served by Swan Inference's decentralized GPU network.

## What You Get

- AI coding assistant in your terminal
- GLM-4.7-Flash: 30B MoE model (3B active params), 131K context window
- Powered by decentralized GPU providers on Swan Chain
- No local GPU required

## Setup

### 1. Install PicoClaw

```bash
# Download binary (Linux amd64)
curl -LO https://github.com/sipeed/picoclaw/releases/latest/download/picoclaw-linux-amd64
chmod +x picoclaw-linux-amd64
sudo mv picoclaw-linux-amd64 /usr/local/bin/picoclaw

# Or from source
git clone https://github.com/sipeed/picoclaw.git
cd picoclaw && make build && sudo make install
```

### 2. Get a Swan Inference API Key

Sign up at [inference.swanchain.io/signup](https://inference.swanchain.io/signup) to get your `sk-swan-*` API key.

### 3. Configure

Copy the config file:

```bash
mkdir -p ~/.picoclaw
cp config.json ~/.picoclaw/config.json
```

Edit `~/.picoclaw/config.json` and replace `sk-swan-YOUR-API-KEY` with your actual key.

### 4. Run

```bash
picoclaw
```

PicoClaw will start an interactive session using GLM-4.7-Flash via Swan Inference.

## Configuration

The `config.json` connects PicoClaw to Swan Inference:

```json
{
  "agents": {
    "defaults": {
      "model_name": "glm-4.7-flash",
      "max_tokens": 8192,
      "temperature": 0.7
    }
  },
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

Key points:
- `openai/` prefix tells PicoClaw to use the OpenAI-compatible provider
- `api_base` points to Swan Inference
- The model ID `zai-org/GLM-4.7-Flash` matches the Swan Inference catalog

## Try Other Models

Swan Inference serves many models. Change the `model` field:

```json
{
  "model_name": "cydonia-24b",
  "model": "openai/TheDrummer/Cydonia-24B-v4.1",
  "api_base": "https://inference.swanchain.io/v1",
  "api_key": "sk-swan-YOUR-API-KEY"
}
```

Available models: [inference.swanchain.io/models](https://inference.swanchain.io/models)

## Why Swan Inference?

- No GPU required locally
- Models served by decentralized GPU providers
- OpenAI-compatible API (works with any tool)
- Pay-per-token or $6/month Pro subscription
- Providers earn rewards for serving inference
