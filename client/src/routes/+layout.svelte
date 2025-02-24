<script>
	import { browser } from '$app/environment';
	import { setUserContext } from '$lib/context/user';
	import { pb } from '$lib/utils/pb';
	import { onDestroy } from 'svelte';
	import { writable } from 'svelte/store';
	import '../app.css';

	let { data, children } = $props();

	// Initialize user store
	const user = writable(data.user);
	setUserContext(user);

	if (browser) {
		// Load user from cookie (client-side only)
		pb.authStore.loadFromCookie(document.cookie);

		// Update user store when auth store changes
		const unsubscribe = pb.authStore.onChange(() => {
			user.set(pb.authStore.record);
			document.cookie = pb.authStore.exportToCookie({ httpOnly: false });
		}, true);

		onDestroy(unsubscribe);
	}
</script>

<div class="bg-base-100 h-screen w-screen overflow-x-auto">
	{@render children()}
</div>
