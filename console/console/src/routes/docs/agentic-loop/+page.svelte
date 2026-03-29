<script lang="ts">
	import { baseUrl, getSnippets } from '$lib/docs-store';
	import CodeBlock from '$lib/CodeBlock.svelte';

	$: s = getSnippets($baseUrl);
</script>

<svelte:head><title>Agentic Loop — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Agentic Loop</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		A gateway-native feature: add <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">x-endpoint</code> to any tool definition and the gateway will automatically execute the tool, feed results back to the model, and repeat — up to 10 iterations — until a final answer is produced. No client-side loop needed.
	</p>

	<!-- Comparison -->
	<div class="grid grid-cols-2 gap-3 mb-8">
		<div class="border border-[#e5e5e5] rounded-2xl p-4 bg-[#fafafa]">
			<p class="text-xs font-semibold text-[#888] uppercase tracking-wide mb-2">Standard tool calls</p>
			<p class="text-sm text-[#555] leading-snug">Your code executes each tool and sends results back manually.</p>
			<p class="text-xs text-[#aaa] mt-2">Many round-trips required</p>
		</div>
		<div class="border border-[#e5e5e5] rounded-2xl p-4 bg-[#f4f4f4]">
			<p class="text-xs font-semibold text-[#0f0f0f] uppercase tracking-wide mb-2">Agentic Loop ✦</p>
			<p class="text-sm text-[#555] leading-snug">The gateway executes tools automatically via x-endpoint URLs.</p>
			<p class="text-xs text-[#888] mt-2">Single request → final answer</p>
		</div>
	</div>

	<!-- How it works steps -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">How it works</h2>
	<div class="border border-[#e5e5e5] rounded-2xl overflow-hidden mb-8">
		{#each [
			{ n: 1, title: 'Client sends request', desc: 'You send a request with tools that include x-endpoint URLs. A single HTTP call.' },
			{ n: 2, title: 'Gateway calls the model', desc: 'Gateway makes a non-streaming call to the LLM to get tool call decisions.' },
			{ n: 3, title: 'Gateway executes tools', desc: 'For each tool_call, the gateway POSTs to the x-endpoint with the model\'s JSON arguments.' },
			{ n: 4, title: 'Results fed back', desc: 'Tool results are appended to the conversation as "tool" messages and sent to the model again.' },
			{ n: 5, title: 'Repeat until done', desc: 'Steps 2–4 repeat up to 10 times. If the model returns no tool calls, the final response is sent to you.' },
		] as step, i}
			<div class="flex items-start gap-4 p-4 {i < 4 ? 'border-b border-[#f0f0f0]' : ''} hover:bg-[#fafafa] transition-colors">
				<span class="flex-shrink-0 w-6 h-6 rounded-full bg-[#0f0f0f] text-white text-[11px] font-bold flex items-center justify-center">{step.n}</span>
				<div>
					<p class="text-sm font-semibold text-[#0f0f0f]">{step.title}</p>
					<p class="text-sm text-[#666] mt-0.5">{step.desc}</p>
				</div>
			</div>
		{/each}
	</div>

	<!-- Example request -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Example — automatic web search</h2>
	<p class="text-sm text-[#666] mb-1">Include <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">x-endpoint</code> in the tool's function object:</p>
	<CodeBlock code={s.agenticCurl} id="ag-curl" label="POST /v1/chat/completions" method="POST" lang="bash" />

	<!-- What your endpoint receives -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">Your endpoint contract</h2>
	<p class="text-sm text-[#666] mb-1">The gateway POSTs the tool arguments as JSON. Your endpoint should return any JSON:</p>
	<CodeBlock code={s.agenticEndpoint} id="ag-ep" lang="json" />

	<!-- x-endpoint spec -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-3">x-endpoint specification</h2>
	<div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
		<div class="divide-y divide-[#f0f0f0]">
			{#each [
				{ field: 'Location',       value: 'Inside function.x-endpoint (sibling of name, description, parameters)' },
				{ field: 'Type',           value: 'string — a valid HTTP/HTTPS URL' },
				{ field: 'Request method', value: 'POST' },
				{ field: 'Request body',   value: 'JSON of the model\'s tool call arguments' },
				{ field: 'Response',       value: 'Any JSON — forwarded as-is to the model as a tool message' },
				{ field: 'Timeout',        value: '30 seconds per tool call' },
				{ field: 'Max iterations', value: '10 (gateway retries without tools after limit)' },
			] as row}
				<div class="px-5 py-3 flex flex-col sm:flex-row sm:items-center gap-1 sm:gap-4">
					<p class="text-xs font-semibold text-[#888] w-36 flex-shrink-0">{row.field}</p>
					<p class="text-sm text-[#444]">{row.value}</p>
				</div>
			{/each}
		</div>
	</div>

	<!-- Warnings -->
	<div class="space-y-3">
		<div class="flex items-start gap-3 p-4 rounded-xl border border-amber-100 bg-amber-50">
			<svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
			<p class="text-sm text-amber-800">
				The agentic loop always returns a <strong>non-streaming JSON response</strong>, regardless of the <code class="bg-amber-100 px-1 rounded font-mono">stream</code> setting. Set <code class="bg-amber-100 px-1 rounded font-mono">stream: false</code> explicitly.
			</p>
		</div>
		<div class="flex items-start gap-3 p-4 rounded-xl border border-[#e5e5e5] bg-[#f4f4f4]">
			<svg class="w-4 h-4 text-[#888] flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
			<p class="text-sm text-[#444]">
				You can mix agentic tools (with <code class="bg-[#e5e5e5] px-1 rounded font-mono">x-endpoint</code>) and regular tools (without) in the same request. Only tools with <code class="bg-[#e5e5e5] px-1 rounded font-mono">x-endpoint</code> are auto-executed.
			</p>
		</div>
	</div>

</div>
