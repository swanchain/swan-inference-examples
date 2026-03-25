import OpenAI from "openai";

const client = new OpenAI({
  baseURL: "https://inference.swanchain.io/v1",
  apiKey: process.env.SWAN_API_KEY || "sk-swan-YOUR-API-KEY",
});

const response = await client.chat.completions.create({
  model: "zai-org/GLM-4.7-Flash",
  messages: [
    { role: "system", content: "You are a helpful assistant." },
    { role: "user", content: "What is decentralized AI inference?" },
  ],
  max_tokens: 200,
});

console.log(response.choices[0].message.content);
console.log(`\nTokens: ${response.usage.prompt_tokens} in, ${response.usage.completion_tokens} out`);
