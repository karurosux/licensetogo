<script>
	/**
	 * @type {{
	 *  items: import('./Breadcrumb').Breadcrumb[]
	 * }}
	 */
	let { items } = $props();
</script>

<div class="breadcrumbs border-b-base-300 bg-base-200 w-full border-b p-6 text-sm">
	<ul>
		{#each items as item}
			<li>
				{#if item.icon}
					<item.icon class="mr-2" />
				{/if}
				<!-- svelte-ignore a11y_missing_attribute -->
				<a> {item.label} </a>
			</li>
		{/each}
	</ul>
</div>
