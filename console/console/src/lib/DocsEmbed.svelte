<script lang="ts">
	import { onMount } from 'svelte';
	import { baseUrl, NAV, FLAT_NAV } from '$lib/docs-store';

	import IntroductionPage from '../routes/docs/introduction/+page.svelte';
	import QuickstartPage from '../routes/docs/quickstart/+page.svelte';
	import ChangelogPage from '../routes/docs/changelog/+page.svelte';
	import AuthenticationPage from '../routes/docs/authentication/+page.svelte';
	import ModelsPage from '../routes/docs/models/+page.svelte';
	import ChatCompletionsPage from '../routes/docs/chat-completions/+page.svelte';
	import ParametersPage from '../routes/docs/parameters/+page.svelte';
	import ResponseFormatPage from '../routes/docs/response-format/+page.svelte';
	import StreamingPage from '../routes/docs/streaming/+page.svelte';
	import RateLimitsPage from '../routes/docs/rate-limits/+page.svelte';
	import StructuredOutputsPage from '../routes/docs/structured-outputs/+page.svelte';
	import ToolCallsPage from '../routes/docs/tool-calls/+page.svelte';
	import AgenticLoopPage from '../routes/docs/agentic-loop/+page.svelte';
	import ErrorsPage from '../routes/docs/errors/+page.svelte';
	import BestPracticesPage from '../routes/docs/best-practices/+page.svelte';

	export let activePage: string = 'introduction';

	let mobileNavOpen = false;

	const PAGES: Record<string, any> = {
		introduction: IntroductionPage,
		quickstart: QuickstartPage,
		changelog: ChangelogPage,
		authentication: AuthenticationPage,
		models: ModelsPage,
		'chat-completions': ChatCompletionsPage,
		parameters: ParametersPage,
		'response-format': ResponseFormatPage,
		streaming: StreamingPage,
		'rate-limits': RateLimitsPage,
		'structured-outputs': StructuredOutputsPage,
		'tool-calls': ToolCallsPage,
		'agentic-loop': AgenticLoopPage,
		errors: ErrorsPage,
		'best-practices': BestPracticesPage,
	};

	$: currentIdx = FLAT_NAV.findIndex(n => n.id === activePage);
	$: prev = currentIdx > 0 ? FLAT_NAV[currentIdx - 1] : null;
	$: next = currentIdx < FLAT_NAV.length - 1 ? FLAT_NAV[currentIdx + 1] : null;
	$: CurrentComponent = PAGES[activePage] ?? IntroductionPage;

	function navigate(id: string) {
		activePage = id;
		mobileNavOpen = false;
		scrollTop();
	}

	let contentEl: HTMLElement;
	function scrollTop() {
		setTimeout(() => contentEl?.scrollTo({ top: 0, behavior: 'smooth' }), 0);
	}

	function handleContentClick(e: MouseEvent) {
		const link = (e.target as HTMLElement).closest('a');
		if (!link) return;
		const href = link.getAttribute('href');
		if (href?.startsWith('/docs/')) {
			const id = href.replace('/docs/', '').replace(/\/$/, '');
			if (PAGES[id]) {
				e.preventDefault();
				navigate(id);
			}
		}
	}

	onMount(() => {
		const host = window.location.hostname;
		if (host === 'localhost' || host === '127.0.0.1') {
			baseUrl.set('http://localhost:3001/v1');
		} else {
			baseUrl.set('https://' + host + '/v1');
		}
	});
</script>

<div class="flex flex-1 min-h-0 overflow-hidden bg-white dark:bg-gray-950 rounded-tl-[1.5rem]">

	<!-- Mobile nav overlay -->
	{#if mobileNavOpen}
		<div
			class="fixed inset-0 z-30 bg-black/20 dark:bg-black/50 md:hidden"
			on:click={() => mobileNavOpen = false}
			role="presentation"
		></div>
	{/if}

	<!-- Docs Sidebar -->
	<aside class="
		{mobileNavOpen ? 'fixed inset-y-0 left-0 z-40 translate-x-0' : 'fixed inset-y-0 left-0 z-40 -translate-x-full'}
		md:relative md:translate-x-0 md:z-auto
		w-60 flex-shrink-0 border-r border-[#f0f0f0] dark:border-gray-900 bg-white dark:bg-gray-950
		flex flex-col
		transition-transform duration-200
		md:h-full md:overflow-hidden
	">
		<!-- Nav -->
		<nav class="flex-1 overflow-y-auto px-3 py-4 space-y-5">
			{#each NAV as group}
				<div>
					<p class="text-[10px] font-semibold text-[#bbbbbb] dark:text-gray-700 uppercase tracking-wider mb-1.5 px-2">{group.group}</p>
					{#each group.items as item}
						<button
							on:click={() => navigate(item.id)}
							class="w-full flex items-center gap-2 px-3 py-2 rounded-lg text-sm transition-all text-left
								{activePage === item.id
									? 'bg-[#ebebeb] dark:bg-gray-800 text-[#0f0f0f] dark:text-gray-100 font-semibold'
									: 'text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:bg-[#e5e5e5]/40 dark:hover:bg-gray-800/60'}"
						>
							{#if activePage === item.id}
								<span class="w-1.5 h-1.5 rounded-full bg-[#0f0f0f] dark:bg-gray-300 flex-shrink-0"></span>
							{:else}
								<span class="w-1.5 h-1.5 rounded-full bg-transparent flex-shrink-0"></span>
							{/if}
							{item.label}
						</button>
					{/each}
				</div>
			{/each}
		</nav>

		<div class="px-5 py-3 border-t border-[#f0f0f0] dark:border-gray-900 flex-shrink-0">
			<p class="text-[10px] text-[#cccccc] dark:text-gray-800 text-center">AI Gateway API · v1</p>
		</div>
	</aside>

	<!-- Content -->
	<div class="flex-1 min-w-0 flex flex-col md:h-full md:overflow-hidden">

		<!-- Mobile top bar -->
		<div class="md:hidden flex items-center gap-3 px-4 py-3 border-b border-[#f0f0f0] dark:border-gray-900 flex-shrink-0 bg-white dark:bg-gray-950">
			<button
				on:click={() => mobileNavOpen = !mobileNavOpen}
				class="p-1.5 rounded-lg hover:bg-[#f4f4f4] dark:hover:bg-gray-800 transition-colors"
				aria-label="Toggle docs menu"
			>
				<svg class="w-5 h-5 text-[#444] dark:text-gray-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
					<line x1="3" y1="6" x2="21" y2="6"/>
					<line x1="3" y1="12" x2="21" y2="12"/>
					<line x1="3" y1="18" x2="21" y2="18"/>
				</svg>
			</button>
			<span class="text-sm font-semibold text-[#0f0f0f] dark:text-gray-100">
				{FLAT_NAV.find(n => n.id === activePage)?.label ?? 'API Docs'}
			</span>
			<span class="ml-auto flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-[#f4f4f4] dark:bg-gray-900 text-xs text-[#666666] dark:text-gray-500 font-medium">
				<span class="w-1.5 h-1.5 rounded-full bg-green-500"></span>
				v1
			</span>
		</div>

		<!-- Scrollable content -->
		<main bind:this={contentEl} on:click={handleContentClick} class="flex-1 overflow-y-auto px-6 md:px-12 py-8 md:py-12 bg-white dark:bg-gray-950">
			<svelte:component this={CurrentComponent} />

			<!-- Prev / Next -->
			{#if prev || next}
				<div class="mt-16 pt-8 border-t border-[#f0f0f0] dark:border-gray-900 flex items-center justify-between gap-4">
					{#if prev}
						<button
							on:click={() => navigate(prev.id)}
							class="flex items-center gap-2 px-4 py-3 rounded-xl border border-[#e5e5e5] dark:border-gray-850 text-sm text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:border-[#cccccc] dark:hover:border-gray-700 hover:bg-[#fafafa] dark:hover:bg-gray-900 transition-all group"
						>
							<svg class="w-4 h-4 group-hover:-translate-x-0.5 transition-transform" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
							<div class="text-left">
								<p class="text-[10px] text-[#aaa] dark:text-gray-700 uppercase tracking-wide">Previous</p>
								<p class="font-medium text-[#222] dark:text-gray-200">{prev.label}</p>
							</div>
						</button>
					{:else}
						<div></div>
					{/if}
					{#if next}
						<button
							on:click={() => navigate(next.id)}
							class="flex items-center gap-2 px-4 py-3 rounded-xl border border-[#e5e5e5] dark:border-gray-850 text-sm text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:border-[#cccccc] dark:hover:border-gray-700 hover:bg-[#fafafa] dark:hover:bg-gray-900 transition-all group"
						>
							<div class="text-right">
								<p class="text-[10px] text-[#aaa] dark:text-gray-700 uppercase tracking-wide">Next</p>
								<p class="font-medium text-[#222] dark:text-gray-200">{next.label}</p>
							</div>
							<svg class="w-4 h-4 group-hover:translate-x-0.5 transition-transform" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M12 5l7 7-7 7"/></svg>
						</button>
					{/if}
				</div>
			{/if}
		</main>
	</div>
</div>
