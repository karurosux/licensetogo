<script lang="ts">
	import { browser } from '$app/environment';
	import { locale, locales, t } from '$lib/i18n/i18n';
	import * as Select from '$lib/components/ui/select/index.js';

	$effect(() => {
		if (browser) {
			window.localStorage.setItem('locale', $locale);
		}
	});
</script>

<Select.Root type="single" name="language" bind:value={$locale}>
	<Select.Trigger>{$t('langs.' + $locale)}</Select.Trigger>
	<Select.Content>
		{#each locales as l}
			<Select.Item value={l}>{$t(`langs.${l}`)}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>
