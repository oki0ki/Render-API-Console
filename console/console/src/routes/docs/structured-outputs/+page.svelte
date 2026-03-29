<script lang="ts">
        import { baseUrl } from '$lib/docs-store';
        import CodeBlock from '$lib/CodeBlock.svelte';

        $: url = $baseUrl || 'https://api.example.com/v1';

        $: jsonModeCurl = `curl ${url}/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -d '{
    "model": "DeepSeek-V3.1",
    "messages": [
      {
        "role": "system",
        "content": "You are a data extraction assistant. Always respond with valid JSON."
      },
      {
        "role": "user",
        "content": "Extract name and price from: \\"MacBook Pro 16\\" costs $2499\\""
      }
    ],
    "response_format": { "type": "json_object" }
  }'`;

        $: jsonModePython = `from openai import OpenAI
import json

client = OpenAI(
    api_key="sk-<your-key>",
    base_url="${url}"
)

response = client.chat.completions.create(
    model="DeepSeek-V3.1",
    messages=[
        {
            "role": "system",
            "content": "You are a data extraction assistant. Always respond with valid JSON."
        },
        {
            "role": "user",
            "content": 'Extract name and price from: "MacBook Pro 16\\" costs $2499"'
        }
    ],
    response_format={"type": "json_object"}
)

data = json.loads(response.choices[0].message.content)
print(data)
# → {"name": "MacBook Pro 16", "price": 2499}`;

        $: jsonModeJS = `import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-<your-key>',
  baseURL: '${url}'
});

const response = await client.chat.completions.create({
  model: 'DeepSeek-V3.1',
  messages: [
    {
      role: 'system',
      content: 'You are a data extraction assistant. Always respond with valid JSON.'
    },
    {
      role: 'user',
      content: 'Extract name and price from: "MacBook Pro 16" costs $2499'
    }
  ],
  response_format: { type: 'json_object' }
});

const data = JSON.parse(response.choices[0].message.content);
console.log(data);
// → { name: 'MacBook Pro 16', price: 2499 }`;

        $: jsonSchemaExample = `{
  "model": "DeepSeek-V3.1",
  "messages": [
    { "role": "user", "content": "List three planets with their distance from the Sun in AU." }
  ],
  "response_format": {
    "type": "json_object"
  }
}
// Prompt the model to follow your schema in the system message:
// "Respond ONLY with a JSON object: { \\"planets\\": [{ \\"name\\": string, \\"distance_au\\": number }] }"`;

        let tab = 'python';
        const codeMap: Record<string, () => string> = {};
        $: codeMap['curl']   = () => jsonModeCurl;
        $: codeMap['python'] = () => jsonModePython;
        $: codeMap['js']     = () => jsonModeJS;
</script>

<svelte:head><title>Structured Outputs — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Structured Outputs</h1>
        <p class="text-sm text-[#666] mb-8 leading-relaxed">
                Force the model to respond with valid JSON using <code class="font-mono bg-[#f4f4f4] px-1 rounded">response_format</code>.
                This is ideal for data extraction, classification, and any workflow that must parse the model's reply programmatically.
        </p>

        <!-- What it does -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-4">How it works</h2>
        <div class="space-y-3 mb-8">
                {#each [
                        { step: '1', title: 'Set response_format', desc: 'Pass { "type": "json_object" } in your request body.' },
                        { step: '2', title: 'Instruct in the system prompt', desc: 'Tell the model exactly which JSON shape you expect — field names, types, and nesting.' },
                        { step: '3', title: 'Parse the response', desc: 'message.content will always be a valid JSON string. Parse it safely with JSON.parse() or json.loads().' },
                ] as s}
                        <div class="flex items-start gap-4 p-4 rounded-2xl border border-[#e5e5e5] bg-[#fafafa]">
                                <div class="flex-shrink-0 w-7 h-7 rounded-full bg-[#0f0f0f] text-white text-xs font-bold flex items-center justify-center">{s.step}</div>
                                <div>
                                        <p class="text-sm font-semibold text-[#111] mb-0.5">{s.title}</p>
                                        <p class="text-sm text-[#555]">{s.desc}</p>
                                </div>
                        </div>
                {/each}
        </div>

        <!-- response_format parameter -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">response_format parameter</h2>
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <table class="w-full text-sm">
                        <thead>
                                <tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Field</th>
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Type</th>
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Description</th>
                                </tr>
                        </thead>
                        <tbody class="divide-y divide-[#f0f0f0]">
                                {#each [
                                        { field: 'type', type: '"text" | "json_object"', desc: 'Set to "json_object" to enable JSON mode. Defaults to "text".' },
                                ] as row}
                                        <tr class="hover:bg-[#fafafa] transition-colors">
                                                <td class="px-4 py-3 font-mono text-[13px] text-[#333]">{row.field}</td>
                                                <td class="px-4 py-3"><span class="text-xs font-mono px-2 py-0.5 rounded bg-[#f4f4f4] text-[#555]">{row.type}</span></td>
                                                <td class="px-4 py-3 text-[13px] text-[#666]">{row.desc}</td>
                                        </tr>
                                {/each}
                        </tbody>
                </table>
        </div>

        <!-- Code example with tabs -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Code example — data extraction</h2>

        <div class="border border-[#e5e5e5] rounded-2xl overflow-hidden mb-8">
                <div class="flex border-b border-[#e5e5e5] bg-[#f9f9f9]">
                        {#each [['python','Python'], ['js','JavaScript'], ['curl','cURL']] as [id, label]}
                                <button
                                        class="px-4 py-2.5 text-xs font-semibold transition-colors
                                                {tab === id ? 'text-[#0f0f0f] bg-white' : 'text-[#888] hover:text-[#555]'}"
                                        on:click={() => tab = id}
                                >{label}</button>
                        {/each}
                </div>
                {#if tab === 'python'}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{jsonModePython}</code></pre>
                {:else if tab === 'js'}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{jsonModeJS}</code></pre>
                {:else}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{jsonModeCurl}</code></pre>
                {/if}
        </div>

        <!-- Schema tip -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Defining your schema in the system prompt</h2>
        <p class="text-sm text-[#666] mb-4">
                The model follows schema instructions from the system message. Be explicit about field names and types:
        </p>
        <CodeBlock code={jsonSchemaExample} id="json-schema" lang="json" />

        <!-- Best practices -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Tips</h2>
        <div class="space-y-2 mb-8">
                {#each [
                        { icon: '✓', text: 'Always include the schema in the system message — don\'t rely on the model to guess your format.' },
                        { icon: '✓', text: 'Wrap JSON.parse() in a try/catch; some models occasionally output near-valid JSON on edge cases.' },
                        { icon: '✓', text: 'Use DeepSeek-V3.1 or Mistral-Small-4 for best JSON adherence in structured extraction tasks.' },
                        { icon: '✓', text: 'For deeply nested schemas, describe each field with its type and a short description in the prompt.' },
                        { icon: '⚠', text: 'JSON mode only guarantees valid JSON — it does not enforce a specific schema. Validate fields yourself.' },
                ] as tip}
                        <div class="flex items-start gap-3 p-3 rounded-xl bg-[#fafafa] border border-[#f0f0f0]">
                                <span class="{tip.icon === '⚠' ? 'text-amber-500' : 'text-green-500'} font-bold text-sm flex-shrink-0">{tip.icon}</span>
                                <p class="text-sm text-[#555]">{tip.text}</p>
                        </div>
                {/each}
        </div>

</div>
