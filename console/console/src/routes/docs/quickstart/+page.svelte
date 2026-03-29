<script lang="ts">
        import { baseUrl, getSnippets } from '$lib/docs-store';

        $: s = getSnippets($baseUrl);
        let activeTab1 = 0;
        let activeTab2 = 0;
</script>

<svelte:head><title>Quick Start — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Quick Start</h1>
        <p class="text-sm text-[#666] mb-8 leading-relaxed">
                Make your first API call in under a minute. The gateway is fully compatible with the OpenAI SDK — just change the base URL and API key.
        </p>

        <!-- Step 1 -->
        <div class="flex items-start gap-4 mb-8">
                <div class="flex-shrink-0 w-7 h-7 rounded-full bg-[#0f0f0f] text-white text-xs font-bold flex items-center justify-center mt-0.5">1</div>
                <div class="flex-1">
                        <h2 class="text-base font-semibold text-[#0f0f0f] mb-1">Install the OpenAI SDK</h2>
                        <p class="text-sm text-[#666] mb-3">The gateway is a drop-in replacement — use any existing OpenAI SDK.</p>
                        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden">
                                <div class="flex border-b border-[#e5e5e5] bg-[#f9f9f9]">
                                        {#each ['Python', 'TypeScript / Node'] as tab, i}
                                                <button
                                                        on:click={() => activeTab1 = i}
                                                        class="px-4 py-2 text-xs font-medium transition-colors {activeTab1 === i ? 'text-[#0f0f0f] font-semibold bg-white' : 'text-[#888] hover:text-[#444]'}"
                                                >{tab}</button>
                                        {/each}
                                </div>
                                {#if activeTab1 === 0}
                                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>pip install openai</code></pre>
                                {:else}
                                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>npm install openai</code></pre>
                                {/if}
                        </div>
                </div>
        </div>

        <!-- Step 2 -->
        <div class="flex items-start gap-4 mb-8">
                <div class="flex-shrink-0 w-7 h-7 rounded-full bg-[#0f0f0f] text-white text-xs font-bold flex items-center justify-center mt-0.5">2</div>
                <div class="flex-1">
                        <h2 class="text-base font-semibold text-[#0f0f0f] mb-1">Send your first request</h2>
                        <p class="text-sm text-[#666] mb-3">Copy your API key from the Console and paste it below.</p>

                        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden">
                                <div class="flex border-b border-[#e5e5e5] bg-[#f9f9f9]">
                                        {#each ['Python', 'TypeScript', 'cURL'] as tab, i}
                                                <button
                                                        on:click={() => activeTab2 = i}
                                                        class="px-4 py-2 text-xs font-medium transition-colors {activeTab2 === i ? 'text-[#0f0f0f] font-semibold bg-white' : 'text-[#888] hover:text-[#444]'}"
                                                >{tab}</button>
                                        {/each}
                                </div>
                                {#if activeTab2 === 0}
                                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.quickstartPython}</code></pre>
                                {:else if activeTab2 === 1}
                                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.quickstartJS}</code></pre>
                                {:else}
                                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.quickstartCurl}</code></pre>
                                {/if}
                        </div>
                </div>
        </div>

        <!-- Step 3 -->
        <div class="flex items-start gap-4 mb-10">
                <div class="flex-shrink-0 w-7 h-7 rounded-full bg-[#0f0f0f] text-white text-xs font-bold flex items-center justify-center mt-0.5">3</div>
                <div class="flex-1">
                        <h2 class="text-base font-semibold text-[#0f0f0f] mb-1">Check the response</h2>
                        <p class="text-sm text-[#666] mb-3">A successful response looks like this:</p>
                        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden">
                                <div class="flex items-center px-4 py-2.5 bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                        <span class="text-xs font-medium text-[#888888]">JSON</span>
                                </div>
                                <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{`{
  "id": "chatcmpl-abc123",
  "object": "chat.completion",
  "model": "Mistral-Small-4",
  "choices": [{
    "index": 0,
    "message": {
      "role": "assistant",
      "content": "Hello! I'm an AI assistant..."
    },
    "finish_reason": "stop"
  }],
  "usage": { "prompt_tokens": 12, "completion_tokens": 38, "total_tokens": 50 }
}`}</code></pre>
                        </div>
                </div>
        </div>

        <!-- Tip -->
        <div class="flex items-start gap-3 p-4 rounded-xl border border-[#e5e5e5] bg-[#f4f4f4]">
                <svg class="w-4 h-4 text-[#888] flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
                <div class="text-sm text-[#444]">
                        <strong>Migrating from OpenAI?</strong> You only need to change <code class="bg-[#e5e5e5] px-1 rounded font-mono">base_url</code> and <code class="bg-[#e5e5e5] px-1 rounded font-mono">api_key</code>. All existing code works as-is.
                </div>
        </div>

</div>
