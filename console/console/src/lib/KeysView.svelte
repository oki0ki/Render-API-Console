<script lang="ts">
	import { X } from 'lucide-svelte';

	export let keys: any[];
	export let keysLoading: boolean;
	export let creating: boolean;
	export let newlyCreatedKey: string | null;
	export let copiedId: string | null;
	export let onCreateKey: () => void;
	export let onDeleteKey: (id: string) => void;
	export let onCopy: (text: string, id: string) => void;
	export let fmtDate: (iso: string) => string;
</script>

<!-- API Keys View -->
<div class="mb-6 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
	<div>
		<h3 class="text-xl font-semibold text-[#0f0f0f]">API keys</h3>
		<p class="text-sm text-[#666666] mt-1">Your secret API keys for authenticating requests.</p>
	</div>
	<button on:click={onCreateKey} disabled={creating} class="self-start sm:self-auto flex items-center gap-2 bg-[#0f0f0f] text-white px-4 py-2 rounded-full text-sm font-medium hover:bg-[#222222] disabled:opacity-40 transition-colors">
		<svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor"><path d="M16 9.75C16.9665 9.75 17.75 8.9665 17.75 8C17.75 7.0335 16.9665 6.25 16 6.25C15.0335 6.25 14.25 7.0335 14.25 8C14.25 8.9665 15.0335 9.75 16 9.75Z" /><path d="M15 2C11.134 2 8 5.13401 8 9C8 9.49204 8.05092 9.97307 8.14801 10.4378L3.73223 14.8536C3.26339 15.3224 3 15.9583 3 16.6213V20C3 20.5523 3.44772 21 4 21H7.37868C8.04172 21 8.6776 20.7366 9.14645 20.2678L10.2071 19.2071C10.3946 19.0196 10.5 18.7652 10.5 18.5V17.5H11.5C11.7652 17.5 12.0196 17.3946 12.2071 17.2071L13.5622 15.852C14.0269 15.9491 14.508 16 15 16C18.866 16 22 12.866 22 9C22 5.13401 18.866 2 15 2ZM10 9C10 6.23858 12.2386 4 15 4C17.7614 4 20 6.23858 20 9C20 11.7614 17.7614 14 15 14C14.4932 14 14.0057 13.9249 13.5471 13.7859C13.1941 13.6789 12.8108 13.775 12.55 14.0358L11.0858 15.5H9.5C8.94772 15.5 8.5 15.9477 8.5 16.5V18.0858L7.73223 18.8536C7.63847 18.9473 7.51129 19 7.37868 19H5V16.6213C5 16.4887 5.05268 16.3615 5.14645 16.2678L9.96418 11.45C10.225 11.1892 10.3211 10.8059 10.2141 10.4529C10.0751 9.99431 10 9.50683 10 9Z" /></svg>
		{creating ? 'Creating…' : 'Create new secret key'}
	</button>
</div>

{#if newlyCreatedKey}
<div class="border border-[#e5e5e5] rounded-2xl p-5 bg-white mb-6">
	<p class="text-sm font-semibold text-[#0f0f0f] mb-1">Save your key</p>
	<p class="text-sm text-[#666666] mb-3">This won't be shown again. Copy it now and store it safely.</p>
	<div class="flex items-center gap-2 bg-[#f4f4f4] rounded-lg px-4 py-3 mb-3">
		<code class="text-sm font-mono text-[#0f0f0f] flex-1 break-all">{newlyCreatedKey}</code>
		<button on:click={() => onCopy(newlyCreatedKey ?? '', 'new-key')} class="text-[#888888] hover:text-[#0f0f0f] transition-colors text-sm">
			{copiedId === 'new-key' ? 'Copied' : 'Copy'}
		</button>
	</div>
	<button on:click={() => newlyCreatedKey = null} class="text-sm text-[#888888] hover:text-[#0f0f0f] transition-colors">Dismiss</button>
</div>
{/if}

<div class="border border-[#e5e5e5] rounded-2xl bg-white overflow-hidden">
	<div class="grid grid-cols-[1fr_auto] sm:grid-cols-[1fr_auto_auto] gap-4 px-5 py-3 border-b border-[#e5e5e5] bg-[#f9f9f9]">
		<p class="text-xs font-medium text-[#888888] uppercase tracking-wide">Name</p>
		<p class="hidden sm:block text-xs font-medium text-[#888888] uppercase tracking-wide">Created</p>
		<p class="w-8"></p>
	</div>
	{#if keysLoading}
	<div class="flex items-center justify-center py-12">
		<div class="w-4 h-4 rounded-full border-[1.5px] border-black/10 border-t-black animate-spin"></div>
	</div>
	{:else if keys.length === 0}
	<div class="flex flex-col items-center py-12 gap-2">
		<p class="text-sm font-medium text-[#444444]">No API keys yet</p>
		<p class="text-sm text-[#888888]">Create your first key to start using the API</p>
	</div>
	{:else}
	{#each keys as k, i}
	<div class="grid grid-cols-[1fr_auto] sm:grid-cols-[1fr_auto_auto] gap-4 items-center px-5 py-4 {i < keys.length - 1 ? 'border-b border-[#e5e5e5]' : ''} hover:bg-gray-50 transition-colors group">
		<div>
			<p class="text-sm font-medium text-[#0f0f0f]">{k.name ?? 'API Key'}</p>
			<div class="flex items-center gap-1.5 mt-0.5">
				<code class="text-xs sm:text-sm font-mono text-[#666666] truncate max-w-[180px] sm:max-w-none">sk-••••••••••••••••••••</code>
				<button on:click={() => onCopy(k.key ?? '', k.id)} class="text-xs text-[#888888] hover:text-[#0f0f0f] transition-colors sm:opacity-0 sm:group-hover:opacity-100">
					{copiedId === k.id ? 'Copied' : 'Copy'}
				</button>
			</div>
		</div>
		<p class="hidden sm:block text-sm text-[#666666] whitespace-nowrap">{fmtDate(k.created_at)}</p>
		<button on:click={() => onDeleteKey(k.id)} class="w-8 h-8 flex items-center justify-center rounded-md text-[#888888] hover:text-red-500 hover:bg-red-50 transition-colors sm:opacity-0 sm:group-hover:opacity-100">
			<X class="w-4 h-4" strokeWidth={2} />
		</button>
	</div>
	{/each}
	{/if}
</div>
