<script lang="ts">
	import type { Snippet } from 'svelte'

	type MenuItem = {
		route: string
		label: string
	}

	type Props = { children: Snippet<[MenuItem]> }

	let { children }: Props = $props()
	let items = $state<MenuItem[]>([])
</script>

{#each items || [] as i (i.route)}
	{@render children?.(i)}
{/each}
