<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { APP_NAME } from '$lib/constants';
	import dayjs from 'dayjs';
	import { KeyRound, Plus, XCircle } from 'lucide-svelte';

	let { form } = $props();
	let loading = $state(false);

	const minDate = new Date();
</script>

<svelte:head>
	<title>Create API Key | {APP_NAME}</title>
</svelte:head>

<div class="breadcrumbs bg-base-200 border-b-base-300 w-full border-b p-6 text-sm">
	<ul>
		<li>
			<a>
				<KeyRound />
				API Keys Manager
			</a>
		</li>
		<li>
			<a>
				<Plus />
				Create API Key
			</a>
		</li>
	</ul>
</div>
<div class="mt-8 flex w-full flex-col items-center justify-center gap-4">
	<div class="w-64 flex-col items-center gap-4">
		<h1 class="mb-4 text-center text-lg font-bold">Create a new API Key</h1>
		{#if form?.failed}
			<div role="alert" class="alert alert-error my-4">
				<XCircle />
				<span>Failed to create API key.</span>
			</div>
		{/if}
		<form
			class="form-control grid w-full max-w-lg grid-cols-2 gap-4"
			method="POST"
			action="?/create"
			use:enhance={() => {
				loading = true;
				return async ({ update }) => {
					loading = false;
					return update();
				};
			}}
		>
			<label class="label col-span-2">
				<input
					name="name"
					type="text"
					class="input input-bordered w-full"
					placeholder="Name*"
					required
					maxlength="120"
				/>
			</label>
			<button type="submit" class="btn btn-primary col-span-2">
				{#if loading}
					<span class="loading loading-spinner loading-xs"></span>
				{:else}
					<Plus />
				{/if}
				Create
			</button>
			<button type="button" class="btn btn-ghost col-span-2" onclick={() => goto('/apikeys')}
				>Cancel</button
			>
		</form>
	</div>
</div>
