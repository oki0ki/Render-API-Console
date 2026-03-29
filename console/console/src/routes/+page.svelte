<script lang="ts">
        import { onMount } from 'svelte';
        import { goto } from '$app/navigation';
        import { supabase } from '$lib/supabase';
        import type { User } from '@supabase/supabase-js';
        import DocsEmbed from '$lib/DocsEmbed.svelte';
        import Sidebar from '$lib/Sidebar.svelte';
        import AppHeader from '$lib/AppHeader.svelte';
        import HomeView from '$lib/HomeView.svelte';
        import KeysView from '$lib/KeysView.svelte';
        let user: User | null = null;
        let loading = true;
        let keys: any[] = [];
        let keysLoading = false;
        let creating = false;
        let newlyCreatedKey: string | null = null;
        let copiedId: string | null = null;
        let activeView: 'home' | 'keys' | 'docs' = 'home';
        let docsPage = 'introduction';
        let showBanner = true;
        let sidebarOpen = true;

        onMount(() => {
                sidebarOpen = window.innerWidth >= 1024;

                let touchstartX = 0;
                let touchendX = 0;
                const checkSwipe = () => {
                        const dist = Math.abs(touchendX - touchstartX);
                        if (dist >= window.innerWidth / 4) {
                                sidebarOpen = touchendX > touchstartX;
                        }
                };
                const onTouchStart = (e: TouchEvent) => { touchstartX = e.changedTouches[0].screenX; };
                const onTouchEnd = (e: TouchEvent) => { touchendX = e.changedTouches[0].screenX; checkSwipe(); };
                document.addEventListener('touchstart', onTouchStart);
                document.addEventListener('touchend', onTouchEnd);
                return () => {
                        document.removeEventListener('touchstart', onTouchStart);
                        document.removeEventListener('touchend', onTouchEnd);
                };
        });

        onMount(async () => {
                const { data: { session } } = await supabase.auth.getSession();
                if (!session) { window.location.replace('/login'); return; }
                user = session.user;
                loading = false;
                await loadKeys();
        });

        async function loadKeys() {
                keysLoading = true;
                try {
                        const { data, error } = await supabase
                                .from('api_keys')
                                .select('id,name,key,active,created_at,requests_count')
                                .eq('user_id', user!.id)
                                .order('created_at', { ascending: false });
                        if (error) console.error('[loadKeys] error:', error);
                        else keys = data ?? [];
                } catch (e) {
                        console.error('[loadKeys] failed:', e);
                } finally {
                        keysLoading = false;
                }
        }

        async function createKey() {
                creating = true;
                try {
                        activeView = 'keys';
                        const { data: existing } = await supabase
                                .from('api_keys')
                                .select('id')
                                .eq('user_id', user!.id)
                                .eq('active', true);
                        if ((existing?.length ?? 0) >= 5) {
                                console.error('Max 5 keys reached');
                                return;
                        }
                        const rawKey = crypto.getRandomValues(new Uint8Array(24));
                        const hex = Array.from(rawKey).map(b => b.toString(16).padStart(2, '0')).join('');
                        const key = 'sk-' + hex;
                        const { error } = await supabase
                                .from('api_keys')
                                .insert({ key, name: 'API Key', user_id: user!.id });
                        if (error) console.error('Key creation failed:', error);
                        else newlyCreatedKey = key;
                        await loadKeys();
                } finally {
                        creating = false;
                }
        }

        async function deleteKey(id: string) {
                const { error } = await supabase
                        .from('api_keys')
                        .delete()
                        .eq('id', id)
                        .eq('user_id', user!.id);
                if (!error) keys = keys.filter(k => k.id !== id);
        }

        async function signOut() {
                await supabase.auth.signOut();
                goto('/login');
        }

        function copy(text: string, id: string) {
                navigator.clipboard.writeText(text);
                copiedId = id;
                setTimeout(() => copiedId = null, 2000);
        }

        function fmtDate(iso: string) {
                return new Date(iso).toLocaleDateString('en-US', { day: 'numeric', month: 'short', year: 'numeric' });
        }

        $: initial = (user?.email ?? 'U').charAt(0).toUpperCase();
        $: totalRequests = keys.reduce((sum: number, k: any) => sum + (k.requests_count ?? 0), 0);
        $: totalCompletions = totalRequests;
</script>

{#if loading}
<div class="min-h-screen flex items-center justify-center bg-white dark:bg-gray-950">
        <div class="w-4 h-4 rounded-full border-[1.5px] border-black/10 dark:border-white/10 border-t-black dark:border-t-white animate-spin"></div>
</div>
{:else}

<div class="flex h-screen w-full bg-[#f9f9f9] dark:bg-gray-950 text-[#0f0f0f] dark:text-gray-100 font-sans overflow-hidden lg:mt-[60px] lg:h-[calc(100vh-60px)]">

        <!-- Mobile overlay -->
        {#if sidebarOpen}
        <div
                class="fixed inset-0 z-[250] bg-black/20 lg:hidden"
                on:click={() => sidebarOpen = false}
                role="button"
                tabindex="-1"
                aria-label="Close sidebar"
        ></div>
        {/if}

        <Sidebar bind:activeView bind:sidebarOpen />

        <!-- Main Content -->
        <div class="flex-1 flex flex-col h-full bg-[#f9f9f9] overflow-x-hidden min-w-0">
                <AppHeader
                        {user}
                        {initial}
                        bind:activeView
                        bind:sidebarOpen
                        onSignOut={signOut}
                />

                <!-- Scrollable Content -->
                <main class="flex-1 min-h-0 bg-white dark:bg-gray-950 rounded-tl-[1.5rem] rounded-tr-[1.5rem] sm:rounded-tr-none sm:rounded-tl-[1.5rem] {activeView === 'docs' ? 'overflow-hidden flex flex-col' : 'overflow-y-auto p-4 sm:p-8'}">
                        <div class="{activeView === 'docs' ? 'flex-1 min-h-0 flex flex-col' : 'max-w-[1200px] mx-auto w-full min-w-0'}">

                                {#if activeView === 'home'}

                                <HomeView
                                        bind:showBanner
                                        {keysLoading}
                                        {totalRequests}
                                        {totalCompletions}
                                        {keys}
                                        onNavigateDocs={() => activeView = 'docs'}
                                        onNavigateKeys={() => activeView = 'keys'}
                                />

                                {:else if activeView === 'keys'}

                                <KeysView
                                        {keys}
                                        {keysLoading}
                                        {creating}
                                        bind:newlyCreatedKey
                                        {copiedId}
                                        onCreateKey={createKey}
                                        onDeleteKey={deleteKey}
                                        onCopy={copy}
                                        {fmtDate}
                                />

                                {:else if activeView === 'docs'}

                                <DocsEmbed bind:activePage={docsPage} />

                                {/if}

                        </div>
                </main>
        </div>
</div>
{/if}
