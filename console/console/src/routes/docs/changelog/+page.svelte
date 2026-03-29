<svelte:head><title>Changelog — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Changelog</h1>
	<p class="text-sm text-[#666] mb-10 leading-relaxed">
		A running log of updates, new models, and improvements to the AI Gateway API.
	</p>

	<div class="relative">
		<!-- Timeline line -->
		<div class="absolute left-[7px] top-2 bottom-0 w-px bg-[#e5e5e5]"></div>

		<div class="space-y-10">

			{#each [
				{
					date: 'March 2026',
					tag: 'Latest',
					tagColor: 'bg-[#f4f4f4] text-[#555] border-[#e5e5e5]',
					items: [
						{ type: 'new',  text: 'Added Kimi-K2 model (Moonshot AI, 1T MoE, 32B active parameters).' },
						{ type: 'new',  text: 'Agentic Loop now supports up to 10 automatic tool-call iterations.' },
						{ type: 'fix',  text: 'Fixed streaming chunk ordering for DeepSeek-V3.1 under high load.' },
						{ type: 'new',  text: 'Added x-api-key header as an alternative to Authorization: Bearer.' },
					]
				},
				{
					date: 'February 2026',
					tag: null,
					tagColor: '',
					items: [
						{ type: 'new',  text: 'Added DeepSeek-V3.1 and Mistral-Small-4 models.' },
						{ type: 'new',  text: 'Introduced Tool Calls (function calling) support across all models.' },
						{ type: 'impr', text: 'Improved cold-start latency by ~40% for NVIDIA NIM-backed models.' },
						{ type: 'fix',  text: 'Corrected token count reporting in usage.completion_tokens for streamed responses.' },
					]
				},
				{
					date: 'January 2026',
					tag: null,
					tagColor: '',
					items: [
						{ type: 'new',  text: 'Public launch of AI Gateway API v1.' },
						{ type: 'new',  text: 'Initial model roster: 4 Amazon Nova variants + GPT-OSS-120b.' },
						{ type: 'new',  text: 'Full OpenAI SDK compatibility — drop-in base_url replacement.' },
						{ type: 'new',  text: 'Server-Sent Events (SSE) streaming for chat completions.' },
					]
				}
			] as entry}
				<div class="relative pl-7">
					<!-- Dot -->
					<div class="absolute left-0 top-1.5 w-3.5 h-3.5 rounded-full bg-white border-2 border-[#d1d5db] flex items-center justify-center">
						<div class="w-1.5 h-1.5 rounded-full bg-[#9ca3af]"></div>
					</div>

					<div class="flex items-center gap-2 mb-3">
						<p class="text-sm font-semibold text-[#111]">{entry.date}</p>
						{#if entry.tag}
							<span class="text-[10px] font-semibold px-2 py-0.5 rounded-full border {entry.tagColor}">{entry.tag}</span>
						{/if}
					</div>

					<div class="space-y-2">
						{#each entry.items as item}
							<div class="flex items-start gap-2.5 p-3 rounded-xl bg-[#fafafa] border border-[#f0f0f0]">
								{#if item.type === 'new'}
									<span class="flex-shrink-0 text-[10px] font-bold px-1.5 py-0.5 rounded bg-green-100 text-green-700 mt-0.5">NEW</span>
								{:else if item.type === 'fix'}
									<span class="flex-shrink-0 text-[10px] font-bold px-1.5 py-0.5 rounded bg-amber-100 text-amber-700 mt-0.5">FIX</span>
								{:else}
									<span class="flex-shrink-0 text-[10px] font-bold px-1.5 py-0.5 rounded bg-[#f4f4f4] text-[#666] mt-0.5">IMPR</span>
								{/if}
								<p class="text-sm text-[#444]">{item.text}</p>
							</div>
						{/each}
					</div>
				</div>
			{/each}

		</div>
	</div>

</div>
