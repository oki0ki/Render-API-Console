<svelte:head><title>Best Practices — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Best Practices</h1>
	<p class="text-sm text-[#666] mb-10 leading-relaxed">
		Recommendations for building reliable, efficient, and cost-effective applications on top of the AI Gateway API.
	</p>

	<!-- Sections -->
	{#each [
		{
			title: 'Reliability',
			icon: '⟳',
			color: 'text-[#555]',
			bg: 'bg-[#f4f4f4]',
			border: 'border-[#e5e5e5]',
			tips: [
				{ ok: true,  text: 'Always implement exponential backoff for 502 errors: 1s → 2s → 4s, max 3 retries.' },
				{ ok: true,  text: 'Set a client-side timeout of 120 seconds — equal to the upstream limit.' },
				{ ok: true,  text: 'For streaming, handle disconnections and implement SSE reconnection logic.' },
				{ ok: true,  text: 'Check finish_reason in every response: "length" means your output was truncated.' },
				{ ok: false, text: 'Don\'t retry immediately on a 429 — read x-ratelimit-reset-requests and wait.' },
			]
		},
		{
			title: 'Performance',
			icon: '◈',
			color: 'text-[#555]',
			bg: 'bg-[#f4f4f4]',
			border: 'border-[#e5e5e5]',
			tips: [
				{ ok: true,  text: 'Use streaming (stream: true) for long responses to reduce perceived latency.' },
				{ ok: true,  text: 'Trim your system prompt — shorter prompts mean faster first-token time and lower cost.' },
				{ ok: true,  text: 'Cache responses for deterministic prompts (temperature: 0) using a prompt hash as cache key.' },
				{ ok: true,  text: 'Re-use client instances across requests instead of creating a new OpenAI() on each call.' },
				{ ok: false, text: 'Don\'t set max_tokens much larger than you need — it reserves capacity on the upstream model.' },
			]
		},
		{
			title: 'Prompt engineering',
			icon: '✦',
			color: 'text-[#555]',
			bg: 'bg-[#f4f4f4]',
			border: 'border-[#e5e5e5]',
			tips: [
				{ ok: true,  text: 'Put stable instructions in the system message; put dynamic content in the user message.' },
				{ ok: true,  text: 'For JSON output, always describe the exact schema in the system prompt and set response_format.' },
				{ ok: true,  text: 'Use few-shot examples in the messages array for consistent formatting.' },
				{ ok: true,  text: 'Lower temperature (0.1–0.3) for factual / structured tasks; higher (0.7–1.0) for creative tasks.' },
				{ ok: false, text: 'Don\'t include private or sensitive data in prompts — treat every request as potentially logged.' },
			]
		},
		{
			title: 'Security',
			icon: '◉',
			color: 'text-red-600',
			bg: 'bg-red-50',
			border: 'border-red-100',
			tips: [
				{ ok: true,  text: 'Store API keys in environment variables — never hard-code them in source code.' },
				{ ok: true,  text: 'Rotate keys regularly and immediately if you suspect exposure.' },
				{ ok: true,  text: 'Use separate API keys per environment (dev, staging, production) for fine-grained revocation.' },
				{ ok: true,  text: 'Sanitise user-supplied content before inserting it into prompts (prompt injection).' },
				{ ok: false, text: 'Don\'t share API keys across teams or services — each consumer should have its own key.' },
			]
		},
		{
			title: 'Cost management',
			icon: '⊛',
			color: 'text-amber-600',
			bg: 'bg-amber-50',
			border: 'border-amber-100',
			tips: [
				{ ok: true,  text: 'Monitor usage.total_tokens per request and alert when it exceeds expected thresholds.' },
				{ ok: true,  text: 'Choose the smallest model that meets your quality bar — smaller models are faster and cheaper.' },
				{ ok: true,  text: 'Use max_tokens to cap runaway completions in production.' },
				{ ok: true,  text: 'Batch short, independent questions into a single request to reduce per-request overhead.' },
				{ ok: false, text: 'Don\'t send redundant context — only include the conversation history that the model actually needs.' },
			]
		},
	] as section}
		<div class="mb-8">
			<div class="flex items-center gap-2 mb-4">
				<span class="text-lg {section.color}">{section.icon}</span>
				<h2 class="text-base font-semibold text-[#0f0f0f]">{section.title}</h2>
			</div>
			<div class="space-y-2">
				{#each section.tips as tip}
					<div class="flex items-start gap-3 p-3 rounded-xl border {section.border} {section.bg}">
						<span class="{tip.ok ? section.color : 'text-[#aaa]'} font-bold text-sm flex-shrink-0">{tip.ok ? '✓' : '✗'}</span>
						<p class="text-sm text-[#444]">{tip.text}</p>
					</div>
				{/each}
			</div>
		</div>
	{/each}

	<!-- Summary card -->
	<div class="bg-[#f4f4f4] border border-[#e5e5e5] rounded-2xl p-6 mt-2">
		<p class="text-sm font-semibold text-[#0f0f0f] mb-2">Quick-start checklist</p>
		<ul class="space-y-1.5 text-sm text-[#555]">
			{#each [
				'API key stored in an environment variable',
				'Client timeout set to 120s',
				'Exponential backoff on 502 / 429 errors',
				'finish_reason checked on every response',
				'usage.total_tokens logged per request',
			] as item}
				<li class="flex items-center gap-2">
					<span class="w-4 h-4 rounded border-2 border-[#0f0f0f] flex items-center justify-center flex-shrink-0">
						<span class="w-2 h-2 rounded-sm bg-[#0f0f0f]"></span>
					</span>
					{item}
				</li>
			{/each}
		</ul>
	</div>

</div>
