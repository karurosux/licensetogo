<script>
	import { page } from '$app/state';
	import Button from '$lib/components/ui/button/button.svelte';
	import { t } from '$lib/i18n/i18n';
</script>

{#if page.status === 403}
	<main class="bg-base3 mt-10 grid min-h-full place-items-center px-6 py-24 sm:py-32 lg:px-8">
		<div class="text-center">
			<p class="!text-primary-500 text-3xl font-semibold">{page.status}</p>
			<h1 class="mt-4 text-5xl font-semibold tracking-tight text-balance text-red-500 sm:text-7xl">
				{$t('errorPage.forbidden.title')}
			</h1>
			<p class="text-base1 mt-6 text-lg font-medium text-pretty sm:text-xl/8">
				{$t('errorPage.forbidden.description')}
			</p>
			<div class="mt-10 flex items-center justify-center gap-6">
				<Button href="/">{$t('errorPage.goBack')}</Button>
			</div>
		</div>
	</main>
{/if}
{#if page.status === 404}
	<main class="bg-base3 mt-10 grid min-h-full place-items-center px-6 py-24 sm:py-32 lg:px-8">
		<div class="text-center">
			<p class="!text-primary-500 text-3xl font-semibold">{page.status}</p>
			<h1 class="mt-4 text-5xl font-semibold tracking-tight text-balance sm:text-7xl">
				{$t('errorPage.notFound.title')}
			</h1>
			<p class="text-base1 mt-6 text-lg font-medium text-pretty sm:text-xl/8">
				{$t('errorPage.notFound.description')}
			</p>
			<div class="mt-10 flex items-center justify-center gap-6">
				<Button href="/">{$t('errorPage.goBack')}</Button>
			</div>
		</div>
	</main>
{/if}
