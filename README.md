# Swan Inference Examples

Code examples showing how to use [Swan Inference](https://inference.swanchain.io) — a decentralized AI inference marketplace on Swan Chain.

Swan Inference provides an **OpenAI-compatible API** backed by a decentralized network of GPU providers. Use any OpenAI SDK, Claw-family AI agent, or HTTP client.

## Quick Start

```bash
# No signup needed — try the public playground
curl https://inference.swanchain.io/v1/playground/chat \
  -H "Content-Type: application/json" \
  -d '{"model":"glm-4.7-flash","messages":[{"role":"user","content":"Hello!"}]}'
```

For full access, [sign up](https://inference.swanchain.io/signup) and get an API key (`sk-swan-*`).

## Examples

| Example | Language | Description |
|---------|----------|-------------|
| [picoclaw-glm-flash](picoclaw-glm-flash/) | Go (PicoClaw) | AI coding agent powered by GLM-4.7-Flash via Swan Inference |
| [python-openai](python-openai/) | Python | OpenAI SDK with Swan Inference backend |
| [curl-examples](curl-examples/) | Shell | Raw HTTP examples for all endpoints |
| [nodejs-openai](nodejs-openai/) | JavaScript | Node.js OpenAI SDK integration |

## Available Models

| Model | Type | Best For |
|-------|------|----------|
| `glm-4.7-flash` | LLM (30B MoE) | Fast reasoning, code, 131K context |
| `deepseek-r1-distill-llama-70b` | LLM | Deep reasoning, math |
| `Qwen/Qwen2.5-7B-Instruct` | LLM | General chat, fast |
| `meta-llama/Llama-3.3-70B-Instruct` | LLM | General purpose |

Full model list: [inference.swanchain.io/models](https://inference.swanchain.io/models)

## Base URL

```
https://inference.swanchain.io/v1
```

Drop-in replacement for `https://api.openai.com/v1` — just change the base URL and API key.

## Links

- [Swan Inference Dashboard](https://inference.swanchain.io)
- [API Documentation](https://docs.swanchain.io/bulders/app-developer/swan-inference-api)
- [Claw Tools Integration Guide](https://docs.swanchain.io/bulders/app-developer/claw-tools-integration)
- [Become a GPU Provider](https://inference.swanchain.io/provider-guide)
