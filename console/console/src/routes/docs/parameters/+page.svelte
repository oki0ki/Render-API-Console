<script lang="ts">
        const params = [
                {
                        name: 'model',
                        type: 'string',
                        req: true,
                        default: '—',
                        desc: 'The model alias or full model ID to use. Example: "Mistral-Small-4" or "mistralai/mistral-small-4-119b-2603".',
                },
                {
                        name: 'messages',
                        type: 'array',
                        req: true,
                        default: '—',
                        desc: 'Array of message objects representing the conversation. Each item requires role (system | user | assistant | tool) and content.',
                },
                {
                        name: 'stream',
                        type: 'boolean',
                        req: false,
                        default: 'true',
                        desc: 'Enable streaming via Server-Sent Events. Set to false to receive a single JSON response object.',
                },
                {
                        name: 'temperature',
                        type: 'number',
                        req: false,
                        default: 'model default',
                        desc: 'Sampling temperature between 0 and 2. Higher = more creative and varied. Lower = more deterministic. Recommended range: 0.1–1.0.',
                },
                {
                        name: 'max_tokens',
                        type: 'integer',
                        req: false,
                        default: 'model default',
                        desc: 'Maximum number of tokens to generate in the response. The model may produce fewer tokens if it reaches a natural stop.',
                },
                {
                        name: 'tools',
                        type: 'array',
                        req: false,
                        default: '—',
                        desc: 'List of tool definitions in OpenAI format. Add x-endpoint to a tool\'s function object for automatic agentic execution.',
                },
                {
                        name: 'tool_choice',
                        type: 'string | object',
                        req: false,
                        default: '"auto" when tools provided',
                        desc: 'Controls whether the model uses tools. Values: "auto" (model decides), "none" (no tools), or { "type": "function", "function": { "name": "..." } } to force a specific tool.',
                },
        ];
</script>

<svelte:head><title>Request Parameters — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Request Parameters</h1>
        <p class="text-sm text-[#666] mb-8 leading-relaxed">
                Full reference for all supported fields in the <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">POST /v1/chat/completions</code> request body.
        </p>

        <!-- Params -->
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <div class="px-5 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
                        <span class="text-xs font-semibold text-[#888] uppercase tracking-wider">Request Body — JSON</span>
                </div>
                <div class="divide-y divide-[#f0f0f0]">
                        {#each params as p}
                                <div class="px-5 py-4 hover:bg-[#fafafa] transition-colors">
                                        <div class="flex flex-wrap items-center gap-2 mb-1.5">
                                                <code class="text-[13px] font-mono font-semibold text-[#0f0f0f]">{p.name}</code>
                                                <span class="text-xs text-[#aaa] font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">{p.type}</span>
                                                {#if p.req}
                                                        <span class="text-[10px] font-bold text-red-500 bg-red-50 border border-red-100 px-1.5 py-0.5 rounded-full uppercase tracking-wide">required</span>
                                                {:else}
                                                        <span class="text-[10px] font-medium text-[#aaa] bg-[#f4f4f4] border border-[#e8e8e8] px-1.5 py-0.5 rounded-full">optional</span>
                                                {/if}
                                                {#if p.default !== '—'}
                                                        <span class="text-[10px] text-[#888]">default: <code class="font-mono">{p.default}</code></span>
                                                {/if}
                                        </div>
                                        <p class="text-sm text-[#555] leading-relaxed">{p.desc}</p>
                                </div>
                        {/each}
                </div>
        </div>

        <!-- Message object reference -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Message object</h2>
        <p class="text-sm text-[#666] mb-3">Each item in the <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">messages</code> array has this shape:</p>

        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <div class="divide-y divide-[#f0f0f0]">
                        {#each [
                                { name: 'role',        type: 'string',           req: true,  desc: 'Message role: "system" | "user" | "assistant" | "tool"' },
                                { name: 'content',     type: 'string | null',    req: true,  desc: 'The message text. Can be null when tool_calls is present.' },
                                { name: 'tool_calls',  type: 'array',            req: false, desc: 'Tool calls made by the assistant (present on assistant messages when finish_reason is "tool_calls").' },
                                { name: 'tool_call_id',type: 'string',           req: false, desc: 'Required on tool role messages. Must match the tool_call id from the assistant message.' },
                                { name: 'name',        type: 'string',           req: false, desc: 'Optional name for the tool role message (the function name that was called).' },
                        ] as f}
                                <div class="px-5 py-3 flex flex-wrap items-center gap-3">
                                        <code class="text-[13px] font-mono font-semibold text-[#333]">{f.name}</code>
                                        <span class="text-xs text-[#aaa] font-mono">{f.type}</span>
                                        {#if f.req}
                                                <span class="text-[10px] font-bold text-red-500 bg-red-50 border border-red-100 px-1.5 py-0.5 rounded-full uppercase tracking-wide">required</span>
                                        {:else}
                                                <span class="text-[10px] text-[#aaa]">optional</span>
                                        {/if}
                                        <p class="w-full text-sm text-[#555] mt-0.5">{f.desc}</p>
                                </div>
                        {/each}
                </div>
        </div>

        <!-- Example full request body -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Example request body</h2>
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden">
                <div class="px-4 py-2.5 bg-[#f9f9f9] border-b border-[#e5e5e5]">
                        <span class="text-xs font-medium text-[#888]">Full example — all optional fields</span>
                </div>
                <pre class="p-4 text-sm font-mono text-[#1e293b] bg-[#f4f4f4] overflow-x-auto leading-relaxed"><code>{`{
  "model": "DeepSeek-V3.1",
  "messages": [
    { "role": "system",    "content": "You are a helpful assistant." },
    { "role": "user",      "content": "What is 2 + 2?" },
    { "role": "assistant", "content": "2 + 2 equals 4." },
    { "role": "user",      "content": "And 4 + 4?" }
  ],
  "stream": false,
  "temperature": 0.7,
  "max_tokens": 512
}`}</code></pre>
        </div>

</div>
