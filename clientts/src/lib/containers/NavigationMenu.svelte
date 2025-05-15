<script lang="ts">
	import { page } from '$app/state';
	import { t } from '$lib/i18n/i18n';
	import type { RolesResponse, UsersResponse } from '$lib/models/generated/pb-models';
	import NavigationMenuItems from './NavigationMenuItems.svelte';

	let { user }: { user: UsersResponse<{ role: RolesResponse }> } = $props();
</script>

<div class="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8">
	<NavigationMenuItems {user}>
		{#snippet children(i)}
			<a
				href={i.route || '#'}
				class="hover:text-primary hover:border-primary inline-flex items-center px-1 pt-1 text-sm font-medium hover:border-b-2"
				class:border-b-2={page.url?.pathname.includes(i.route)}
				class:border-primary={page.url?.pathname.includes(i.route)}
				aria-current="page">{$t(i.label)}</a
			>
		{/snippet}
	</NavigationMenuItems>
</div>
