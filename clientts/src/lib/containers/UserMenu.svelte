<script lang="ts">
	import { goto } from '$app/navigation';
	import Avatar from '$lib/components/Avatar.svelte';
	import * as Dropdown from '$lib/components/ui/dropdown-menu/index.js';
	import { t } from '$lib/i18n/i18n';
	import type { RolesResponse, UsersResponse } from '$lib/models/generated/pb-models';
	import { pb } from '$lib/pocketbase';

	let { user }: { user: UsersResponse<{ role: RolesResponse }> | null } = $props();

	const handleLogoutClick = () => {
		pb.authStore.clear();
		goto('/login');
	};
</script>

<Dropdown.Root>
	<Dropdown.Trigger class="cursor-pointer">
		<Avatar size="sm" url="" name={user?.name} />
	</Dropdown.Trigger>
	<Dropdown.Content>
		<Dropdown.Item onclick={handleLogoutClick}>{$t('userMenu.signOut')}</Dropdown.Item>
	</Dropdown.Content>
</Dropdown.Root>
