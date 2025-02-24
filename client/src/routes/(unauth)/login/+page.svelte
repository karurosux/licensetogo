<script>
	import { applyAction, enhance } from '$app/forms';
	import { APP_NAME } from '$lib/constants';
	import { pb } from '$lib/utils/pb.js';
	import { Lock, LogIn, Mail, ShieldHalf } from 'lucide-svelte';

	let loading = $state(false);
	let { form } = $props();
</script>

<svelte:head>
	<title>Login | {APP_NAME}</title>
</svelte:head>
<div class="flex h-screen w-screen flex-col items-center justify-center gap-4 p-4">
	{#if form?.failed}
		<div class="alert alert-error w-80" role="alert">
			<strong>Error!</strong> Invalid email or password.
		</div>
	{/if}
	<div class="card w-80 border-2 border-gray-600 shadow-2xl">
		<div class="card-body">
			<div class="mb-10 flex items-center justify-center gap-1 text-xl">
				<ShieldHalf class="h-6 w-6" />
				LicenseToGo
			</div>
			<form
				class="form-horizontal flex flex-col gap-4"
				method="POST"
				action="?/login"
				use:enhance={() => {
					loading = true;
					return async ({ update, result }) => {
						loading = false;
						pb.authStore.loadFromCookie(document.cookie);
						await applyAction(result);
						update();
					};
				}}
			>
				<label class="input input-bordered invalid:input-error flex items-center gap-2">
					<Mail class="h-5 w-5" />
					<input
						type="email"
						class="grow"
						id="email"
						name="email"
						placeholder="Email"
						required
						maxlength="60"
					/>
				</label>
				<label class="input input-bordered invalid:input-error flex items-center gap-2">
					<Lock class="h-5 w-5" />
					<input
						type="password"
						class="grow"
						id="password"
						name="password"
						placeholder="Password"
						required
						maxlength="20"
					/>
				</label>
				<div class="form-group">
					<button type="submit" class="btn btn-primary mt-8 w-full" disabled={loading}>
						{#if loading}
							<span class="loading loading-spinner loading-xs"></span>
						{:else}
							<LogIn class="h-4 w-4" />
						{/if}
						Login
					</button>
				</div>
			</form>
		</div>
	</div>
	<p>
		© {new Date().getFullYear()} LicenseToGo
	</p>
</div>
