<script lang="ts">
	import { baseUrl, getSnippets } from '$lib/docs-store';
	import CodeBlock from '$lib/CodeBlock.svelte';

	$: s = getSnippets($baseUrl);
</script>

<svelte:head><title>Tool Calls — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Tool Calls</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		Pass <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">tools</code> in the standard OpenAI format. The model decides when to call a tool and returns structured arguments. Your application is responsible for executing the tool and sending the result back.
	</p>

	<!-- Flow diagram -->
	<div class="border border-[#e5e5e5] rounded-2xl p-5 mb-8 bg-[#fafafa]">
		<p class="text-xs font-semibold text-[#888] uppercase tracking-wide mb-4">Standard tool call flow</p>
		<div class="flex flex-col gap-2">
			{#each [
				{ arrow: false, label: 'You', text: 'Send request with tools + user message' },
				{ arrow: true },
				{ label: 'Model', text: 'Returns finish_reason: "tool_calls" with arguments' },
				{ arrow: true },
				{ label: 'You', text: 'Execute the tool in your application' },
				{ arrow: true },
				{ label: 'You', text: 'Send tool result back with role: "tool"' },
				{ arrow: true },
				{ label: 'Model', text: 'Returns final answer using the tool result' },
			] as step}
				{#if step.arrow}
					<div class="flex items-center justify-center text-[#ccc]">↓</div>
				{:else}
					<div class="flex items-center gap-3 p-3 rounded-xl bg-white border border-[#eee]">
						<span class="text-xs font-semibold w-12 flex-shrink-0
							{step.label === 'Model' ? 'text-[#0f0f0f]' : 'text-[#888]'}">
							{step.label}
						</span>
						<p class="text-sm text-[#444]">{step.text}</p>
					</div>
				{/if}
			{/each}
		</div>
		<p class="text-xs text-[#aaa] mt-3 text-center">
			For automatic tool execution, see <a href="/docs/agentic-loop" class="text-[#0f0f0f] hover:underline font-medium">Agentic Loop →</a>
		</p>
	</div>

	<!-- Tool definition -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Tool definition</h2>
	<p class="text-sm text-[#666] mb-1">Define tools using JSON Schema for parameters:</p>
	<CodeBlock code={s.toolCurl} id="tool-curl" label="POST /v1/chat/completions" method="POST" lang="bash" />

	<!-- Response when tool is called -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Model response when calling a tool</h2>
	<p class="text-sm text-[#666] mb-1">When the model decides to call a tool, <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">finish_reason</code> is <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">"tool_calls"</code>:</p>
	<CodeBlock code={s.toolResponse} id="tool-resp" lang="json" />

	<!-- Continue with tool result -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Continue with tool result</h2>
	<p class="text-sm text-[#666] mb-1">Execute the tool in your code, then send the result back:</p>
	<CodeBlock code={s.toolFollow} id="tool-follow" lang="python" />

	<!-- Tool definition structure -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Tool object structure</h2>
	<div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-4">
		<div class="divide-y divide-[#f0f0f0]">
			{#each [
				{ field: 'type',                      req: true,  desc: 'Always "function".' },
				{ field: 'function.name',             req: true,  desc: 'The function name. The model uses this to call the tool.' },
				{ field: 'function.description',      req: false, desc: 'Human-readable description. Helps the model decide when to use this tool.' },
				{ field: 'function.parameters',       req: false, desc: 'JSON Schema object describing the function\'s parameters.' },
				{ field: 'function.x-endpoint',       req: false, desc: 'Custom extension: URL for automatic tool execution (Agentic Loop only).' },
			] as f}
				<div class="px-5 py-3">
					<div class="flex items-center gap-2 mb-0.5">
						<code class="text-[12px] font-mono font-semibold text-[#0f0f0f]">{f.field}</code>
						{#if f.req}
							<span class="text-[10px] font-bold text-red-500 bg-red-50 border border-red-100 px-1.5 py-0.5 rounded-full uppercase">required</span>
						{:else}
							<span class="text-[10px] text-[#aaa]">optional</span>
						{/if}
					</div>
					<p class="text-sm text-[#555]">{f.desc}</p>
				</div>
			{/each}
		</div>
	</div>

	<div class="flex items-start gap-3 p-4 rounded-xl border border-amber-100 bg-amber-50">
		<svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
		<p class="text-sm text-amber-800">
			Want the gateway to execute tools automatically without any client-side loop? Add <code class="bg-amber-100 px-1 rounded font-mono">x-endpoint</code> to the tool and use the <a href="/docs/agentic-loop" class="underline font-medium">Agentic Loop</a>.
		</p>
	</div>

</div>
