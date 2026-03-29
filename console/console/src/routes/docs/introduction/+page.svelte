<script lang="ts">
        import { baseUrl, copyText, copiedId } from '$lib/docs-store';
        import CodeBlock from '$lib/CodeBlock.svelte';

        const featureCards = [
                {
                        icon: '<svg width="24" height="24" viewBox="0 0 20 20" fill="currentColor"><path d="M4.12167 4.16233C4.14825 3.92766 4.34665 3.75028 4.58282 3.75004C4.81898 3.7498 5.01775 3.92678 5.04481 4.16139C5.21375 5.62616 6.0405 6.45291 7.50528 6.62186C7.73989 6.64892 7.91687 6.84768 7.91663 7.08385C7.91639 7.32001 7.739 7.51842 7.50434 7.545C6.06034 7.70856 5.17626 8.52734 5.04605 9.99324C5.02474 10.2331 4.82362 10.417 4.58277 10.4167C4.34192 10.4164 4.14122 10.2321 4.12045 9.9922C3.99535 8.54703 3.11964 7.67132 1.67447 7.54622C1.43452 7.52545 1.25023 7.32475 1.24996 7.0839C1.24969 6.84305 1.43353 6.64193 1.67343 6.62062C3.13933 6.49041 3.95811 5.60633 4.12167 4.16233Z"/><path fill-rule="evenodd" clip-rule="evenodd" d="M13.1914 5.41671C12.8514 5.41671 12.5549 5.64816 12.4724 5.97809C12.1151 7.40728 11.6948 8.2739 11.0925 8.87625C10.4901 9.47859 9.62351 9.89892 8.19432 10.2562C7.86439 10.3387 7.63294 10.6351 7.63294 10.9752C7.63294 11.3153 7.86439 11.6117 8.19432 11.6942C9.62351 12.0515 10.4901 12.4718 11.0925 13.0742C11.6948 13.6765 12.1151 14.5432 12.4724 15.9723C12.5549 16.3023 12.8514 16.5337 13.1914 16.5337C13.5315 16.5337 13.828 16.3023 13.9105 15.9723C14.2678 14.5432 14.6881 13.6765 15.2904 13.0742C15.8928 12.4718 16.7594 12.0515 18.1886 11.6942C18.5185 11.6117 18.75 11.3153 18.75 10.9752C18.75 10.6351 18.5185 10.3387 18.1886 10.2562C16.7594 9.89892 15.8928 9.47859 15.2904 8.87625C14.6881 8.2739 14.2678 7.40728 13.9105 5.97809C13.828 5.64816 13.5315 5.41671 13.1914 5.41671Z"/></svg>',
                        label: 'OpenAI SDK',
                        desc: 'Drop-in replacement — change only base_url',
                },
                {
                        icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path fill-rule="evenodd" d="M12.5 3.444a1 1 0 0 0-1 0l-6.253 3.61 6.768 3.807 6.955-3.682-6.47-3.735Zm7.16 5.632L13 12.602v7.666l6.16-3.556a1 1 0 0 0 .5-.867V9.076ZM11 20.268v-7.683L4.34 8.839v7.006a1 1 0 0 0 .5.867L11 20.268Zm-.5-18.557a3 3 0 0 1 3 0l6.66 3.846a3 3 0 0 1 1.5 2.598v7.69a3 3 0 0 1-1.5 2.598L13.5 22.29a3 3 0 0 1-3 0l-6.66-3.846a3 3 0 0 1-1.5-2.598v-7.69a3 3 0 0 1 1.5-2.598L10.5 1.71Z" clip-rule="evenodd"></path></svg>',
                        label: '11 Models',
                        desc: 'NVIDIA NIM + Amazon Nova providers',
                },
                {
                        icon: '<svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.5574 7.40381C14.0958 4.8654 18.2113 4.8654 20.7497 7.40381C23.2882 9.94221 23.2882 14.0578 20.7497 16.5962C18.2113 19.1346 14.0958 19.1346 11.5574 16.5962L8.37538 13.4142L7.31472 14.4749C5.94788 15.8417 3.7318 15.8417 2.36497 14.4749C0.998135 13.108 0.998135 10.892 2.36497 9.52513C3.7318 8.15829 5.94788 8.15829 7.31472 9.52513L8.37538 10.5858L11.5574 7.40381ZM9.78959 12L12.9716 15.182C14.7289 16.9393 17.5782 16.9393 19.3355 15.182C21.0929 13.4246 21.0929 10.5754 19.3355 8.81802C17.5782 7.06066 14.7289 7.06066 12.9716 8.81802L9.78959 12ZM6.96116 12L5.9005 10.9393C5.31472 10.3536 4.36497 10.3536 3.77918 10.9393C3.1934 11.5251 3.1934 12.4749 3.77918 13.0607C4.36497 13.6464 5.31472 13.6464 5.9005 13.0607L6.96116 12Z"/></svg>',
                        label: 'Agentic Loop',
                        desc: 'Auto tool execution — up to 10 iterations',
                },
        ];
</script>

<svelte:head><title>Introduction — API Docs</title></svelte:head>

<div class="max-w-[760px]">

        <div class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-[#f4f4f4] text-xs text-[#444] font-semibold mb-6 border border-[#e5e5e5]">
                <span class="w-1.5 h-1.5 rounded-full bg-[#444]"></span>
                OpenAI-compatible
        </div>

        <h1 class="text-[2rem] font-bold tracking-tight text-[#0f0f0f] mb-4 leading-tight">AI Gateway API</h1>
        <p class="text-[16px] text-[#555] leading-relaxed mb-8 max-w-[620px]">
                A unified gateway giving you access to 11 state-of-the-art language models through a single, OpenAI-compatible endpoint. Drop it into any existing OpenAI SDK integration — no code changes beyond the base URL.
        </p>

        <!-- Feature cards -->
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-3 mb-10">
                {#each featureCards as f}
                        <div class="border border-[#f0f0f0] rounded-2xl p-5 bg-[#fafafa] hover:border-[#e5e5e5] hover:bg-[#f4f4f4]/50 transition-colors">
                                <div class="text-[#0f0f0f] mb-3">{@html f.icon}</div>
                                <p class="text-sm font-semibold text-[#111]">{f.label}</p>
                                <p class="text-xs text-[#888] mt-1 leading-snug">{f.desc}</p>
                        </div>
                {/each}
        </div>

        <!-- Base URL -->
        <h2 class="text-lg font-semibold text-[#0f0f0f] mb-3">Base URL</h2>
        <p class="text-sm text-[#666] mb-3">All API requests are made to this base URL. Append the endpoint path.</p>
        <button
                on:click={() => copyText('intro-base-url', $baseUrl)}
                class="w-full text-left bg-[#f9f9f9] border border-[#e5e5e5] rounded-2xl p-5 hover:bg-[#f4f4f4] transition-colors group mb-8"
        >
                <div class="flex items-center justify-between mb-2">
                        <p class="text-[10px] font-semibold text-[#aaa] uppercase tracking-wider">Base URL</p>
                        <span class="text-[10px] text-[#bbb] group-hover:text-[#666] transition-colors">
                                {$copiedId === 'intro-base-url' ? '✓ Copied' : 'Click to copy'}
                        </span>
                </div>
                <code class="text-[15px] font-mono text-[#0f0f0f]">{$baseUrl || 'Loading…'}</code>
        </button>

        <!-- Endpoints overview -->
        <h2 class="text-lg font-semibold text-[#0f0f0f] mb-3">Endpoints</h2>
        <div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
                <table class="w-full text-sm">
                        <thead>
                                <tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Method</th>
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Path</th>
                                        <th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider hidden sm:table-cell">Description</th>
                                </tr>
                        </thead>
                        <tbody class="divide-y divide-[#f0f0f0]">
                                {#each [
                                        { method: 'GET',  path: '/v1/models',           desc: 'List all available models' },
                                        { method: 'POST', path: '/v1/chat/completions', desc: 'Create a chat completion' },
                                        { method: 'GET',  path: '/',                    desc: 'Detect the gateway base URL' },
                                ] as e}
                                        <tr class="hover:bg-[#fafafa] transition-colors">
                                                <td class="px-4 py-3">
                                                        <span class="text-xs font-bold font-mono px-2 py-0.5 rounded
                                                                {e.method === 'POST' ? 'bg-green-100 text-green-700' : 'bg-[#f4f4f4] text-[#555]'}">
                                                                {e.method}
                                                        </span>
                                                </td>
                                                <td class="px-4 py-3 font-mono text-[13px] text-[#333]">{e.path}</td>
                                                <td class="px-4 py-3 text-[13px] text-[#666] hidden sm:table-cell">{e.desc}</td>
                                        </tr>
                                {/each}
                        </tbody>
                </table>
        </div>

        <!-- Auth summary -->
        <h2 class="text-lg font-semibold text-[#0f0f0f] mb-3">Authentication</h2>
        <p class="text-sm text-[#666] mb-3">Every request must include an API key in one of these headers:</p>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 mb-8">
                {#each [
                        { name: 'Authorization', value: 'Bearer sk-<your-key>' },
                        { name: 'x-api-key',    value: 'sk-<your-key>' },
                ] as h}
                        <div class="bg-[#f9f9f9] border border-[#e5e5e5] rounded-xl p-4">
                                <p class="text-[10px] font-semibold text-[#aaa] uppercase tracking-wide mb-2">{h.name}</p>
                                <code class="text-xs font-mono text-[#333]">{h.value}</code>
                        </div>
                {/each}
        </div>

        <!-- Next steps -->
        <div class="bg-[#f4f4f4] border border-[#e5e5e5] rounded-2xl p-5">
                <p class="text-sm font-semibold text-[#0f0f0f] mb-2">Ready to get started?</p>
                <p class="text-sm text-[#555] mb-3">Head to the Quick Start guide to make your first API call in under a minute.</p>
                <a href="/docs/quickstart" class="inline-flex items-center gap-1.5 text-sm font-semibold text-[#0f0f0f] hover:text-[#444] transition-colors">
                        Quick Start
                        <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M12 5l7 7-7 7"/></svg>
                </a>
        </div>

</div>
