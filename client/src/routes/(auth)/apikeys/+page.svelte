<script>
	import { APP_NAME } from '$lib/constants';
	import { Search, Plus, Check, X, EllipsisVertical, XCircle } from 'lucide-svelte';
	import dayjs from 'dayjs';
	import { getPagination } from '$lib/utils/pagination.js';
	import { replaceStateWithQuery } from '$lib/utils/query-params.js';
	import lo from 'lodash';
	import { pb } from '$lib/utils/pb.js';
	import { catchPromise } from '$lib/utils/catch-promise.js';
	import { goto } from '$app/navigation';

	let { data } = $props();
	let filter = $state(data.query?.filter || '');
	let error = $state('');
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

	/**
	 * @param {any} l
	 */
	const handleToggleActive = (l) => async () => {
		const res = await catchPromise(
			pb.collection('apikey').update(l.id, {
				active: !l.active
			})
		);

		if (res.ok) {
			l.active = !l.active;
			location.reload();
		} else {
			error = 'Failed to toggle active status.';
		}
	};

	const handleCreateClick = () => {
		goto('/apikeys/create');
	};
</script>

<svelte:head>
	<title>API Keys | {APP_NAME}</title>
</svelte:head>

<div class="mt-8">
	{#if error}
		<div role="alert" class="alert alert-error">
			<XCircle />
			<span>{error}</span>
			<div>
				<button class="btn btn-ghost" onclick={() => (error = '')}>Dismiss</button>
			</div>
		</div>
	{/if}
	{#await data.license}
		<div class="flex w-full items-center justify-center">
			<span class="loading loading-dots"></span>
		</div>
	{:then lics}
		{#if lics.items?.length > 0 || filter?.length > 0}
			<div class="flex w-full justify-end gap-4 p-4">
				<label class="input input-bordered flex items-center gap-0.5">
					<Search class="h-5 w-5" />
					<input
						type="text"
						class="grow"
						placeholder="Search..."
						bind:value={filter}
						oninput={handleFilterChange}
					/>
				</label>
				<span class="divider divider-horizontal divide-base-300"></span>
				<button class="btn btn-outline" onclick={handleCreateClick}>
					<Plus class="h-4 w-4" />
					Create API Key
				</button>
			</div>
			<table class="table">
				<!-- head -->
				<thead>
					<tr>
						<th>Name</th>
						<th>Active</th>
						<th>Created</th>
						<th class="w-24"></th>
					</tr>
				</thead>
				<tbody>
					{#each lics.items as l}
						<tr>
							<td>{l.name}</td>
							<td>
								<div class="px-2">
									{#if l.active}
										<Check class="h-4 w-4 text-green-500" />
									{:else}
										<X class="h-4 w-4 text-red-500" />
									{/if}
								</div>
							</td>
							<td>{l.created ? dayjs(l.created).format('DD/MM/YYYY') : '-'}</td>
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
									</ul>
								</div>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>

			{#await pagination then p}
				{#if p.pages > 1}
					<div class="join mt-4 flex justify-center">
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
			<div class="flex justify-center p-4">
				<div class="flex w-56 flex-col gap-4">
					<h1 class="text-center text-lg font-bold">No API Key found.</h1>
					<p class="text-center">Feel free to create as many API keys as required.</p>
					<button class="btn btn-primary" onclick={handleCreateClick}>
						<Plus class="h-4 w-4" />
						Create API Key
					</button>
				</div>
			</div>
		{/if}
	{/await}
</div>
