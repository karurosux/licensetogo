<script>
	import { APP_NAME } from '$lib/constants';
	import * as Card from '$lib/components/ui/card';
	import * as Table from '$lib/components/ui/table';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import { Trash } from 'lucide-svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { onMount } from 'svelte';
	import { breadcrumbs } from '$lib/context/breadcrumbs.js';
	import Input from '$lib/components/ui/input/input.svelte';
	import { Plus } from 'lucide-svelte';
	import Pagination from '$lib/components/Pagination.svelte';

	let { data } = $props();

	onMount(() => {
		breadcrumbs.set([{ href: '/', label: 'Home' }, { label: 'Licenses' }]);
	});
</script>

<svelte:head>
	<title>License Manager | {APP_NAME}</title>
</svelte:head>

<div class="flex gap-4">
	<Input placeholder="Search..." class="flex-1" />
	<Button variant="outline">
		<Plus />
		Create License
	</Button>
</div>

<Card.Root
	data-x-chunk-name="dashboard-05-chunk-3"
	data-x-chunk-description="A table of recent orders showing the following columns: Customer, Type, Status, Date, and Amount."
>
	<Card.Header class="px-7">
		<Card.Title>Licenses</Card.Title>
		<Card.Description>All registered licenses in your application.</Card.Description>
	</Card.Header>
	<Card.Content>
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head>Name</Table.Head>
					<Table.Head class="hidden sm:table-cell">Metadata</Table.Head>
					<Table.Head class="hidden sm:table-cell">Permissions</Table.Head>
					<Table.Head class="hidden md:table-cell">Active</Table.Head>
					<Table.Head>Expires</Table.Head>
					<Table.Head>Created</Table.Head>
					<Table.Head></Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each data?.licenses?.items || [] as l (l.id)}
					<Table.Row>
						<Table.Cell>
							<div class="font-medium">{l.name}</div>
						</Table.Cell>
						<Table.Cell class="hidden sm:table-cell">
							<div class="text-muted-foreground hidden text-sm md:inline">
								{JSON.stringify(l.metadata || {})}
							</div>
						</Table.Cell>
						<Table.Cell class="hidden md:table-cell">
							{JSON.stringify(l.permissions || [])}
						</Table.Cell>
						<Table.Cell class="hidden sm:table-cell">
							<Badge class="text-xs text-white {l.active ? 'bg-green-500' : ' bg-red-500'}">
								{l.active ? 'Active' : 'Inactive'}
							</Badge>
						</Table.Cell>
						<Table.Cell
							>{l.expires
								? new Date(l.expires).toLocaleDateString(undefined, {
										dateStyle: 'medium'
									})
								: '---'}</Table.Cell
						>
						<Table.Cell
							>{l.created
								? new Date(l.created).toLocaleDateString(undefined, {
										dateStyle: 'medium'
									})
								: '---'}</Table.Cell
						>
						<Table.Cell>
							<Button variant="outline" class="btn btn-sm btn-square btn-ghost" on:click={() => {}}>
								<Trash class="h-4 w-4" />
							</Button>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</Card.Content>
	<Card.Footer>
		<Pagination count={data.licenses?.totalItems || 0} perPage={data.licenses?.perPage || 5} />
	</Card.Footer>
</Card.Root>
