<script lang="ts">
	import { applyAction, enhance } from '$app/forms'
	import Button from '$lib/components/ui/button/button.svelte'
	import * as Dialog from '$lib/components/ui/dialog/index.js'
	import * as Input from '$lib/components/ui/input/index.js'
	import { t } from '$lib/i18n/i18n'
	import type { LicenseRecord } from '$lib/models/generated/pb-models'
	import { Trash } from '@lucide/svelte'

	type Props = {
		id: LicenseRecord['id']
		name: LicenseRecord['name']
	}

	const { id, name }: Props = $props()

	let open = $state(false)
	let loading = $state(false)
	let failed = $state(false)

	const handleModalClosed = () => {}
</script>

<Dialog.Root bind:open onOpenChange={handleModalClosed}>
	<Dialog.Trigger>
		<Button
			variant="outline"
			class="btn btn-sm btn-square btn-ghost border-red-500 text-red-500"
			onclick={() => {
				open = true
			}}>
			<Trash class="h-4 w-4" />
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{$t((ts) => ts.licenseManager.deleteLicense.title)}</Dialog.Title>
			<Dialog.Description>
				{$t((ts) => ts.licenseManager.deleteLicense.message, { name })}
			</Dialog.Description>
		</Dialog.Header>
		<form
			class="grid grid-cols-2 gap-4"
			method="POST"
			action="?/delete"
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
			<Input.Input type="hidden" name="id" id="id" required value={id} />
			<Dialog.Footer class="col-span-2">
				<Button variant="ghost" type="button" disabled={loading} onclick={() => (open = false)}>
					{$t((ts) => ts.general.cancel)}
				</Button>
				<Button type="submit" disabled={loading}>
					{$t((ts) => ts.general.delete)}
				</Button>
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
