<svelte:head><title>Rate Limits — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Rate Limits</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		Rate limits protect the gateway from overload and ensure fair access across all API consumers.
		Limits are enforced per API key.
	</p>

	<!-- Limits table -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">Current limits</h2>
	<div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
		<table class="w-full text-sm">
			<thead>
				<tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Limit type</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Value</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Scope</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-[#f0f0f0]">
				{#each [
					{ type: 'Requests per minute (RPM)',  value: '60',    scope: 'Per API key' },
					{ type: 'Requests per day (RPD)',     value: '5 000', scope: 'Per API key' },
					{ type: 'Tokens per minute (TPM)',    value: '200 000', scope: 'Per API key' },
					{ type: 'Max tokens per request',    value: '128 000', scope: 'Context window (model-dependent)' },
					{ type: 'Concurrent requests',       value: '10',    scope: 'Per API key' },
					{ type: 'Tool call iterations',      value: '10',    scope: 'Per agentic request' },
					{ type: 'Upstream timeout',          value: '120 s', scope: 'Per request' },
				] as row}
					<tr class="hover:bg-[#fafafa] transition-colors">
						<td class="px-4 py-3 text-[13px] text-[#333]">{row.type}</td>
						<td class="px-4 py-3 font-mono text-[13px] font-semibold text-[#0f0f0f]">{row.value}</td>
						<td class="px-4 py-3 text-[13px] text-[#666]">{row.scope}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<!-- 429 response -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">When you exceed a limit — 429</h2>
	<div class="flex items-start gap-4 p-4 rounded-2xl border border-red-200 bg-red-50 mb-6">
		<span class="font-mono text-sm font-bold px-2.5 py-1 rounded-lg bg-red-100 text-red-700 flex-shrink-0 leading-none">429</span>
		<div>
			<p class="text-sm font-semibold text-[#1a1a1a] mb-0.5">Too Many Requests</p>
			<p class="text-sm text-[#555]">
				The response body will contain a plain-text message indicating which limit was hit (RPM, RPD, or TPM).
				Wait for the window to reset before retrying.
			</p>
		</div>
	</div>

	<!-- Response headers -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-4">Rate-limit response headers</h2>
	<p class="text-sm text-[#666] mb-4">Every API response includes these headers so you can track your current usage programmatically:</p>
	<div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
		<table class="w-full text-sm">
			<thead>
				<tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Header</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Description</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-[#f0f0f0]">
				{#each [
					{ header: 'x-ratelimit-limit-requests',     desc: 'Maximum requests allowed per minute.' },
					{ header: 'x-ratelimit-remaining-requests', desc: 'Requests remaining in the current window.' },
					{ header: 'x-ratelimit-reset-requests',     desc: 'ISO 8601 timestamp when the request window resets.' },
					{ header: 'x-ratelimit-limit-tokens',       desc: 'Maximum tokens allowed per minute.' },
					{ header: 'x-ratelimit-remaining-tokens',   desc: 'Tokens remaining in the current window.' },
					{ header: 'x-ratelimit-reset-tokens',       desc: 'ISO 8601 timestamp when the token window resets.' },
				] as row}
					<tr class="hover:bg-[#fafafa] transition-colors">
						<td class="px-4 py-3 font-mono text-[12px] text-[#333]">{row.header}</td>
						<td class="px-4 py-3 text-[13px] text-[#666]">{row.desc}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<!-- Best practices -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">Handling rate limits in production</h2>
	<div class="space-y-2 mb-8">
		{#each [
			{ icon: '✓', text: 'Read x-ratelimit-remaining-requests before each request to check available capacity.' },
			{ icon: '✓', text: 'On a 429 response, parse x-ratelimit-reset-requests and sleep until that timestamp.' },
			{ icon: '✓', text: 'Use exponential backoff with jitter to avoid thundering-herd retries: 1s → 2s → 4s + random(0–500ms).' },
			{ icon: '✓', text: 'Batch multiple short questions into a single request to reduce RPM consumption.' },
			{ icon: '✓', text: 'Cache responses for identical prompts to avoid redundant API calls.' },
		] as tip}
			<div class="flex items-start gap-3 p-3 rounded-xl bg-[#fafafa] border border-[#f0f0f0]">
				<span class="text-green-500 font-bold text-sm flex-shrink-0">{tip.icon}</span>
				<p class="text-sm text-[#555]">{tip.text}</p>
			</div>
		{/each}
	</div>

	<!-- Info box -->
	<div class="bg-[#f4f4f4] border border-[#e5e5e5] rounded-2xl p-5">
		<p class="text-sm font-semibold text-[#0f0f0f] mb-1">Need higher limits?</p>
		<p class="text-sm text-[#555]">
			Contact us through the Console to request increased rate limits for production workloads.
			Include your use case and estimated daily request volume.
		</p>
	</div>

</div>
