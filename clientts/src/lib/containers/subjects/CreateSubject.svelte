<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import * as Alert from '$lib/components/ui/alert/index.js';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import * as Input from '$lib/components/ui/input/index.js';
	import Label from '$lib/components/ui/label/label.svelte';
	import { t } from '$lib/i18n/i18n';
	import { Plus, Save } from '@lucide/svelte';

	let { failed = false } = $props();
	let loading = $state(false);
	let open = $state(false);
	const maxDate = new Date().toISOString().split('T')[0];

	const handleModalClosed = () => {
		loading = false;
		failed = false;
	};
</script>

<Dialog.Root bind:open onOpenChange={handleModalClosed}>
	<Dialog.Trigger>
		<Button variant="outline">
			<Plus />
			{$t('subjects.createSubject')}
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{$t('subjects.createSubject')}</Dialog.Title>
			<Dialog.Description>
				{$t('subjects.createSubjectDescription')}
			</Dialog.Description>
		</Dialog.Header>
		{#if failed}
			<Alert.Root variant="destructive">
				<Alert.Description>{$t('subjects.createError')}</Alert.Description>
			</Alert.Root>
		{/if}
		<form
			class="grid grid-cols-2 gap-4"
			method="POST"
			action="?/create"
			use:enhance={() => {
				loading = true;
				return async ({ result, update }) => {
					loading = false;
					if (result.type === 'error' || result.type === 'failure') {
						failed = true;
						return;
					}
					applyAction(result);
					open = false;
					return update();
				};
			}}
		>
			<div class="col-span-1">
				<Label for="firstName">{$t('general.firstName')}</Label>
				<Input.Input type="text" name="firstName" id="lastName" required maxlength={120} />
			</div>
			<div class="col-span-1">
				<Label for="lastName">{$t('general.lastName')}</Label>
				<Input.Input type="text" name="lastName" id="lastName" required maxlength={120} />
			</div>

			<div class="col-span-1">
				<Label for="nickname">{$t('general.nickName')}</Label>
				<Input.Input type="text" name="nickname" id="nickname" maxlength={120} />
			</div>

			<div class="col-span-1">
				<Label for="birthdate">{$t('general.birthdate')}</Label>
				<Input.Input type="date" name="birthdate" id="birthdate" required max={maxDate} />
			</div>
			<Dialog.Footer class="col-span-2">
				<Button type="submit" disabled={loading}>
					<Save />
					{$t('subjects.createSubject')}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
