<script lang="ts">
	import { Upload } from '@lucide/svelte';
	import type { Snippet } from 'svelte';

	interface Props {
		class?: string;
		hoverClass?: string;
		loading?: boolean;
		accepts?: HTMLInputElement['accept'];
		children?: Snippet;
		onfilechange?: (file: File) => void;
	}

	let fileInput: HTMLInputElement | null = $state(null);
	let {
		class: className = '',
		hoverClass = '',
		loading = false,
		children,
		accepts,
		onfilechange
	}: Props = $props();

	const handleClick = () => {
		if (fileInput) {
			fileInput.click();
		}
	};

	const handleFileChange = (e: Event) => {
		const file = (e.target as HTMLInputElement).files?.[0];
		if (file) {
			onfilechange?.(file);
		}
	};
</script>

<span class="relative cursor-pointer {className}">
	{@render children?.()}
	<!-- svelte-ignore a11y_consider_explicit_label -->
	<button
		class="absolute top-0 right-0 bottom-0 left-0 z-30 flex cursor-pointer items-center justify-center bg-white opacity-0 transition hover:opacity-90 {hoverClass}"
		onclick={handleClick}
	>
		<Upload class="z-50 text-black" />
	</button>
	{#if loading}
		<div class="absolute top-0 right-0 bottom-0 left-0">
			<!-- <LoadingSpinner absolute /> -->
		</div>
	{/if}
</span>
<input
	bind:this={fileInput}
	type="file"
	class="hidden"
	accept={accepts}
	onchange={handleFileChange}
/>
