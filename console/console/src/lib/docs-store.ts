import { writable } from 'svelte/store';

export const baseUrl = writable('');
export const copiedId = writable<string | null>(null);

export async function copyText(id: string, text: string) {
        try {
                await navigator.clipboard.writeText(text);
                copiedId.set(id);
                setTimeout(() => copiedId.update(v => (v === id ? null : v)), 2000);
        } catch {}
}

export function getSnippets(u: string) {
        const url = u || 'https://api.example.com/v1';
        const O = '{';
        const C = '}';
        return {
                curlModels: `curl ${url}/models \\
  -H "Authorization: Bearer sk-<your-key>"`,

                curlModelsResponse: `{
  "object": "list",
  "data": [
    { "id": "Mistral-Small-4", "object": "model", "created": 1743170400, "owned_by": "nvidia" },
    { "id": "DeepSeek-V3.1",   "object": "model", "created": 1743170400, "owned_by": "nvidia" },
    { "id": "Kimi-K2",         "object": "model", "created": 1743170400, "owned_by": "nvidia" }
  ]
}`,

                curlChat: `curl ${url}/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -d '{
    "model": "DeepSeek-V3.1",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "What is the capital of Poland?"}
    ]
  }'`,

                quickstartPython: `from openai import OpenAI

client = OpenAI(
    api_key="sk-<your-key>",
    base_url="${url}"
)

response = client.chat.completions.create(
    model="Mistral-Small-4",
    messages=[
        {"role": "user", "content": "Hello! What can you do?"}
    ]
)

print(response.choices[0].message.content)`,

                quickstartJS: `import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-<your-key>',
  baseURL: '${url}'
});

const response = await client.chat.completions.create({
  model: 'Mistral-Small-4',
  messages: [{ role: 'user', content: 'Hello! What can you do?' }]
});

console.log(response.choices[0].message.content);`,

                quickstartCurl: `curl ${url}/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -d '{"model":"Mistral-Small-4","messages":[{"role":"user","content":"Hello!"}]}'`,

                python: `from openai import OpenAI

client = OpenAI(
    api_key="sk-<your-key>",
    base_url="${url}"
)

response = client.chat.completions.create(
    model="Kimi-K2",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Explain quantum entanglement simply."}
    ],
    temperature=0.7,
    max_tokens=1024
)

print(response.choices[0].message.content)`,

                js: `import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-<your-key>',
  baseURL: '${url}'
});

const response = await client.chat.completions.create({
  model: 'Mistral-Small-4',
  messages: [
    { role: 'system', content: 'You are a helpful assistant.' },
    { role: 'user',   content: 'Explain quantum entanglement simply.' }
  ],
  temperature: 0.7,
  maxTokens: 1024
});

console.log(response.choices[0].message.content);`,

                responseExample: `{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "created": 1743170400,
  "model": "Kimi-K2",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Quantum entanglement is a phenomenon..."
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 28,
    "completion_tokens": 142,
    "total_tokens": 170
  }
}`,

                streamPython: `# Using the openai SDK — streaming handled automatically
stream = client.chat.completions.create(
    model="GPT-OSS-120b",
    messages=[{"role": "user", "content": "Tell me a story"}],
    stream=True
)
for chunk in stream:
    print(chunk.choices[0].delta.content or "", end="", flush=True)`,

                streamJS: `// Using the openai SDK — streaming handled automatically
const stream = await client.chat.completions.create({
  model: 'GPT-OSS-120b',
  messages: [{ role: 'user', content: 'Tell me a story' }],
  stream: true
});

for await (const chunk of stream) {
  const delta = chunk.choices[0]?.delta?.content ?? '';
  process.stdout.write(delta);
}`,

                streamCurl: `curl ${url}/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-<your-key>" \\
  --no-buffer \\
  -d '{
    "model": "GPT-OSS-120b",
    "messages": [{"role":"user","content":"Tell me a story"}],
    "stream": true
  }'`,

                sseRaw: `data: {"id":"chatcmpl-...","choices":[{"delta":{"role":"assistant","content":""},...}]}

data: {"id":"chatcmpl-...","choices":[{"delta":{"content":"Once"},...}]}

data: {"id":"chatcmpl-...","choices":[{"delta":{"content":" upon"},...}]}

data: {"id":"chatcmpl-...","choices":[{"delta":{},"finish_reason":"stop"}]}

data: [DONE]`,

                toolCurl: `curl ${url}/chat/completions \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -H "Content-Type: application/json" \\
  -d '{
    "model": "DeepSeek-V3.1",
    "messages": [{"role":"user","content":"What is the weather in Warsaw?"}],
    "tools": [{
      "type": "function",
      "function": {
        "name": "get_weather",
        "description": "Get current weather for a city",
        "parameters": {
          "type": "object",
          "properties": {"city": {"type": "string"}},
          "required": ["city"]
        }
      }
    }],
    "tool_choice": "auto"
  }'`,

                toolResponse: `{
  "id": "chatcmpl-abc",
  "choices": [{
    "message": {
      "role": "assistant",
      "content": null,
      "tool_calls": [{
        "id": "call_abc123",
        "type": "function",
        "function": {
          "name": "get_weather",
          "arguments": "{\\"city\\": \\"Warsaw\\"}"
        }
      }]
    },
    "finish_reason": "tool_calls"
  }]
}`,

                toolFollow: `# After executing the tool, send results back to the model:
messages.append({"role": "assistant", "tool_calls": [...]})
messages.append({
    "role": "tool",
    "tool_call_id": "call_abc123",
    "name": "get_weather",
    "content": "{\\"temperature\\": 12, \\"condition\\": \\"Cloudy\\"}"
})

# Continue the conversation
response = client.chat.completions.create(
    model="DeepSeek-V3.1",
    messages=messages
)`,

                agenticCurl: `curl ${url}/chat/completions \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -H "Content-Type: application/json" \\
  -d '{
    "model": "Mistral-Small-4",
    "messages": [{"role":"user","content":"What is the EUR/PLN exchange rate?"}],
    "tools": [{
      "type": "function",
      "function": {
        "name": "search_web",
        "description": "Search the web for current information",
        "parameters": {
          "type": "object",
          "properties": {"query": {"type": "string"}},
          "required": ["query"]
        },
        "x-endpoint": "https://your-service.com/api/search"
      }
    }]
  }'`,

                agenticEndpoint: `// Your endpoint receives a POST with tool arguments:
// POST https://your-service.com/api/search
// {"query": "EUR/PLN exchange rate today"}

// Return any JSON — it gets forwarded to the model:
{
  "result": "EUR/PLN rate is 4.27 (as of 2026-03-28)",
  "source": "European Central Bank"
}`,

                retryPython: `import time
from openai import OpenAI

client = OpenAI(api_key="sk-<your-key>", base_url="${url}")

def chat_with_retry(model, messages, max_retries=3):
    for attempt in range(max_retries):
        try:
            return client.chat.completions.create(
                model=model, messages=messages
            )
        except Exception as e:
            if '502' not in str(e) or attempt == max_retries - 1:
                raise
            wait = 2 ** attempt   # 1s → 2s → 4s
            print(f"Retry {attempt+1} after {wait}s...")
            time.sleep(wait)`,
        };
}

export const NAV = [
        {
                group: 'Overview',
                items: [
                        { id: 'introduction', label: 'Introduction', path: '/docs/introduction' },
                        { id: 'quickstart',   label: 'Quick Start',  path: '/docs/quickstart' },
                        { id: 'changelog',    label: 'Changelog',    path: '/docs/changelog' },
                ],
        },
        {
                group: 'Reference',
                items: [
                        { id: 'authentication',   label: 'Authentication',     path: '/docs/authentication' },
                        { id: 'models',           label: 'Available Models',   path: '/docs/models' },
                        { id: 'chat-completions', label: 'Chat Completions',   path: '/docs/chat-completions' },
                        { id: 'parameters',       label: 'Request Parameters', path: '/docs/parameters' },
                        { id: 'response-format',  label: 'Response Format',    path: '/docs/response-format' },
                        { id: 'streaming',        label: 'Streaming',          path: '/docs/streaming' },
                        { id: 'rate-limits',      label: 'Rate Limits',        path: '/docs/rate-limits' },
                ],
        },
        {
                group: 'Advanced',
                items: [
                        { id: 'structured-outputs', label: 'Structured Outputs', path: '/docs/structured-outputs' },
                        { id: 'tool-calls',         label: 'Tool Calls',         path: '/docs/tool-calls' },
                        { id: 'agentic-loop',       label: 'Agentic Loop',       path: '/docs/agentic-loop' },
                        { id: 'errors',             label: 'Error Handling',     path: '/docs/errors' },
                        { id: 'best-practices',     label: 'Best Practices',     path: '/docs/best-practices' },
                ],
        },
];

export const FLAT_NAV = NAV.flatMap(g => g.items);
