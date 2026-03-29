<script lang="ts">
	import { getSnippets } from '$lib/docs-store';
	import { baseUrl } from '$lib/docs-store';
	import CodeBlock from '$lib/CodeBlock.svelte';

	$: s = getSnippets($baseUrl);

	const models = [
		{ alias: 'Bielik-11b',           id: 'speakleash/bielik-11b-v2.6-instruct',  provider: 'NVIDIA',       desc: 'Polish-language optimized 11B model' },
		{ alias: 'Mistral-Small-4',       id: 'mistralai/mistral-small-4-119b-2603',  provider: 'NVIDIA',       desc: 'Fast & efficient 119B model' },
		{ alias: 'DeepSeek-V3.1',         id: 'deepseek-ai/deepseek-v3.1',            provider: 'NVIDIA',       desc: 'State-of-the-art reasoning model' },
		{ alias: 'Kimi-K2',               id: 'moonshotai/kimi-k2-instruct',          provider: 'NVIDIA',       desc: 'Long-context general-purpose model' },
		{ alias: 'Kimi-K2.5',             id: 'moonshotai/kimi-k2.5',                 provider: 'NVIDIA',       desc: 'Updated Kimi K2 with improvements' },
		{ alias: 'Amazon-Nova-2-lite-v1', id: 'nova-2-lite-v1',                       provider: 'Amazon Nova',  desc: 'Fast, lightweight Amazon Nova model' },
		{ alias: 'Minimax-m2.5',          id: 'minimaxai/minimax-m2.5',               provider: 'NVIDIA',       desc: 'MiniMax multimodal model' },
		{ alias: 'GLM-4.7',               id: 'z-ai/glm4.7',                          provider: 'NVIDIA',       desc: 'ZHIPU AI bilingual model' },
		{ alias: 'GPT-OSS-120b',          id: 'openai/gpt-oss-120b',                  provider: 'NVIDIA',       desc: 'Open-source 120B model' },
		{ alias: 'Step-3.5-Flash',        id: 'stepfun-ai/step-3.5-flash',            provider: 'NVIDIA',       desc: 'Ultra-fast inference model' },
		{ alias: 'Qwen-3.5',              id: 'qwen/qwen3.5-122b-a10b',               provider: 'NVIDIA',       desc: 'Alibaba 122B MoE — strong at code' },
	];
</script>

<svelte:head><title>Available Models — API Docs</title></svelte:head>

<div class="max-w-[760px]">

	<h1 class="text-[1.75rem] font-bold text-[#0f0f0f] mb-2">Available Models</h1>
	<p class="text-sm text-[#666] mb-8 leading-relaxed">
		Use the short <strong class="text-[#333]">alias</strong> or the full model ID in the <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded">model</code> field — both are accepted. Model aliases are resolved server-side.
	</p>

	<!-- Providers -->
	<div class="grid grid-cols-2 gap-3 mb-8">
		{#each [
			{ name: 'NVIDIA NIM', count: 10, color: 'bg-[#f4f4f4] border-[#e5e5e5] text-[#555]', dot: 'bg-[#888]' },
			{ name: 'Amazon Nova', count: 1, color: 'bg-orange-50 border-orange-100 text-orange-700', dot: 'bg-orange-500' },
		] as p}
			<div class="border {p.color} rounded-2xl p-4 flex items-center gap-3">
				<span class="w-2 h-2 rounded-full {p.dot}"></span>
				<div>
					<p class="text-sm font-semibold">{p.name}</p>
					<p class="text-xs opacity-70">{p.count} model{p.count > 1 ? 's' : ''}</p>
				</div>
			</div>
		{/each}
	</div>

	<!-- Models table -->
	<div class="rounded-2xl border border-[#e5e5e5] overflow-hidden mb-8">
		<table class="w-full text-sm">
			<thead>
				<tr class="bg-[#f9f9f9] border-b border-[#e5e5e5]">
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider">Alias</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider hidden lg:table-cell">Full model ID</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider hidden sm:table-cell">Provider</th>
					<th class="text-left px-4 py-3 text-xs font-semibold text-[#888] uppercase tracking-wider hidden md:table-cell">Description</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-[#f0f0f0]">
				{#each models as m}
					<tr class="hover:bg-[#fafafa] transition-colors">
						<td class="px-4 py-3">
							<code class="text-[12px] font-mono font-semibold text-[#0f0f0f] bg-[#f4f4f4] px-2 py-0.5 rounded">{m.alias}</code>
						</td>
						<td class="px-4 py-3 font-mono text-[11px] text-[#aaa] hidden lg:table-cell break-all">{m.id}</td>
						<td class="px-4 py-3 hidden sm:table-cell">
							<span class="text-[11px] px-2 py-0.5 rounded-full font-medium
								{m.provider === 'Amazon Nova'
									? 'bg-orange-50 text-orange-600 border border-orange-100'
									: 'bg-[#f4f4f4] text-[#555] border border-[#e5e5e5]'}">
								{m.provider}
							</span>
						</td>
						<td class="px-4 py-3 text-[12px] text-[#666] hidden md:table-cell">{m.desc}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<!-- List models API -->
	<h2 class="text-base font-semibold text-[#0f0f0f] mb-3">List models via API</h2>
	<p class="text-sm text-[#666] mb-1">Fetch all available models programmatically:</p>
	<CodeBlock code={s.curlModels} id="models-curl" label="GET /v1/models" method="GET" lang="bash" />

	<p class="text-sm text-[#666] mt-4 mb-1">Example response:</p>
	<CodeBlock code={s.curlModelsResponse} id="models-resp" lang="json" />

	<!-- Notes -->
	<div class="flex items-start gap-3 p-4 rounded-xl border border-[#e5e5e5] bg-[#f4f4f4] mt-4">
		<svg class="w-4 h-4 text-[#888] flex-shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
		<p class="text-sm text-[#444]">
			Models are routed to NVIDIA NIM by default. <code class="bg-[#e5e5e5] px-1 rounded font-mono">Amazon-Nova-2-lite-v1</code> is routed to the Amazon Nova endpoint automatically — no extra configuration needed.
		</p>
	</div>

</div>
