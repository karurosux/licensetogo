<script lang="ts">
	import { goto } from '$app/navigation'
	import * as Breadcrumb from '$lib/components/ui/breadcrumb'
	import { Button } from '$lib/components/ui/button/index.js'
	import { APP_NAME } from '$lib/constants'
	import { breadcrumbs } from '$lib/context/breadcrumbs'
	import { t } from '$lib/i18n/i18n'
	import { pb } from '$lib/pocketbase'
	import { LogOut, ShieldHalf } from 'lucide-svelte'

	let { children } = $props()
</script>

<div class="flex min-h-screen w-full flex-col">
	<header class="bg-background sticky top-0 flex h-16 items-center gap-4 border-b px-4 md:px-6">
		<nav
			class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
			<a href="/home" class="flex items-center gap-2 text-lg font-semibold md:text-base">
				<div class="flex items-center justify-center gap-1 text-xl">
					<ShieldHalf class="h-6 w-6" />
					{APP_NAME}
				</div>
				<!-- <span class="sr-only">Acme Inc</span> -->
			</a>
			<a href="/license-manager" class="hover:text-foreground w-16 transition-colors">
				{$t((ts) => ts.general.licenses)}
			</a>
			<a
				href="/api-keys"
				class="text-muted-foreground hover:text-foreground w-16 transition-colors">
				{$t((ts) => ts.general.apiKeys)}
			</a>
			<a
				href="/settings"
				class="text-muted-foreground hover:text-foreground w-16 transition-colors">
				{$t((ts) => ts.general.settings)}
			</a>
		</nav>
		<div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
			<span class="flex-1"></span>
			<Button
				variant="outline"
				onclick={() => {
					pb.authStore.clear()
					goto('/login')
				}}>
				<LogOut />
				Logout
			</Button>
		</div>
	</header>
	<main class="container m-auto flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
		<Breadcrumb.Root class="hidden md:flex">
			<Breadcrumb.List>
				{#each $breadcrumbs as item}
					<Breadcrumb.Item>
						<Breadcrumb.Link href={item.href}>{item.label}</Breadcrumb.Link>
					</Breadcrumb.Item>
					<Breadcrumb.Separator class="last:hidden"></Breadcrumb.Separator>
				{/each}
			</Breadcrumb.List>
		</Breadcrumb.Root>
		{@render children?.()}
	</main>
</div>
