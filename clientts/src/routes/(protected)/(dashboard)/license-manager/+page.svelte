<script lang="ts">
	import { applyAction, enhance } from '$app/forms'
	import Pagination from '$lib/components/Pagination.svelte'
	import Badge from '$lib/components/ui/badge/badge.svelte'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Card from '$lib/components/ui/card'
	import Input from '$lib/components/ui/input/input.svelte'
	import * as Table from '$lib/components/ui/table'
	import { APP_NAME } from '$lib/constants'
	import CreateLicense from '$lib/containers/subjects/CreateLicense.svelte'
	import DeleteLicense from '$lib/containers/subjects/DeleteLicense.svelte'
	import { breadcrumbs } from '$lib/context/breadcrumbs.js'
	import { t } from '$lib/i18n/i18n.js'
	import { getQueryParams, mergeStateWithQuery } from '$lib/utils/query-params'
	import { ToggleLeft, ToggleRight } from '@lucide/svelte'
	import debounce from 'lodash/debounce'
	import { onMount } from 'svelte'
	import type { FormEventHandler } from 'svelte/elements'

	let { data } = $props()
	let filter = $state('')

	const handleFilterChange: FormEventHandler<HTMLInputElement> = debounce(() => {
		mergeStateWithQuery({ filter })
	}, 300)

	onMount(() => {
		const queryParams = getQueryParams()
		filter = queryParams.filter
		breadcrumbs.set([{ href: '/', label: 'Home' }, { label: 'Licenses' }])
	})
</script>

<svelte:head>
	<title>{$t((ts) => ts.licenseManager.createLicense.title)} | {APP_NAME}</title>
</svelte:head>

<div class="flex gap-4">
	<Input
		placeholder={$t((ts) => ts.general.search)}
		class="flex-1"
		bind:value={filter}
		autofocus={!!filter}
		oninput={handleFilterChange} />
	<CreateLicense />
</div>

<Card.Root>
	<Card.Header class="px-7">
		<Card.Title>{$t((ts) => ts.general.licenses)}</Card.Title>
		<Card.Description>{$t((ts) => ts.licenseManager.descriptions)}</Card.Description>
	</Card.Header>
	<Card.Content>
		{#if (data?.licenses?.totalItems || 0) > 0}
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head>{$t((ts) => ts.general.name)}</Table.Head>
						<Table.Head class="hidden sm:table-cell">{$t((ts) => ts.general.metadata)}</Table.Head>
						<Table.Head class="hidden sm:table-cell">
							{$t((ts) => ts.general.permissions)}
						</Table.Head>
						<Table.Head class="hidden md:table-cell">{$t((ts) => ts.general.active)}</Table.Head>
						<Table.Head>{$t((ts) => ts.general.expires)}</Table.Head>
						<Table.Head>{$t((ts) => ts.general.createdAt)}</Table.Head>
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
							<Table.Cell>
								{l.expires
									? new Date(l.expires).toLocaleDateString(undefined, {
											dateStyle: 'medium',
										})
									: '---'}
							</Table.Cell>
							<Table.Cell>
								{l.created
									? new Date(l.created).toLocaleDateString(undefined, {
											dateStyle: 'medium',
										})
									: '---'}
							</Table.Cell>
							<Table.Cell class="flex items-center justify-center gap-2">
								<span class="toggle-active">
									<form
										method="POST"
										action="?/active"
										use:enhance={() => {
											return async ({ result, update }) => {
												applyAction(result)
												return update()
											}
										}}>
										<input type="hidden" name="id" value={l.id} />
										<input type="hidden" name="value" value={!l.active} />
										<Button type="submit" variant="outline" class="btn btn-sm btn-square btn-ghost">
											{#if l.active}
												<ToggleRight class="h-4 w-4 text-green-600" />
											{:else}
												<ToggleLeft class="h-4 w-4" />
											{/if}
										</Button>
									</form>
								</span>
								<DeleteLicense id={l.id} name={l.name} />
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		{:else}
			<span class="p-2">{$t((ts) => ts.general.noContent)}</span>
		{/if}
	</Card.Content>
	<Card.Footer>
		<Pagination count={data.licenses?.totalItems || 0} perPage={data.licenses?.perPage || 5} />
	</Card.Footer>
</Card.Root>
