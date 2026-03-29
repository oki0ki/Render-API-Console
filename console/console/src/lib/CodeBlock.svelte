<script lang="ts">
	import { copyText, copiedId } from './docs-store';

	export let code: string;
	export let id: string;
	export let label: string = '';
	export let method: 'GET' | 'POST' | '' = '';
	export let lang: string = '';
</script>

<div class="rounded-2xl border border-[#e5e5e5] dark:border-gray-850 overflow-hidden">
	{#if label || method}
		<div class="flex items-center justify-between px-4 py-2.5 bg-[#f9f9f9] dark:bg-gray-900 border-b border-[#e5e5e5] dark:border-gray-850">
			<div class="flex items-center gap-2">
				{#if method}
					<span class="text-xs font-bold font-mono text-white px-2 py-0.5 rounded
						{method === 'POST' ? 'bg-green-500' : 'bg-[#555] dark:bg-gray-700'}">
						{method}
					</span>
				{/if}
				{#if label}
					<span class="text-xs font-medium text-[#888888] dark:text-gray-600">{label}</span>
				{/if}
			</div>
			{#if lang}
				<span class="text-[10px] font-mono text-[#cccccc] dark:text-gray-700 uppercase">{lang}</span>
			{/if}
		</div>
	{/if}
	<div class="relative">
		<pre class="p-4 text-sm font-mono text-[#0f0f0f] dark:text-gray-200 overflow-x-auto bg-white dark:bg-gray-950"><code>{@html code}</code></pre>
		<button
			on:click={() => copyText(id, code)}
			class="absolute top-3 right-3 p-1.5 rounded-lg border border-[#e5e5e5] dark:border-gray-800 text-[#888] dark:text-gray-600 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:border-[#ccc] dark:hover:border-gray-700 transition-all text-xs font-mono bg-white dark:bg-gray-950"
			title="Copy to clipboard"
		>
			{$copiedId === id ? '✓' : '⎘'}
		</button>
	</div>
</div>
