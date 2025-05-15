<script lang="ts">
	import { page } from '$app/state';
	import { t } from '$lib/i18n/i18n';
	import type { RolesResponse, UsersResponse } from '$lib/models/generated/pb-models';
	import lo from 'lodash';
	import type { Snippet } from 'svelte';

	type MenuItem = {
		route: string;
		label: string;
	};

	type Props = { user: UsersResponse<{ role: RolesResponse }>; children: Snippet<[MenuItem]> };

	let { user, children }: Props = $props();
	let items = $state<MenuItem[]>([]);

	$effect(() => {
		const i: MenuItem[] = [];
		user.expand?.role?.permission?.forEach((permission) => {
			switch (permission) {
				case 'subject_write':
				case 'subject_read':
					i.push({
						route: '/subjects',
						label: 'general.subjects'
					});
					break;
				case 'product_read':
				case 'product_write':
					i.push({
						route: '/products',
						label: 'general.products'
					});
					break;
				case 'user_read':
				case 'user_write':
					i.push({
						route: '/users',
						label: 'general.users'
					});
					break;
			}
		});

		items = lo.uniqBy(i, 'route');
	});
</script>

{#each items || [] as i (i.route)}
	{@render children?.(i)}
{/each}
