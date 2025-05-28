<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import PageTitle from '$lib/components/PageTitle.svelte';
	import { pb } from '$lib/pocketbase.js';
	import { Lock, LogIn, Mail, ShieldHalf } from 'lucide-svelte';
	import * as Card from '$lib/components/ui/card';
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	let loading = $state(false);
	let { form } = $props();
</script>

<PageTitle title="Login" />
<div class="flex h-screen w-screen flex-col items-center justify-center gap-4 p-4">
	{#if form?.failed}
		<div class="alert alert-error w-80" role="alert">
			<strong>Error!</strong> Invalid email or password.
		</div>
	{/if}
	<Card.Root class="w-80">
		<Card.Content>
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
					<Input
						type="email"
						class="grow"
						id="email"
						name="email"
						placeholder="Email"
						required
						maxlength={60}
					/>
				</label>
				<label class="input input-bordered invalid:input-error flex items-center gap-2">
					<Lock class="h-5 w-5" />
					<Input
						type="password"
						class="grow"
						id="password"
						name="password"
						placeholder="Password"
						required
						maxlength={20}
					/>
				</label>
				<div class="form-group">
					<Button type="submit" class="btn btn-primary mt-8 w-full" disabled={loading}>
						{#if loading}
							<span class="loading loading-spinner loading-xs"></span>
						{:else}
							<LogIn class="h-4 w-4" />
						{/if}
						Login
					</Button>
				</div>
			</form>
		</Card.Content>
	</Card.Root>
	<Card.Content></Card.Content>
	<p>
		© {new Date().getFullYear()} LicenseToGo
	</p>
</div>
