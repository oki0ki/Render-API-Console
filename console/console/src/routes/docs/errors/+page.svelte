<script lang="ts">
	import { baseUrl, getSnippets } from '$lib/docs-store';
	import CodeBlock from '$lib/CodeBlock.svelte';

	$: s = getSnippets($baseUrl);

	const errors = [
		{
			code: '200', label: 'OK',
			bg: 'bg-green-50', border: 'border-green-200', badge: 'bg-green-100 text-green-700',
			desc: 'Request succeeded. The response body contains the completion or event stream.',
		},
		{
			code: '400', label: 'Bad Request',
			bg: 'bg-amber-50', border: 'border-amber-200', badge: 'bg-amber-100 text-amber-700',
			desc: 'The request body is malformed or missing required fields. Check your JSON syntax and required parameters (model, messages).',
		},
		{
			code: '401', label: 'Unauthorized',
			bg: 'bg-red-50', border: 'border-red-200', badge: 'bg-red-100 text-red-700',
			desc: 'The API key is missing, invalid, or has been deactivated. Check your Authorization header.',
		},
		{
			code: '502', label: 'Bad Gateway',
			bg: 'bg-orange-50', border: 'border-orange-200', badge: 'bg-orange-100 text-orange-700',
			desc: 'The upstream model provider returned an error or is temporarily unavailable. Retry with exponential backoff.',
		},
	];
</script>

<svelte:head><title>Error Handling — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Error Handling</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		The API uses standard HTTP status codes to indicate success or failure. Error response bodies are plain text.
	</p>

	<!-- Status codes -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">HTTP status codes</h2>
	<div class="space-y-3 mb-8">
		{#each errors as e}
			<div class="flex items-start gap-4 p-4 rounded-2xl border {e.border} {e.bg}">
				<span class="font-mono text-sm font-bold px-2.5 py-1 rounded-lg {e.badge} flex-shrink-0 leading-none">{e.code}</span>
				<div>
					<p class="text-sm font-semibold text-[#1a1a1a] mb-0.5">{e.label}</p>
					<p class="text-sm text-[#555]">{e.desc}</p>
				</div>
			</div>
		{/each}
	</div>

	<!-- Handling errors -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Handling specific errors</h2>

	<div class="space-y-5 mb-8">

		<!-- 401 handling -->
		<div class="border border-[#e5e5e5] rounded-2xl overflow-hidden">
			<div class="flex items-center gap-2 px-4 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
				<span class="font-mono text-xs font-bold px-2 py-0.5 rounded bg-red-100 text-red-700">401</span>
				<span class="text-sm font-semibold text-[#333]">Unauthorized</span>
			</div>
			<div class="p-4 text-sm text-[#555] space-y-2">
				<p>Check these common causes:</p>
				<ul class="list-disc pl-5 space-y-1 text-sm text-[#666]">
					<li>Missing <code class="font-mono bg-[#f4f4f4] px-1 rounded">Authorization</code> header</li>
					<li>Key is incorrectly formatted (must start with <code class="font-mono bg-[#f4f4f4] px-1 rounded">sk-</code>)</li>
					<li>Key has been deactivated in the Console</li>
					<li>Extra whitespace around the key value</li>
				</ul>
			</div>
		</div>

		<!-- 400 handling -->
		<div class="border border-[#e5e5e5] rounded-2xl overflow-hidden">
			<div class="flex items-center gap-2 px-4 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
				<span class="font-mono text-xs font-bold px-2 py-0.5 rounded bg-amber-100 text-amber-700">400</span>
				<span class="text-sm font-semibold text-[#333]">Bad Request</span>
			</div>
			<div class="p-4 text-sm text-[#555] space-y-2">
				<p>Common causes:</p>
				<ul class="list-disc pl-5 space-y-1 text-sm text-[#666]">
					<li>Missing required fields: <code class="font-mono bg-[#f4f4f4] px-1 rounded">model</code> or <code class="font-mono bg-[#f4f4f4] px-1 rounded">messages</code></li>
					<li>Invalid JSON in the request body</li>
					<li>Wrong Content-Type header (must be <code class="font-mono bg-[#f4f4f4] px-1 rounded">application/json</code>)</li>
				</ul>
			</div>
		</div>

		<!-- 502 handling -->
		<div class="border border-[#e5e5e5] rounded-2xl overflow-hidden">
			<div class="flex items-center gap-2 px-4 py-3 bg-[#f9f9f9] border-b border-[#e5e5e5]">
				<span class="font-mono text-xs font-bold px-2 py-0.5 rounded bg-orange-100 text-orange-700">502</span>
				<span class="text-sm font-semibold text-[#333]">Bad Gateway — Retry strategy</span>
			</div>
			<div class="p-4">
				<p class="text-sm text-[#555] mb-3">502 errors are transient. Implement exponential backoff for production workloads:</p>
				<CodeBlock code={s.retryPython} id="retry-py" lang="python" />
			</div>
		</div>
	</div>

	<!-- Best practices -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Best practices</h2>
	<div class="space-y-2 mb-8">
		{#each [
			{ icon: '✓', text: 'Always check finish_reason in the response (stop | tool_calls | length | content_filter).' },
			{ icon: '✓', text: 'Log the request id from error responses to help with debugging.' },
			{ icon: '✓', text: 'For 502 errors, retry with exponential backoff: 1s → 2s → 4s (max 3 retries).' },
			{ icon: '✓', text: 'For streaming, handle stream disconnections and implement reconnection logic.' },
			{ icon: '✓', text: 'Set a client-side timeout (recommended: 120s) to catch stuck requests.' },
		] as tip}
			<div class="flex items-start gap-3 p-3 rounded-xl bg-[#fafafa] border border-[#f0f0f0]">
				<span class="text-green-500 font-bold text-sm flex-shrink-0">{tip.icon}</span>
				<p class="text-sm text-[#555]">{tip.text}</p>
			</div>
		{/each}
	</div>

	<!-- Footer -->
	<div class="mt-4 pt-6 border-t border-[#f0f0f0] text-center">
		<p class="text-sm text-[#ccc]">AI Gateway API · OpenAI-compatible · v1</p>
		<p class="text-xs text-[#ddd] mt-1">All endpoints · Max 10 tool iterations · 120s upstream timeout</p>
	</div>

</div>
