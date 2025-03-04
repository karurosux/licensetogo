<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { APP_NAME } from '$lib/constants';
	import { catchPromise } from '$lib/utils/catch-promise.js';
	import { getPagination } from '$lib/utils/pagination.js';
	import { replaceStateWithQuery } from '$lib/utils/query-params.js';
	import dayjs from 'dayjs';
	import lo from 'lodash';
	import { Check, EllipsisVertical, Plus, Search, X, XCircle, Scroll, Trash } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let { data } = $props();
	let filter = $state(data.query?.filter || '');
	let error = $state('');
	let deleting = $state(false);
	let showCreateNew = $state(false);
	let deleteLicense = $state(null);
	let setActiveLicense = $state(null);
	let pagination = $derived(
		data.license?.then((lic) => getPagination(lic.totalItems, lic.perPage, lic.page - 1))
	);

	const handleFilterChange = lo.debounce(() => {
		replaceStateWithQuery({ filter, offset: 0 });
	}, 300);

	/**
	 * @param {number} page
	 */
	const handlePageChange = (page) => () => {
		replaceStateWithQuery({ filter, offset: page - 1 });
	};

	const handleCreateClick = () => {
		goto('/license-manager/create');
	};

	const handleDeleteClick = (l) => () => {
		deleteLicense = l;
	};

	/**
	 * @param {any} l
	 */
	const handleToggleActive = (l) => () => {
		setActiveLicense = l;
	};

	onMount(() => {
		data.license.then((lic) => (showCreateNew = lic.totalItems === 0));
	});
</script>

<svelte:head>
	<title>License Manager | {APP_NAME}</title>
</svelte:head>

<div class="breadcrumbs bg-base-200 w-full p-6 text-sm">
	<ul>
		<li>
			<a>
				<Scroll />
				License Manager
			</a>
		</li>
	</ul>
</div>

<div>
	{#if error}
		<div role="alert" class="alert alert-error rounded-none">
			<XCircle />
			<span>{error}</span>
			<div>
				<button class="btn btn-ghost" onclick={() => (error = '')}>Dismiss</button>
			</div>
		</div>
	{/if}

	<div class="border-y-base-300 bg-base-200 flex w-full flex-row rounded-none border-y p-8">
		<label class="input input-md bg-base-100 border-base-300 flex flex-1 items-center gap-0.5">
			<Search class="h-5 w-5" />
			<input
				type="text"
				class="grow"
				placeholder="Search by License Name"
				bind:value={filter}
				oninput={handleFilterChange}
			/>
		</label>
		<span class="divider divider-horizontal divide-base-300"></span>
		<button class="btn btn-md btn-primary" onclick={handleCreateClick}>
			<Plus class="h-4 w-4" />
			Create License
		</button>
	</div>
	{#await data.license}
		<div class="flex w-full items-center justify-center">
			<span class="loyding loading-dots"></span>
		</div>
	{:then lics}
		{#if lics.items?.length > 0}
			<table class="bg-base-200 table-md table rounded-none">
				<!-- head -->
				<thead class="bg-base-300">
					<tr>
						<th>Name</th>
						<th>Permissions</th>
						<th>Metadata</th>
						<th>Active</th>
						<th>Expires At</th>
						<th>Last Used At</th>
						<th class="w-24"></th>
					</tr>
				</thead>
				<tbody>
					{#each lics.items as l}
						<tr>
							<td>{l.name}</td>
							<td>{l.permissions?.toString() || 'N/A'}</td>
							<td>{l.metadata ? JSON.stringify(l.metadata) : 'N/A'}</td>
							<td>
								<div class="px-2">
									{#if l.active}
										<Check class="h-4 w-4 text-green-500" />
									{:else}
										<X class="h-4 w-4 text-red-500" />
									{/if}
								</div>
							</td>
							<td>{l.expires ? dayjs(l.expires).format('DD/MM/YYYY') : '-'}</td>
							<td>{l.lastused ? dayjs(l.lastused).format('DD/MM/YYYY') : '-'}</td>
							<td>
								<div class="dropdown dropdown-end">
									<div tabindex="0" role="button" class="btn m-1">
										<EllipsisVertical class="h-4 w-4" />
									</div>
									<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
									<ul
										tabindex="0"
										class="dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow"
									>
										{#if l.active}
											<li><button onclick={handleToggleActive(l)}><X /> Disable</button></li>
										{:else}
											<li><button onclick={handleToggleActive(l)}><Check /> Enable</button></li>
										{/if}
										<li><button onclick={handleDeleteClick(l)}><Trash /> Delete</button></li>
									</ul>
								</div>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>

			{#await pagination then p}
				{#if p.pages > 1}
					<div class="join border-y-base-300 flex justify-center border-y p-8">
						{#if p.prev}
							<button class="join-item btn" onclick={handlePageChange(p.prev)}>Prev</button>
						{/if}
						{#each p.visiblePages as page}
							{#if page !== null}
								<button
									class="join-item btn"
									class:btn-active={p.current === page}
									disabled={p.current === page}
									onclick={handlePageChange(page)}>{page}</button
								>
							{/if}
						{/each}
						{#if p.next}
							<button class="join-item btn" onclick={handlePageChange(p.next)}>Next</button>
						{/if}
					</div>
				{/if}
			{/await}
		{:else}
			<div class="mt-8 flex justify-center p-4">
				<div class="flex w-56 flex-col gap-4">
					<h1 class="text-base-300 text-neutral flex items-center justify-center gap-2 text-center">
						<XCircle />
						No License Found
					</h1>
				</div>
			</div>
		{/if}
	{/await}
</div>

<dialog id="deleteDialog" class="modal" open={!!deleteLicense}>
	<div class="modal-box">
		<h3 class="text-lg font-bold">Delete License</h3>
		<p class="py-4">Are you sure you want to delete <b>"{deleteLicense?.name}"</b> License?</p>
		<div class="modal-action">
			<form
				method="POST"
				action="?/delete"
				use:enhance={() => {
					deleting = true;
					return async ({ update }) => {
						deleting = false;
						deleteLicense = null;
						return update();
					};
				}}
			>
				<input type="hidden" name="id" value={deleteLicense?.id} />
				<button type="button" class="btn" onclick={() => (deleteLicense = null)}>Close</button>
				<button type="submit" class="btn btn-primary">
					{#if deleting}
						<span class="loading loading-spinner loading-xs"></span>
					{/if}
					Delete
				</button>
			</form>
		</div>
	</div>
</dialog>

<dialog id="setActiveDialog" class="modal" open={!!setActiveLicense}>
	<div class="modal-box">
		<h3 class="text-lg font-bold">{setActiveLicense?.active ? 'Disable' : 'Enable'} License</h3>
		<p class="py-4">
			Are you sure you want to {setActiveLicense?.active ? 'disable' : 'enable'}
			<b>"{setActiveLicense?.name}"</b> License?
		</p>
		<div class="modal-action">
			<form
				method="POST"
				action="?/setActive"
				use:enhance={() => {
					return async ({ update }) => {
						setActiveLicense = null;
						return update();
					};
				}}
			>
				<input type="hidden" name="id" value={setActiveLicense?.id} />
				<input type="hidden" name="value" value={!setActiveLicense?.active} />
				<button type="button" class="btn" onclick={() => (setActiveLicense = null)}>Close</button>
				<button type="submit" class="btn btn-primary">
					{#if deleting}
						<span class="loading loading-spinner loading-xs"></span>
					{/if}
					{setActiveLicense?.active ? 'Disable' : 'Enable'}
				</button>
			</form>
		</div>
	</div>
</dialog>
