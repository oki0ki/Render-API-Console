<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { supabase } from '$lib/supabase';

	let loaded = false;
	let view: 'login' | 'signup' = 'login';

	let loginStep: 'email' | 'password' = 'email';
	let loginEmail = '';
	let loginPassword = '';
	let loginLoading = false;
	let loginError = '';

	let signupStep: 'email' | 'password' = 'email';
	let signupEmail = '';
	let signupPassword = '';
	let signupLoading = false;
	let signupError = '';
	let signupSuccess = '';

	onMount(async () => {
		const { data: { session } } = await supabase.auth.getSession();
		if (session) { goto('/'); return; }
		loaded = true;
	});

	async function handleLoginEmailSubmit() {
		if (!loginEmail) return;
		loginError = '';
		loginLoading = true;
		await new Promise(r => setTimeout(r, 200));
		loginLoading = false;
		loginStep = 'password';
	}

	async function handleLoginPasswordSubmit() {
		loginLoading = true;
		loginError = '';
		try {
			const { error } = await supabase.auth.signInWithPassword({ email: loginEmail, password: loginPassword });
			if (error) { loginError = error.message; loginLoading = false; }
			else goto('/');
		} catch {
			loginError = 'Błąd logowania.';
			loginLoading = false;
		}
	}

	async function handleSignupEmailSubmit() {
		if (!signupEmail) return;
		signupError = '';
		signupLoading = true;
		await new Promise(r => setTimeout(r, 200));
		signupLoading = false;
		signupStep = 'password';
	}

	async function handleSignupPasswordSubmit() {
		signupLoading = true;
		signupError = '';
		try {
			const { error } = await supabase.auth.signUp({ email: signupEmail, password: signupPassword });
			if (error) { signupError = error.message; signupLoading = false; }
			else {
				signupSuccess = 'Konto zostało utworzone. Możesz się teraz zalogować.';
				switchToLogin();
			}
		} catch {
			signupError = 'Błąd rejestracji.';
			signupLoading = false;
		}
	}

	function switchToSignup() {
		view = 'signup';
		signupStep = 'email';
		signupEmail = '';
		signupPassword = '';
		signupError = '';
		signupSuccess = '';
	}

	function switchToLogin() {
		view = 'login';
		loginStep = 'email';
		loginEmail = '';
		loginPassword = '';
		loginError = '';
	}

	$: loginTitle = loginStep === 'email' ? 'Zaloguj się' : 'Wprowadź hasło';
	$: signupTitle = signupStep === 'email' ? 'Utwórz konto' : 'Ustaw hasło';
</script>

<div class="w-full h-screen max-h-[100dvh] relative">
	<div class="w-full h-full absolute top-0 left-0 bg-white dark:bg-black"></div>

	{#if loaded}
		<div class="fixed min-h-screen w-full flex justify-center z-50 text-black dark:text-white">
			<div class="w-full px-6 sm:px-10 min-h-screen flex flex-col text-center">
				<div class="mt-[32vh] mb-auto flex flex-col justify-center items-center">
					<div class="sm:max-w-md my-auto pb-10 w-full dark:text-gray-100">

						{#if view === 'login'}
							<form
								on:submit|preventDefault={() => loginStep === 'email' ? handleLoginEmailSubmit() : handleLoginPasswordSubmit()}
								class="flex flex-col justify-center"
							>
								<div class="mb-1">
									<div class="text-2xl font-medium">{loginTitle}</div>
								</div>

								<div class="flex flex-col mt-4">
									{#if loginStep === 'email'}
										<div class="mb-2">
											<input
												bind:value={loginEmail}
												type="email"
												autocomplete="email"
												placeholder="Adres e-mail"
												required
												class="my-0.5 w-full text-sm outline-none bg-transparent border border-gray-200 dark:border-gray-700 rounded-full px-4 py-2.5 placeholder:text-gray-400 dark:placeholder:text-gray-500 focus:ring-0 focus:outline-none focus:border-gray-400 dark:focus:border-gray-500 transition"
											/>
										</div>
									{:else}
										<input type="email" autocomplete="email" value={loginEmail} aria-hidden="true" tabindex="-1" style="position:absolute;width:1px;height:1px;opacity:0;pointer-events:none;" readonly />
										<div class="flex items-center gap-2 mb-3">
											<button
												type="button"
												aria-label="Wróć"
												on:click={() => { loginStep = 'email'; loginError = ''; }}
												class="h-7 w-7 flex items-center justify-center rounded-full hover:bg-gray-100 dark:hover:bg-gray-800 transition"
											>
												<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 dark:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m12 19-7-7 7-7"/><path d="M19 12H5"/></svg>
											</button>
											<span class="text-sm text-gray-500 dark:text-gray-400">{loginEmail}</span>
										</div>
										<div class="mb-2">
											<input
												bind:value={loginPassword}
												type="password"
												autocomplete="current-password"
												placeholder="Hasło"
												required
												class="my-0.5 w-full text-sm outline-none bg-transparent border border-gray-200 dark:border-gray-700 rounded-full px-4 py-2.5 placeholder:text-gray-400 dark:placeholder:text-gray-500 focus:ring-0 focus:outline-none focus:border-gray-400 dark:focus:border-gray-500 transition"
											/>
										</div>
									{/if}
								</div>

								{#if loginError}
									<p class="text-sm text-red-500 mb-2">{loginError}</p>
								{/if}

								{#if signupSuccess}
									<p class="text-sm text-green-600 mb-2">{signupSuccess}</p>
								{/if}

								<div class="mt-2">
									<button
										type="submit"
										disabled={loginLoading}
										class="bg-gray-700/5 hover:bg-gray-700/10 dark:bg-gray-100/5 dark:hover:bg-gray-100/10 dark:text-gray-300 dark:hover:text-white transition w-full rounded-full font-medium text-sm py-2.5 disabled:opacity-60 flex items-center justify-center"
									>
										{#if loginLoading}
											<svg class="h-4 w-4 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path></svg>
										{:else}
											{loginStep === 'email' ? 'Kontynuuj' : 'Zaloguj się'}
										{/if}
									</button>

									<div class="mt-4 text-sm text-center">
										Nie masz konta?
										<button type="button" on:click={switchToSignup} class="font-medium underline">
											Zarejestruj się
										</button>
									</div>
								</div>
							</form>

						{:else}
							<form
								on:submit|preventDefault={() => signupStep === 'email' ? handleSignupEmailSubmit() : handleSignupPasswordSubmit()}
								class="flex flex-col justify-center"
							>
								<div class="mb-1">
									<div class="text-2xl font-medium">{signupTitle}</div>
								</div>

								<div class="flex flex-col mt-4">
									{#if signupStep === 'email'}
										<div class="mb-2">
											<input
												bind:value={signupEmail}
												type="email"
												autocomplete="email"
												placeholder="Adres e-mail"
												required
												class="my-0.5 w-full text-sm outline-none bg-transparent border border-gray-200 dark:border-gray-700 rounded-full px-4 py-2.5 placeholder:text-gray-400 dark:placeholder:text-gray-500 focus:ring-0 focus:outline-none focus:border-gray-400 dark:focus:border-gray-500 transition"
											/>
										</div>
									{:else}
										<input type="email" autocomplete="email" value={signupEmail} aria-hidden="true" tabindex="-1" style="position:absolute;width:1px;height:1px;opacity:0;pointer-events:none;" readonly />
										<div class="flex items-center gap-2 mb-3">
											<button
												type="button"
												aria-label="Wróć"
												on:click={() => { signupStep = 'email'; signupError = ''; }}
												class="h-7 w-7 flex items-center justify-center rounded-full hover:bg-gray-100 dark:hover:bg-gray-800 transition"
											>
												<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 dark:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m12 19-7-7 7-7"/><path d="M19 12H5"/></svg>
											</button>
											<span class="text-sm text-gray-500 dark:text-gray-400">{signupEmail}</span>
										</div>
										<div class="mb-2">
											<input
												bind:value={signupPassword}
												type="password"
												autocomplete="new-password"
												placeholder="Hasło"
												required
												class="my-0.5 w-full text-sm outline-none bg-transparent border border-gray-200 dark:border-gray-700 rounded-full px-4 py-2.5 placeholder:text-gray-400 dark:placeholder:text-gray-500 focus:ring-0 focus:outline-none focus:border-gray-400 dark:focus:border-gray-500 transition"
											/>
										</div>
									{/if}
								</div>

								{#if signupError}
									<p class="text-sm text-red-500 mb-2">{signupError}</p>
								{/if}

								<div class="mt-2">
									<button
										type="submit"
										disabled={signupLoading}
										class="bg-gray-700/5 hover:bg-gray-700/10 dark:bg-gray-100/5 dark:hover:bg-gray-100/10 dark:text-gray-300 dark:hover:text-white transition w-full rounded-full font-medium text-sm py-2.5 disabled:opacity-60 flex items-center justify-center"
									>
										{#if signupLoading}
											<svg class="h-4 w-4 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path></svg>
										{:else}
											{signupStep === 'email' ? 'Kontynuuj' : 'Zarejestruj się'}
										{/if}
									</button>

									<div class="mt-4 text-sm text-center">
										Masz już konto?
										<button type="button" on:click={switchToLogin} class="font-medium underline">
											Zaloguj się
										</button>
									</div>
								</div>
							</form>
						{/if}

					</div>
				</div>
			</div>
		</div>
	{/if}
</div>
