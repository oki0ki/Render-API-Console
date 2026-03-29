<script lang="ts">
        import { ChevronDown, Settings, LogOut, Sun, Moon } from 'lucide-svelte';
        import type { User } from '@supabase/supabase-js';
        import { theme } from './theme';

        export let user: User | null;
        export let initial: string;
        export let activeView: 'home' | 'keys' | 'docs';
        export let sidebarOpen: boolean;
        export let onSignOut: () => void;

        let showUserMenu = false;
</script>

<header class="h-[60px] flex-shrink-0 sticky top-0 z-[200] flex items-center justify-between px-4
        bg-[#f9f9f9] dark:bg-gray-900
        lg:fixed lg:top-0 lg:left-0 lg:right-0 lg:z-[400]">
        <div class="flex items-center gap-2 min-w-0 flex-1">
                <!-- Mobile sidebar toggle -->
                <button
                        on:click={() => sidebarOpen = !sidebarOpen}
                        class="lg:hidden flex-shrink-0 flex items-center justify-center w-9 h-9 rounded-lg text-[#666666] dark:text-gray-500 hover:bg-[#e5e5e5]/60 dark:hover:bg-gray-800/60 transition-colors"
                >
                        <svg class="w-7 h-7 flex-shrink-0" viewBox="0 0 24 24" fill="currentColor"><path d="M3 8C3 7.44772 3.44772 7 4 7H20C20.5523 7 21 7.44772 21 8C21 8.55228 20.5523 9 20 9H4C3.44772 9 3 8.55228 3 8ZM3 16C3 15.4477 3.44772 15 4 15H14C14.5523 15 15 15.4477 15 16C15 16.5523 14.5523 17 14 17H4C3.44772 17 3 16.5523 3 16Z" fill="currentColor" /></svg>
                </button>
                <div class="hidden sm:flex items-center gap-1.5 cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-800 px-2 py-1 rounded-md transition-colors">
                        <div class="w-5 h-5 bg-[#202124] dark:bg-gray-200 text-white dark:text-gray-900 rounded flex items-center justify-center text-[11px] font-medium">
                                {initial}
                        </div>
                        <span class="text-sm font-medium">api</span>
                        <ChevronDown class="w-4 h-4 text-[#888888] dark:text-gray-600" strokeWidth={2} />
                </div>
                <span class="hidden sm:block text-[#d2d2d2] dark:text-gray-800">/</span>
                <div class="hidden sm:flex items-center gap-1.5 cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-800 px-2 py-1 rounded-md transition-colors">
                        <span class="text-sm font-medium">Default project</span>
                        <ChevronDown class="w-4 h-4 text-[#888888] dark:text-gray-600" strokeWidth={2} />
                </div>
        </div>
        <div class="flex items-center gap-3 sm:gap-6">
                <nav class="hidden sm:flex items-center gap-6">
                        <button on:click={() => activeView = 'home'} class="text-sm font-medium {activeView === 'home' ? 'text-[#0f0f0f] dark:text-gray-50' : 'text-[#666666] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200'} transition-colors">Dashboard</button>
                        <button on:click={() => activeView = 'docs'} class="text-sm {activeView === 'docs' ? 'font-medium text-[#0f0f0f] dark:text-gray-50' : 'text-[#666666] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200'} transition-colors">API Docs</button>
                </nav>
                <div class="flex items-center gap-2 sm:gap-3 sm:ml-2">
                        <!-- Dark mode toggle -->
                        <button
                                on:click={theme.toggle}
                                class="flex items-center justify-center w-8 h-8 rounded-lg text-[#666666] dark:text-gray-500 hover:bg-[#e5e5e5]/60 dark:hover:bg-gray-800/60 transition-colors"
                                title="Toggle dark mode"
                        >
                                {#if $theme}
                                        <Sun class="w-4 h-4" strokeWidth={2} />
                                {:else}
                                        <Moon class="w-4 h-4" strokeWidth={2} />
                                {/if}
                        </button>
                        <button class="text-[#666666] dark:text-gray-500 hover:text-[#0f0f0f] dark:hover:text-gray-200 transition-colors">
                                <Settings class="w-5 h-5" strokeWidth={2} />
                        </button>
                        <div class="relative">
                                <button
                                        on:click={() => showUserMenu = !showUserMenu}
                                        class="w-7 h-7 rounded-full border border-[#e5e5e5] dark:border-gray-800 flex items-center justify-center text-xs font-medium text-[#444444] dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
                                >
                                        {initial}
                                </button>
                                {#if showUserMenu}
                                <div
                                        role="menu"
                                        tabindex="-1"
                                        class="absolute right-0 top-9 w-56 bg-white dark:bg-gray-900 border border-[#e5e5e5] dark:border-gray-850 rounded-xl shadow-lg dark:shadow-black/40 z-50 overflow-hidden"
                                        on:mouseleave={() => showUserMenu = false}
                                >
                                        <div class="px-4 py-3 border-b border-[#f0f0f0] dark:border-gray-850">
                                                <p class="text-xs text-[#888888] dark:text-gray-600">Zalogowany jako</p>
                                                <p class="text-sm font-medium text-[#0f0f0f] dark:text-gray-100 truncate">{user?.email ?? ''}</p>
                                        </div>
                                        <button
                                                on:click={onSignOut}
                                                class="w-full flex items-center gap-3 px-4 py-3 text-sm text-[#444444] dark:text-gray-400 hover:bg-[#f9f9f9] dark:hover:bg-gray-800 transition-colors"
                                        >
                                                <LogOut class="w-4 h-4" strokeWidth={2} />
                                                Wyloguj
                                        </button>
                                </div>
                                {/if}
                        </div>
                </div>
        </div>
</header>
