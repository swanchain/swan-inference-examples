"""Streaming chat completion using Swan Inference."""
import os
from openai import OpenAI

client = OpenAI(
    base_url="https://inference.swanchain.io/v1",
    api_key=os.environ.get("SWAN_API_KEY", "sk-swan-YOUR-API-KEY"),
)

stream = client.chat.completions.create(
    model="zai-org/GLM-4.7-Flash",
    messages=[
        {"role": "user", "content": "Write a haiku about GPU computing."},
    ],
    stream=True,
    max_tokens=100,
)

for chunk in stream:
    content = chunk.choices[0].delta.content
    if content:
        print(content, end="", flush=True)
print()
