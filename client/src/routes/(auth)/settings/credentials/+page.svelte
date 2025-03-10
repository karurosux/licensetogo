<script>
	import { enhance } from '$app/forms';
	import Breadcrumbs from '$lib/components/breadcrumbs/Breadcrumbs.svelte';
	import { APP_NAME } from '$lib/constants';
	import { CheckCircle, Cog, Info, Save, User, XCircle } from 'lucide-svelte';

	let { form } = $props();
	let loading = $state(false);
</script>

<svelte:head>
	<title>Update Credentials | {APP_NAME}</title>
</svelte:head>

<Breadcrumbs
	items={[
		{ label: 'Settings', icon: Cog },
		{ label: 'Application Credentials', icon: User }
	]}
></Breadcrumbs>

<div class="mt-8 flex w-full flex-col items-center justify-center gap-4">
	<div class="w-64 flex-col items-center gap-4">
		<h1 class="mb-4 text-center text-lg font-bold">Update Credentials</h1>

		<div role="alert" class="alert alert-info my-4">
			<Info />
			<span>You will be logged out after updating credentials.</span>
		</div>

		{#if form?.success}
			<div role="alert" class="alert alert-success my-4">
				<CheckCircle />
				<span>Updated credentials successfully.</span>
			</div>
		{/if}
		{#if form?.error}
			<div role="alert" class="alert alert-error my-4">
				<XCircle />
				<span>Failed to update credentials.</span>
			</div>
		{/if}
		<form
			class="form-control grid w-full max-w-lg grid-cols-2 gap-4"
			method="POST"
			action="?/update"
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
					name="email"
					type="email"
					class="input input-bordered invalid:input-error w-full"
					placeholder="Email*"
					required
				/>
			</label>
			<label class="label col-span-2">
				<input
					name="oldPassword"
					type="password"
					class="input input-bordered invalid:input-error w-full"
					placeholder="Password*"
					required
				/>
			</label>
			<label class="label col-span-2">
				<input
					name="password"
					type="password"
					class="input input-bordered invalid:input-error w-full"
					placeholder="New Password*"
					required
				/>
			</label>
			<button type="submit" class="btn btn-primary col-span-2">
				{#if loading}
					<span class="loading loading-spinner loading-xs"></span>
				{:else}
					<Save />
				{/if}
				Update
			</button>
		</form>
	</div>
</div>
