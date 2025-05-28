<script lang="ts">
	import * as Pagination from '$lib/components/ui/pagination/index.js';
	import { t } from '$lib/i18n/i18n';
	import { mergeStateWithQuery } from '$lib/utils/query-params';
	import ChevronLeft from '@lucide/svelte/icons/chevron-left';
	import ChevronRight from '@lucide/svelte/icons/chevron-right';
	import { MediaQuery } from 'svelte/reactivity';

	let { count = 0, perPage = 5 } = $props();

	const isDesktop = new MediaQuery('(min-width: 768px)');
	const siblingCount = $derived(isDesktop.current ? 1 : 0);
	let currentPage = $state(1);

	$effect(() => {
		mergeStateWithQuery({ page: currentPage, perPage });
	});
</script>

<Pagination.Root class="mt-4" bind:page={currentPage} {count} {perPage} {siblingCount}>
	{#snippet children({ pages, currentPage })}
		<Pagination.Content>
			<Pagination.Item>
				<Pagination.PrevButton>
					<ChevronLeft class="size-4" />
					<span class="hidden sm:block">{$t('general.previous')}</span>
				</Pagination.PrevButton>
			</Pagination.Item>
			{#each pages as page (page.key)}
				{#if page.type === 'ellipsis'}
					<Pagination.Item>
						<Pagination.Ellipsis />
					</Pagination.Item>
				{:else}
					<Pagination.Item>
						<Pagination.Link {page} isActive={currentPage === page.value}>
							{page.value}
						</Pagination.Link>
					</Pagination.Item>
				{/if}
			{/each}
			<Pagination.Item>
				<Pagination.NextButton>
					<span class="hidden sm:block">{$t('general.next')}</span>
					<ChevronRight class="size-4" />
				</Pagination.NextButton>
			</Pagination.Item>
		</Pagination.Content>
	{/snippet}
</Pagination.Root>
