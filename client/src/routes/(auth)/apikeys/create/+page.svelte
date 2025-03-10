<script>
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import Breadcrumbs from '$lib/components/breadcrumbs/Breadcrumbs.svelte';
	import { APP_NAME } from '$lib/constants';
	import { KeyRound, Plus, XCircle } from 'lucide-svelte';

	let { form } = $props();
	let loading = $state(false);
</script>

<svelte:head>
	<title>Create API Key | {APP_NAME}</title>
</svelte:head>

<Breadcrumbs
	items={[
		{ label: 'API Keys Manager', icon: KeyRound },
		{ label: 'Create API Key', icon: Plus }
	]}
></Breadcrumbs>

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
