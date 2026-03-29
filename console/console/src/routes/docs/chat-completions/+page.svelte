<script lang="ts">
        import { baseUrl, getSnippets } from '$lib/docs-store';
        import CodeBlock from '$lib/CodeBlock.svelte';

        $: s = getSnippets($baseUrl);
        let activeTab = 0;
</script>

<svelte:head><title>Chat Completions — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Chat Completions</h1>
        <p class="text-sm text-[#666] mb-2 leading-relaxed">
                The primary endpoint. Fully compatible with the OpenAI Chat Completions API — change only the <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">base_url</code>.
        </p>

        <div class="flex items-center gap-2 mb-8">
                <span class="text-xs font-bold font-mono text-white bg-green-500 px-2.5 py-1 rounded-lg">POST</span>
                <code class="text-sm text-[#444] font-mono">/v1/chat/completions</code>
        </div>

        <!-- Tabbed code examples -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Code examples</h2>
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
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.python}</code></pre>
                {:else if activeTab === 1}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.js}</code></pre>
                {:else}
                        <pre class="p-4 text-sm font-mono text-[#0f0f0f] overflow-x-auto"><code>{@html s.curlChat}</code></pre>
                {/if}
        </div>

        <!-- Response format -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Response format</h2>
        <p class="text-sm text-[#666] mb-1">Successful non-streaming response:</p>
        <CodeBlock code={s.responseExample} id="chat-resp" lang="json" />

        <!-- Response fields table -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Response fields</h2>
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <table class="w-full text-sm">
                        <thead>
                                <tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Field</th>
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider hidden sm:table-cell">Description</th>
                                </tr>
                        </thead>
                        <tbody class="divide-y divide-[#f0f0f0]">
                                {#each [
                                        { field: 'id',                         desc: 'Unique identifier for the completion' },
                                        { field: 'object',                     desc: '"chat.completion" for non-streaming, "chat.completion.chunk" for streaming' },
                                        { field: 'model',                      desc: 'The model alias that was used (as sent in the request)' },
                                        { field: 'choices[].message.content',  desc: 'The generated text response' },
                                        { field: 'choices[].finish_reason',    desc: '"stop" | "tool_calls" | "length" | "content_filter"' },
                                        { field: 'usage.prompt_tokens',        desc: 'Number of tokens in the input messages' },
                                        { field: 'usage.completion_tokens',    desc: 'Number of tokens in the generated response' },
                                        { field: 'usage.total_tokens',         desc: 'Total tokens used in this request' },
                                ] as row}
                                        <tr class="hover:bg-[#fafafa] transition-colors">
                                                <td class="px-4 py-3"><code class="font-mono text-[12px] text-[#0f0f0f]">{row.field}</code></td>
                                                <td class="px-4 py-3 text-[13px] text-[#666] hidden sm:table-cell">{row.desc}</td>
                                        </tr>
                                {/each}
                        </tbody>
                </table>
        </div>

        <!-- Message roles -->
        <h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Message roles</h2>
        <div class="space-y-2 mb-8">
                {#each [
                        { role: 'system',    color: 'bg-[#f4f4f4] border-[#e5e5e5] text-[#444]', desc: 'Sets the behavior and persona of the assistant. Appears at the start of the conversation.' },
                        { role: 'user',      color: 'bg-[#f4f4f4] border-[#e5e5e5] text-[#444]', desc: 'The human\'s turn in the conversation. Can contain questions, instructions, or context.' },
                        { role: 'assistant', color: 'bg-[#f4f4f4] border-[#e5e5e5] text-[#444]', desc: 'The model\'s previous responses. Include for multi-turn conversations.' },
                        { role: 'tool',      color: 'bg-[#f4f4f4] border-[#e5e5e5] text-[#444]', desc: 'Tool execution results. Sent back to the model after it calls a tool.' },
                ] as m}
                        <div class="flex items-start gap-3 p-3 rounded-xl border {m.color}">
                                <code class="text-xs font-mono font-bold w-16 flex-shrink-0 mt-0.5 text-[#333]">{m.role}</code>
                                <p class="text-sm text-[#555]">{m.desc}</p>
                        </div>
                {/each}
        </div>

        <!-- Note -->
        <div class="flex items-start gap-3 p-4 rounded-xl border border-[#e5e5e5] bg-[#f4f4f4]">
                <svg class="w-4 h-4 text-[#888] flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
                <p class="text-sm text-[#444]">
                        By default, responses are streamed via Server-Sent Events. Set <code class="bg-[#e5e5e5] px-1 rounded font-mono">stream: false</code> to receive a single JSON response. See <a href="/docs/streaming" class="underline font-medium">Streaming</a> for details.
                </p>
        </div>

</div>
