<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { getCookie, setCookie } from '$lib/utils/cookie-store';
	import { Sun, Moon } from '@lucide/svelte';
	import { onMount } from 'svelte';

	let { class: className = '' } = $props();
	let isDark = $state(false);

	function handleToggleTheme() {
		const theme = getCookie('theme', window.document.cookie) || 'dark';
		isDark = theme === 'dark';
		setCookie('theme', isDark ? 'light' : 'dark', { path: '/' });
		isDark = !isDark;
		document.documentElement.classList.toggle('dark', isDark);
		updateColorScheme();
	}

	function updateColorScheme() {
		document.documentElement.style.colorScheme = isDark ? 'dark' : 'light';
	}

	onMount(() => {
		const theme = getCookie('theme', window.document.cookie) || 'dark';
		isDark = theme === 'dark';
		updateColorScheme();
	});
</script>

<Button variant="ghost" class="rounded-full" onclick={handleToggleTheme}>
	{#if isDark}
		<Sun />
	{:else}
		<Moon />
	{/if}
</Button>
