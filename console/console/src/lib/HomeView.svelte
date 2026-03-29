<script lang="ts">
	import { X } from 'lucide-svelte';

	export let showBanner: boolean;
	export let keysLoading: boolean;
	export let totalRequests: number;
	export let totalCompletions: number;
	export let keys: any[];
	export let onNavigateDocs: () => void;
	export let onNavigateKeys: () => void;
</script>

<!-- Banner -->
{#if showBanner}
<div class="relative overflow-hidden rounded-2xl border border-[#e5e5e5] bg-[#fdfdfd] mb-6 sm:mb-10">
	<div class="absolute inset-y-0 right-0 w-[50%] bg-gradient-to-l from-[#ffb38a]/40 via-[#ff8eb3]/20 to-transparent pointer-events-none"></div>
	<div class="absolute inset-y-0 right-0 w-[30%] bg-gradient-to-bl from-[#ffdfa0]/60 via-transparent to-transparent pointer-events-none"></div>

	<div class="relative z-10 p-5 sm:p-8">
		<button on:click={() => showBanner = false} class="absolute top-4 right-4 w-8 h-8 flex items-center justify-center rounded-full bg-[#fce8df] text-[#c25e3a] hover:bg-[#fad8c9] transition-colors">
			<X class="w-4 h-4" strokeWidth={2.5} />
		</button>

		<h2 class="text-[18px] sm:text-[22px] font-semibold tracking-tight text-[#0f0f0f] mb-2 pr-10">
			Access top AI models with one OpenAI-compatible API
		</h2>
		<p class="text-[14px] sm:text-[15px] text-[#666666] max-w-[600px] leading-relaxed">
			Use Bielik, Mistral, DeepSeek, Kimi, GPT-OSS and more through a single <code class="font-mono bg-[#f4f4f4] px-1.5 py-0.5 rounded text-[13px]">/v1/chat/completions</code> endpoint.
		</p>

		<div class="flex flex-wrap items-center gap-3 mt-5 sm:mt-6">
			<button class="bg-[#0f0f0f] text-white px-4 py-2 rounded-full text-sm font-medium hover:bg-[#222222] transition-colors">
				Get started
			</button>
			<button on:click={onNavigateDocs} class="bg-white border border-[#e5e5e5] text-[#0f0f0f] px-4 py-2 rounded-full text-sm font-medium hover:bg-gray-50 transition-colors">
				Read the docs
			</button>
		</div>
	</div>
</div>
{/if}

<!-- Home Section -->
<div class="mb-8 sm:mb-12">
	<div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mb-4">
		<h3 class="text-xl font-semibold text-[#0f0f0f]">Home</h3>
		<div class="flex flex-wrap items-center gap-2 sm:gap-3">
			<button on:click={onNavigateKeys} class="flex items-center gap-2 border border-[#e5e5e5] bg-white px-3 py-1.5 rounded-lg text-sm font-medium text-[#0f0f0f] hover:bg-gray-50 transition-colors">
				<svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor"><path d="M16 9.75C16.9665 9.75 17.75 8.9665 17.75 8C17.75 7.0335 16.9665 6.25 16 6.25C15.0335 6.25 14.25 7.0335 14.25 8C14.25 8.9665 15.0335 9.75 16 9.75Z" /><path d="M15 2C11.134 2 8 5.13401 8 9C8 9.49204 8.05092 9.97307 8.14801 10.4378L3.73223 14.8536C3.26339 15.3224 3 15.9583 3 16.6213V20C3 20.5523 3.44772 21 4 21H7.37868C8.04172 21 8.6776 20.7366 9.14645 20.2678L10.2071 19.2071C10.3946 19.0196 10.5 18.7652 10.5 18.5V17.5H11.5C11.7652 17.5 12.0196 17.3946 12.2071 17.2071L13.5622 15.852C14.0269 15.9491 14.508 16 15 16C18.866 16 22 12.866 22 9C22 5.13401 18.866 2 15 2ZM10 9C10 6.23858 12.2386 4 15 4C17.7614 4 20 6.23858 20 9C20 11.7614 17.7614 14 15 14C14.4932 14 14.0057 13.9249 13.5471 13.7859C13.1941 13.6789 12.8108 13.775 12.55 14.0358L11.0858 15.5H9.5C8.94772 15.5 8.5 15.9477 8.5 16.5V18.0858L7.73223 18.8536C7.63847 18.9473 7.51129 19 7.37868 19H5V16.6213C5 16.4887 5.05268 16.3615 5.14645 16.2678L9.96418 11.45C10.225 11.1892 10.3211 10.8059 10.2141 10.4529C10.0751 9.99431 10 9.50683 10 9Z" /></svg>
				Create API keys
			</button>
			<div class="flex items-center bg-[#f4f4f4] p-1 rounded-lg">
				<button class="px-3 py-1 text-sm font-medium bg-white shadow-sm rounded-md text-[#0f0f0f]">24h</button>
				<button class="px-3 py-1 text-sm font-medium text-[#666666] hover:text-[#0f0f0f] transition-colors">7d</button>
				<button class="px-3 py-1 text-sm font-medium text-[#666666] hover:text-[#0f0f0f] transition-colors">30d</button>
				<button class="px-3 py-1 text-sm font-medium text-[#666666] hover:text-[#0f0f0f] transition-colors">90d</button>
			</div>
		</div>
	</div>

	<div class="border border-[#e5e5e5] rounded-2xl bg-white overflow-hidden">
		<div class="grid grid-cols-1 sm:grid-cols-2 border-b border-[#e5e5e5]">
			<div class="p-5 border-b sm:border-b-0 sm:border-r border-[#e5e5e5]">
				<div class="flex items-center text-sm text-[#666666] w-fit mb-3">
					Total tokens
				</div>
				<div class="text-2xl font-semibold text-[#0f0f0f] mb-4">—</div>
				<p class="text-xs text-[#aaaaaa]">Token tracking coming soon</p>
			</div>
			<div class="p-5">
				<div class="flex items-center text-sm text-[#666666] w-fit mb-3">
					Responses and Chat Completions
				</div>
				{#if keysLoading}
					<div class="w-16 h-6 bg-[#f4f4f4] rounded-full mb-4 animate-pulse"></div>
				{:else}
					<div class="text-2xl font-semibold text-[#0f0f0f] mb-4">{totalCompletions.toLocaleString()}</div>
				{/if}
				<div class="w-full h-2 bg-[#f4f4f4] rounded-full overflow-hidden flex">
					{#if totalRequests > 0}
						{#each keys.filter((k: any) => (k.requests_count ?? 0) > 0) as k, i}
							<div
								class="h-full"
								style="width:{((k.requests_count ?? 0) / totalRequests * 100).toFixed(1)}%; background: {['#0f0f0f','#555','#888','#aaa','#ccc'][i % 5]}"
							></div>
						{/each}
					{/if}
				</div>
			</div>
		</div>
		<div class="p-5">
			<div class="flex items-center text-sm text-[#666666] w-fit mb-3">
				Total requests
			</div>
			{#if keysLoading}
				<div class="w-16 h-6 bg-[#f4f4f4] rounded-full mb-4 animate-pulse"></div>
			{:else}
				<div class="text-2xl font-semibold text-[#0f0f0f] mb-4">{totalRequests.toLocaleString()}</div>
			{/if}
			<div class="flex gap-1.5 flex-wrap">
				{#each keys as k, i}
					<div class="flex items-center gap-1.5 text-xs text-[#666666]">
						<span class="w-2 h-2 rounded-full flex-shrink-0" style="background: {['#0f0f0f','#555','#888','#aaa','#ccc'][i % 5]}"></span>
						<span class="truncate max-w-[80px]">{k.name}</span>
						<span class="text-[#aaa]">{(k.requests_count ?? 0).toLocaleString()}</span>
					</div>
				{/each}
				{#if keys.length === 0}
					<span class="text-xs text-[#aaaaaa]">No API keys yet</span>
				{/if}
			</div>
		</div>
	</div>
</div>

<!-- Available Models Section -->
<div>
	<h3 class="text-lg sm:text-xl font-semibold text-[#0f0f0f] mb-4">Recommended models</h3>
	<div class="flex gap-3 overflow-x-auto pb-2">

		{#each [
			{ alias: 'Bielik-11b', id: 'speakleash/bielik-11b-v2.6-instruct' },
			{ alias: 'Mistral-Small-4', id: 'mistralai/mistral-small-4-119b-2603' },
			{ alias: 'DeepSeek-V3.1', id: 'deepseek-ai/deepseek-v3.1' },
			{ alias: 'Kimi-K2', id: 'moonshotai/kimi-k2-instruct' },
			{ alias: 'Amazon-Nova-2-lite-v1', id: 'nova-2-lite-v1' },
			{ alias: 'Minimax-m2.5', id: 'minimaxai/minimax-m2.5' },
			{ alias: 'GLM-4.7', id: 'z-ai/glm4.7' },
			{ alias: 'GPT-OSS-120b', id: 'openai/gpt-oss-120b' },
			{ alias: 'Step-3.5-Flash', id: 'stepfun-ai/step-3.5-flash' },
			{ alias: 'Qwen-3.5', id: 'qwen/qwen3.5-122b-a10b' },
			{ alias: 'Kimi-K2.5', id: 'moonshotai/kimi-k2.5' },
		] as model}
		<div class="flex-shrink-0 w-48 border border-[#e5e5e5] rounded-2xl p-5 bg-white hover:shadow-sm transition-shadow cursor-pointer">
			<div class="w-9 h-9 border border-[#e5e5e5] rounded-xl flex items-center justify-center mb-3">
				<svg class="w-5 h-5 text-[#0f0f0f]" width="24" height="24" viewBox="0 0 24 24" fill="currentColor"><path fill-rule="evenodd" d="M12.5 3.444a1 1 0 0 0-1 0l-6.253 3.61 6.768 3.807 6.955-3.682-6.47-3.735Zm7.16 5.632L13 12.602v7.666l6.16-3.556a1 1 0 0 0 .5-.867V9.076ZM11 20.268v-7.683L4.34 8.839v7.006a1 1 0 0 0 .5.867L11 20.268Zm-.5-18.557a3 3 0 0 1 3 0l6.66 3.846a3 3 0 0 1 1.5 2.598v7.69a3 3 0 0 1-1.5 2.598L13.5 22.29a3 3 0 0 1-3 0l-6.66-3.846a3 3 0 0 1-1.5-2.598v-7.69a3 3 0 0 1 1.5-2.598L10.5 1.71Z" clip-rule="evenodd" /></svg>
			</div>
			<h4 class="text-[15px] font-semibold text-[#0f0f0f] mb-1">{model.alias}</h4>
			<p class="text-[12px] text-[#888888] font-mono truncate">{model.id}</p>
		</div>
		{/each}

	</div>
</div>
