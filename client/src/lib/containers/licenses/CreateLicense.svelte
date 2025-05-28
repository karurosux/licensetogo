<script lang="ts">
	import { applyAction, enhance } from '$app/forms'
	import * as Alert from '$lib/components/ui/alert/index.js'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import * as Input from '$lib/components/ui/input/index.js'
	import Label from '$lib/components/ui/label/label.svelte'
	import { t } from '$lib/i18n/i18n'
	import { Plus, Save } from '@lucide/svelte'

	let { failed = false } = $props()
	let loading = $state(false)
	let open = $state(false)
	const minDate = new Date().toISOString().split('T')[0]

	const handleModalClosed = () => {
		loading = false
		failed = false
	}
</script>

<Dialog.Root bind:open onOpenChange={handleModalClosed}>
	<Dialog.Trigger>
		<Button variant="outline">
			<Plus />
			{$t((ts) => ts.licenseManager.createLicense.title)}
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{$t((ts) => ts.licenseManager.createLicense.title)}</Dialog.Title>
			<Dialog.Description>
				{$t((ts) => ts.licenseManager.createLicense.description)}
			</Dialog.Description>
		</Dialog.Header>
		{#if failed}
			<Alert.Root variant="destructive">
				<Alert.Description>
					{$t((ts) => ts.licenseManager.createLicense.createError)}
				</Alert.Description>
			</Alert.Root>
		{/if}
		<form
			class="grid grid-cols-2 gap-4"
			method="POST"
			action="?/create"
			use:enhance={() => {
				loading = true
				return async ({ result, update }) => {
					loading = false
					if (result.type === 'error' || result.type === 'failure') {
						failed = true
						return
					}
					applyAction(result)
					open = false
					return update()
				}
			}}>
			<div class="col-span-2">
				<Label for="name">{$t((ts) => ts.general.name)}</Label>
				<Input.Input type="text" name="name" id="name" required maxlength={120} />
			</div>

			<div class="col-span-2">
				<Label for="expires">{$t((ts) => ts.general.expires)}</Label>
				<Input.Input type="date" name="expires" id="expires" min={minDate} />
			</div>

			<Dialog.Footer class="col-span-2">
				<Button type="submit" disabled={loading}>
					<Save />
					{$t((ts) => ts.licenseManager.createLicense.title)}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
