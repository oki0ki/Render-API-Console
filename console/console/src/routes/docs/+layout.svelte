<script lang="ts">
        import { onMount } from 'svelte';
        import { goto, beforeNavigate } from '$app/navigation';
        import { page } from '$app/stores';
        import { baseUrl, copyText, copiedId, NAV, FLAT_NAV } from '$lib/docs-store';

        let mobileOpen = false;

        onMount(() => {
                const host = window.location.hostname;
                if (host === 'localhost' || host === '127.0.0.1') {
                        baseUrl.set('http://localhost:3001/v1');
                } else {
                        baseUrl.set('https://' + host + '/v1');
                }
        });

        beforeNavigate(() => { mobileOpen = false; });

        $: currentPath = $page.url.pathname;
        $: currentIdx = FLAT_NAV.findIndex(n => n.path === currentPath);
        $: prev = currentIdx > 0 ? FLAT_NAV[currentIdx - 1] : null;
        $: next = currentIdx < FLAT_NAV.length - 1 ? FLAT_NAV[currentIdx + 1] : null;
</script>

<div class="min-h-screen bg-white dark:bg-gray-950 text-[#0f0f0f] dark:text-gray-100 font-sans">

        <!-- Top header -->
        <header class="sticky top-0 z-50 h-[60px] flex items-center justify-between px-5 border-b border-[#f0f0f0] dark:border-gray-900 bg-white/95 dark:bg-gray-950/95 backdrop-blur-sm">
                <div class="flex items-center gap-3">
                        <!-- Mobile menu toggle -->
                        <button
                                class="md:hidden p-1.5 rounded-lg hover:bg-[#f4f4f4] dark:hover:bg-gray-800 transition-colors"
                                on:click={() => mobileOpen = !mobileOpen}
                                aria-label="Toggle menu"
                        >
                                {#if mobileOpen}
                                        <svg class="w-5 h-5 text-[#444] dark:text-gray-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/></svg>
                                {:else}
                                        <svg class="w-5 h-5 text-[#444] dark:text-gray-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="3" y1="6" x2="21" y2="6"/><line x1="3" y1="12" x2="21" y2="12"/><line x1="3" y1="18" x2="21" y2="18"/></svg>
                                {/if}
                        </button>

                        <!-- Logo / breadcrumb -->
                        <button
                                on:click={() => goto('/')}
                                class="flex items-center gap-2 text-sm text-[#666666] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 transition-colors"
                        >
                                <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
                                Console
                        </button>
                        <span class="text-[#e5e5e5] dark:text-gray-800">/</span>
                        <span class="text-sm font-semibold text-[#0f0f0f] dark:text-gray-100">API Docs</span>
                        {#if currentIdx >= 0}
                                <span class="hidden sm:inline text-[#e5e5e5] dark:text-gray-800">/</span>
                                <span class="hidden sm:inline text-sm text-[#666666] dark:text-gray-500">{FLAT_NAV[currentIdx]?.label}</span>
                        {/if}
                </div>

                <div class="flex items-center gap-2">
                        <span class="flex items-center gap-1.5 px-3 py-1 rounded-full bg-[#f4f4f4] dark:bg-gray-900 text-xs text-[#666666] dark:text-gray-500 font-medium">
                                <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span>
                                v1
                        </span>
                        <span class="hidden sm:inline px-2.5 py-1 rounded-full bg-[#f4f4f4] dark:bg-gray-900 text-xs text-[#888888] dark:text-gray-600 font-medium border border-[#e5e5e5] dark:border-gray-850">OpenAI-compatible</span>
                </div>
        </header>

        <div class="flex max-w-[1260px] mx-auto relative md:h-[calc(100vh-60px)] md:overflow-hidden">

                <!-- Mobile overlay -->
                {#if mobileOpen}
                        <div
                                class="fixed inset-0 top-[60px] bg-black/20 dark:bg-black/50 z-30 md:hidden"
                                on:click={() => mobileOpen = false}
                                role="presentation"
                        ></div>
                {/if}

                <!-- Sidebar -->
                <aside class="
                        {mobileOpen ? 'fixed top-[60px] left-0 bottom-0 z-40 translate-x-0' : 'fixed top-[60px] left-0 bottom-0 z-40 -translate-x-full'}
                        md:relative md:translate-x-0 md:z-auto md:h-full md:overflow-y-auto
                        w-64 flex-shrink-0 border-r border-[#f0f0f0] dark:border-gray-900 bg-white dark:bg-gray-950
                        flex flex-col
                        transition-transform duration-200
                ">
                        <!-- Base URL -->
                        <div class="px-5 pt-6 pb-4 border-b border-[#f4f4f4] dark:border-gray-900">
                                <p class="text-[10px] font-semibold text-[#bbbbbb] dark:text-gray-700 uppercase tracking-wider mb-2">Base URL</p>
                                <button
                                        on:click={() => copyText('sidebar-url', $baseUrl)}
                                        class="w-full text-left group"
                                        title="Click to copy"
                                >
                                        <code class="block text-[11px] font-mono text-[#444] dark:text-gray-400 bg-[#f4f4f4] dark:bg-gray-900 px-2.5 py-2 rounded-lg break-all leading-snug group-hover:bg-[#eaeaea] dark:group-hover:bg-gray-800 transition-colors">
                                                {$baseUrl || 'Loading…'}
                                        </code>
                                        <span class="text-[10px] text-[#aaa] dark:text-gray-700 mt-1 block text-right pr-0.5">
                                                {$copiedId === 'sidebar-url' ? '✓ Copied' : 'Click to copy'}
                                        </span>
                                </button>
                        </div>

                        <!-- Navigation -->
                        <nav class="flex-1 px-3 py-4 space-y-5">
                                {#each NAV as group}
                                        <div>
                                                <p class="text-[10px] font-semibold text-[#bbbbbb] dark:text-gray-700 uppercase tracking-wider mb-1.5 px-2">{group.group}</p>
                                                {#each group.items as item}
                                                        <a
                                                                href={item.path}
                                                                class="flex items-center gap-2 px-3 py-2 rounded-xl text-sm transition-all
                                                                        {currentPath === item.path
                                                                                ? 'bg-[#ebebeb] dark:bg-gray-800 text-[#0f0f0f] dark:text-gray-100 font-semibold'
                                                                                : 'text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:bg-[#f9f9f9] dark:hover:bg-gray-900'}"
                                                        >
                                                                {#if currentPath === item.path}
                                                                        <span class="w-1.5 h-1.5 rounded-full bg-[#0f0f0f] dark:bg-gray-300 flex-shrink-0"></span>
                                                                {:else}
                                                                        <span class="w-1.5 h-1.5 rounded-full bg-transparent flex-shrink-0"></span>
                                                                {/if}
                                                                {item.label}
                                                        </a>
                                                {/each}
                                        </div>
                                {/each}
                        </nav>

                        <!-- Sidebar footer -->
                        <div class="px-5 py-4 border-t border-[#f4f4f4] dark:border-gray-900">
                                <p class="text-[10px] text-[#cccccc] dark:text-gray-800 text-center">AI Gateway API · v1</p>
                        </div>
                </aside>

                <!-- Main content -->
                <main class="flex-1 min-w-0 px-6 md:px-12 py-12 md:h-full md:overflow-y-auto">
                        <slot />

                        <!-- Prev / Next navigation -->
                        {#if prev || next}
                                <div class="mt-16 pt-8 border-t border-[#f0f0f0] dark:border-gray-900 flex items-center justify-between gap-4">
                                        {#if prev}
                                                <a
                                                        href={prev.path}
                                                        class="flex items-center gap-2 px-4 py-3 rounded-xl border border-[#e5e5e5] dark:border-gray-850 text-sm text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:border-[#cccccc] dark:hover:border-gray-700 hover:bg-[#fafafa] dark:hover:bg-gray-900 transition-all group"
                                                >
                                                        <svg class="w-4 h-4 group-hover:-translate-x-0.5 transition-transform" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M19 12H5M12 5l-7 7 7 7"/></svg>
                                                        <div class="text-left">
                                                                <p class="text-[10px] text-[#aaa] dark:text-gray-700 uppercase tracking-wide">Previous</p>
                                                                <p class="font-medium text-[#222] dark:text-gray-200">{prev.label}</p>
                                                        </div>
                                                </a>
                                        {:else}
                                                <div></div>
                                        {/if}
                                        {#if next}
                                                <a
                                                        href={next.path}
                                                        class="flex items-center gap-2 px-4 py-3 rounded-xl border border-[#e5e5e5] dark:border-gray-850 text-sm text-[#555] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 hover:border-[#cccccc] dark:hover:border-gray-700 hover:bg-[#fafafa] dark:hover:bg-gray-900 transition-all group"
                                                >
                                                        <div class="text-right">
                                                                <p class="text-[10px] text-[#aaa] dark:text-gray-700 uppercase tracking-wide">Next</p>
                                                                <p class="font-medium text-[#222] dark:text-gray-200">{next.label}</p>
                                                        </div>
                                                        <svg class="w-4 h-4 group-hover:translate-x-0.5 transition-transform" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 12h14M12 5l7 7-7 7"/></svg>
                                                </a>
                                        {/if}
                                </div>
                        {/if}
                </main>
        </div>
</div>
