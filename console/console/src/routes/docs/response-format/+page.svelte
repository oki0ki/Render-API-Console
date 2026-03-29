<script lang="ts">
	import { baseUrl, getSnippets } from '$lib/docs-store';
	import CodeBlock from '$lib/CodeBlock.svelte';

	$: s = getSnippets($baseUrl);
</script>

<svelte:head><title>Response Format — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Response Format</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		All successful responses from <code class="font-mono bg-[#f4f4f4] px-1 rounded">POST /v1/chat/completions</code> follow the
		OpenAI Chat Completion object schema. This page describes every field you may encounter.
	</p>

	<!-- Full example -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Example response</h2>
	<CodeBlock code={s.responseExample} id="resp-example" lang="json" />

	<!-- Top-level fields -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mt-8 mb-4">Top-level fields</h2>
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
					{ field: 'id',      type: 'string',  desc: 'Unique identifier for this completion (chatcmpl-…).' },
					{ field: 'object',  type: 'string',  desc: 'Always "chat.completion" for non-streaming responses.' },
					{ field: 'created', type: 'integer', desc: 'Unix timestamp (seconds) when the completion was generated.' },
					{ field: 'model',   type: 'string',  desc: 'The model ID that produced the response.' },
					{ field: 'choices', type: 'array',   desc: 'One or more candidate completions. Contains the message.' },
					{ field: 'usage',   type: 'object',  desc: 'Token consumption breakdown for billing and monitoring.' },
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

	<!-- choices[].message -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">choices[ ] object</h2>
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
					{ field: 'index',         type: 'integer', desc: 'Zero-based index of this choice (always 0 — only one choice is returned).' },
					{ field: 'message',       type: 'object',  desc: 'The assistant message. Contains role and content (and tool_calls when applicable).' },
					{ field: 'message.role',  type: 'string',  desc: 'Always "assistant".' },
					{ field: 'message.content', type: 'string | null', desc: 'The text reply from the model. null when finish_reason is "tool_calls".' },
					{ field: 'message.tool_calls', type: 'array | null', desc: 'Present when the model wants to invoke one or more tools.' },
					{ field: 'finish_reason', type: 'string',  desc: '"stop" | "tool_calls" | "length" | "content_filter"' },
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

	<!-- finish_reason values -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">finish_reason values</h2>
	<div class="space-y-2 mb-8">
		{#each [
			{ value: 'stop',           color: 'bg-green-100 text-green-700',  desc: 'Normal completion — the model finished naturally.' },
			{ value: 'tool_calls',     color: 'bg-violet-100 text-violet-700', desc: 'The model wants to call one or more tools. Handle tool_calls and continue the conversation.' },
			{ value: 'length',         color: 'bg-amber-100 text-amber-700',  desc: 'The response was truncated because it reached max_tokens.' },
			{ value: 'content_filter', color: 'bg-red-100 text-red-700',      desc: 'The output was blocked by a content safety filter.' },
		] as fr}
			<div class="flex items-start gap-3 p-3 rounded-xl border border-[#e5e5e5] bg-[#fafafa]">
				<code class="flex-shrink-0 text-xs font-bold px-2 py-0.5 rounded mt-0.5 {fr.color}">{fr.value}</code>
				<p class="text-sm text-[#555]">{fr.desc}</p>
			</div>
		{/each}
	</div>

	<!-- usage -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">usage object</h2>
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
					{ field: 'prompt_tokens',     type: 'integer', desc: 'Tokens in the input messages (system + user + any prior assistant turns).' },
					{ field: 'completion_tokens', type: 'integer', desc: 'Tokens generated in the assistant reply.' },
					{ field: 'total_tokens',      type: 'integer', desc: 'Sum of prompt_tokens and completion_tokens.' },
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

	<!-- Streaming note -->
	<div class="bg-[#fffbeb] border border-[#fde68a] rounded-2xl p-5">
		<p class="text-sm font-semibold text-[#92400e] mb-1">Streaming responses differ</p>
		<p class="text-sm text-[#78350f]">
			When <code class="font-mono bg-[#fef3c7] px-1 rounded">stream: true</code>, the API returns a series of
			<code class="font-mono bg-[#fef3c7] px-1 rounded">chat.completion.chunk</code> objects via SSE.
			Each chunk has a <code class="font-mono bg-[#fef3c7] px-1 rounded">delta</code> field instead of <code class="font-mono bg-[#fef3c7] px-1 rounded">message</code>,
			and <code class="font-mono bg-[#fef3c7] px-1 rounded">usage</code> is only present in the final chunk.
			See the <a href="/docs/streaming" class="underline font-semibold">Streaming</a> page for details.
		</p>
	</div>

</div>
