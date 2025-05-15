<script lang="ts">
	import { goto } from '$app/navigation';
	import Avatar from '$lib/components/Avatar.svelte';
	import Pagination from '$lib/components/Pagination.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Dropdown from '$lib/components/ui/dropdown-menu/index.js';
	import * as Table from '$lib/components/ui/table/index.js';
	import { locale, t } from '$lib/i18n/i18n';
	import type { BranchesResponse, SubjectResponse } from '$lib/models/generated/pb-models';
	import { getFileUrl } from '$lib/utils/get-file-url';
	import { MoreVerticalIcon, PackageOpen } from '@lucide/svelte';
	import type { ListResult } from 'pocketbase';

	type Props = {
		data: ListResult<SubjectResponse<{ branch: BranchesResponse }>> | null;
	};

	let { data }: Props = $props();
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			<Table.Head>{$t('general.name')}</Table.Head>
			<Table.Head class="hidden sm:table-cell">{$t('general.nickName')}</Table.Head>
			<Table.Head class="hidden sm:table-cell">{$t('general.birthdate')}</Table.Head>
			<Table.Head class="hidden sm:table-cell">{$t('general.createdAt')}</Table.Head>
			<Table.Head class="hidden sm:table-cell">{$t('general.updatedAt')}</Table.Head>
			<Table.Head></Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#each data?.items || [] as s (s.id)}
			<Table.Row>
				<Table.Cell>
					<div class="flex items-center gap-x-4">
						<Avatar
							url={getFileUrl(s.collectionId, s.id, s.picture)}
							name={s.firstName + ' ' + s.lastName}
						/>
						<div class="truncate text-sm/6 font-medium">
							{s.firstName + ' ' + s.lastName}
						</div>
					</div>
				</Table.Cell>
				<Table.Cell class="hidden sm:table-cell">{s.nickname}</Table.Cell>
				<Table.Cell class="hidden sm:table-cell">
					<div class="flex items-center justify-end gap-x-2 sm:justify-start">
						<div>
							{new Date(s.birthdate).toLocaleDateString($locale, {
								year: 'numeric',
								month: 'long',
								day: 'numeric'
							})}
						</div>
					</div>
				</Table.Cell>
				<Table.Cell class="hidden sm:table-cell">
					<div class="flex items-center justify-end gap-x-2 sm:justify-start">
						<div>
							{new Date(s.created).toLocaleDateString($locale, {
								year: 'numeric',
								month: 'long',
								day: 'numeric'
							})}
						</div>
					</div>
				</Table.Cell>
				<Table.Cell class="hidden sm:table-cell">
					<div class="flex items-center justify-end gap-x-2 sm:justify-start">
						<div>
							{new Date(s.updated).toLocaleDateString($locale, {
								year: 'numeric',
								month: 'long',
								day: 'numeric'
							})}
						</div>
					</div>
				</Table.Cell>
				<Table.Cell>
					<Dropdown.Root>
						<Dropdown.Trigger>
							<Button variant="ghost">
								<MoreVerticalIcon />
							</Button>
						</Dropdown.Trigger>
						<Dropdown.Content>
							<Dropdown.Item onclick={() => goto(`/subjects/${s.id}`)}>
								<PackageOpen />
								{$t('general.open')}
							</Dropdown.Item>
						</Dropdown.Content>
					</Dropdown.Root>
				</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
<Pagination count={data?.totalItems || 0} perPage={data?.perPage} />
