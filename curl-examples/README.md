# cURL Examples for Swan Inference

Raw HTTP examples for all Swan Inference endpoints. No dependencies needed.

## Setup

```bash
export SWAN_API_KEY=sk-swan-YOUR-API-KEY
export SWAN_URL=https://inference.swanchain.io
```

## Examples

### Chat Completion

```bash
curl $SWAN_URL/v1/chat/completions \
  -H "Authorization: Bearer $SWAN_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zai-org/GLM-4.7-Flash",
    "messages": [{"role": "user", "content": "Hello!"}],
    "max_tokens": 100
  }'
```

### Streaming

```bash
curl $SWAN_URL/v1/chat/completions \
  -H "Authorization: Bearer $SWAN_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zai-org/GLM-4.7-Flash",
    "messages": [{"role": "user", "content": "Count to 10"}],
    "stream": true,
    "max_tokens": 100
  }'
```

### List Models

```bash
curl $SWAN_URL/v1/models \
  -H "Authorization: Bearer $SWAN_API_KEY"
```

### Available Models (online providers only)

```bash
curl $SWAN_URL/v1/available-models
```

### Health Check

```bash
curl $SWAN_URL/v1/health
```

### Public Playground (no API key)

```bash
curl $SWAN_URL/v1/playground/chat \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zai-org/GLM-4.7-Flash",
    "messages": [{"role": "user", "content": "What is Swan Chain?"}]
  }'
```

### Playground Models

```bash
curl $SWAN_URL/v1/playground/models
```
