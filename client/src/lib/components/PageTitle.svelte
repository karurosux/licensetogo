<script lang="ts">
	import { t } from '$lib/i18n/i18n';
	import { title as titl } from '$lib/context/title';

	let { title = '' } = $props();

	$effect(() => {
		titl.set($t(title));
	});
</script>

<svelte:head>
	<title>{$t(title)} | {$t('appName')}</title>
</svelte:head>
