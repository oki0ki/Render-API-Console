<script lang="ts">
        import { baseUrl, getSnippets } from '$lib/docs-store';
        import CodeBlock from '$lib/CodeBlock.svelte';

        $: s = getSnippets($baseUrl);
        let activeTab = 0;
</script>

<svelte:head><title>Streaming — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Streaming</h1>
        <p class="text-sm text-[#666] mb-8 leading-relaxed">
                Responses are streamed by default using Server-Sent Events (SSE). Tokens are sent as they are generated, enabling low-latency user experiences. Set <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">stream: false</code> to receive a single JSON object instead.
        </p>

        <!-- Note -->
        <div class="flex items-start gap-3 p-4 rounded-xl border border-[#e5e5e5] bg-[#f4f4f4] mb-8">
                <svg class="w-4 h-4 text-[#888] flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
                <p class="text-sm text-[#444]">
                        The OpenAI SDK handles SSE streaming transparently. You don't need to parse raw SSE manually — just use <code class="bg-[#e5e5e5] px-1 rounded font-mono">stream=True</code> / <code class="bg-[#e5e5e5] px-1 rounded font-mono">stream: true</code>.
                </p>
        </div>

        <!-- SDK examples -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">SDK streaming examples</h2>
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <div class="flex border-b border-[#e5e5e5] bg-[#f9f9f9]">
                        {#each ['Python', 'TypeScript', 'cURL'] as tab, i}
                                <button
                                        on:click={() => activeTab = i}
                                        class="px-4 py-2.5 text-xs font-medium transition-colors {activeTab === i ? 'text-[#0f0f0f] font-semibold bg-white' : 'text-[#888] hover:text-[#444]'}"
                                >{tab}</button>
                        {/each}
                </div>
                {#if activeTab === 0}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.streamPython}</code></pre>
                {:else if activeTab === 1}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.streamJS}</code></pre>
                {:else}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.streamCurl}</code></pre>
                {/if}
        </div>

        <!-- SSE wire format -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Raw SSE wire format</h2>
        <p class="text-sm text-[#666] mb-3">Each chunk is a <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">data: ...</code> line followed by two newlines. The stream terminates with <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">data: [DONE]</code>.</p>
        <CodeBlock code={s.sseRaw} id="stream-sse" lang="text" />

        <!-- Chunk structure -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Chunk structure</h2>
        <p class="text-sm text-[#666] mb-3">Each streamed chunk has this shape:</p>
        <CodeBlock code={`{
  "id": "chatcmpl-abc",
  "object": "chat.completion.chunk",
  "created": 1743170400,
  "model": "GPT-OSS-120b",
  "choices": [
    {
      "index": 0,
      "delta": {
        "role": "assistant",   // present only in the first chunk
        "content": "Once"      // the next token(s)
      },
      "finish_reason": null    // "stop" or "tool_calls" in the last chunk
    }
  ]
}`} id="stream-chunk" lang="json" />

        <!-- Disable streaming -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Disable streaming</h2>
        <p class="text-sm text-[#666] mb-3">Set <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">stream: false</code> to receive the entire response as a single JSON object:</p>
        <CodeBlock code={`curl ${$baseUrl || 'https://api.example.com/v1'}/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-<your-key>" \\
  -d '{
    "model": "Mistral-Small-4",
    "messages": [{"role": "user", "content": "Hello"}],
    "stream": false
  }'`} id="stream-off" lang="bash" />

</div>
