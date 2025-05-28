<script lang="ts">
	import { browser } from '$app/environment';
	import { currentUser, setUserContext } from '$lib/context/user';
	import type { RolesResponse, UsersResponse } from '$lib/models/generated/pb-models';
	import { pb } from '$lib/pocketbase';
	import type { Writable } from 'svelte/store';
	import '../app.css';
	import { onDestroy } from 'svelte';

	let { data, children } = $props();

	// Initialize user store
	currentUser.set(data.user);
	setUserContext(currentUser as Writable<UsersResponse<{ role: RolesResponse }>>);

	if (browser) {
		// Load user from cookie (client-side only)
		pb.authStore.loadFromCookie(document.cookie);

		// Update user store when auth store changes
		const unsubscribe = pb.authStore.onChange(() => {
			currentUser.set(pb.authStore.record as UsersResponse<{ role: RolesResponse }>);
			document.cookie = pb.authStore.exportToCookie({ httpOnly: false });
		}, true);

		onDestroy(unsubscribe);
	}
</script>

{@render children()}
