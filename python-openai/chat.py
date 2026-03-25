"""Basic chat completion using Swan Inference with OpenAI SDK."""
import os
from openai import OpenAI

client = OpenAI(
    base_url="https://inference.swanchain.io/v1",
    api_key=os.environ.get("SWAN_API_KEY", "sk-swan-YOUR-API-KEY"),
)

response = client.chat.completions.create(
    model="zai-org/GLM-4.7-Flash",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Explain decentralized AI inference in 3 sentences."},
    ],
    max_tokens=200,
)

print(response.choices[0].message.content)
print(f"\nTokens: {response.usage.prompt_tokens} in, {response.usage.completion_tokens} out")
