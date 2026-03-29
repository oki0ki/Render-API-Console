import { writable } from 'svelte/store';
import { browser } from '$app/environment';

function createTheme() {
	const { subscribe, set, update } = writable(false);

	return {
		subscribe,
		init() {
			if (!browser) return;
			const saved = localStorage.getItem('theme');
			const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
			const dark = saved === 'dark' || (!saved && prefersDark);
			set(dark);
			document.documentElement.classList.toggle('dark', dark);
		},
		toggle() {
			if (!browser) return;
			const isDark = document.documentElement.classList.toggle('dark');
			localStorage.setItem('theme', isDark ? 'dark' : 'light');
			set(isDark);
		}
	};
}

export const theme = createTheme();
