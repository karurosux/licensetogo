<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button/index.js';
	import { pb } from '$lib/pocketbase';
	import { LogOut, ShieldHalf } from 'lucide-svelte';
	import * as Breadcrumb from '$lib/components/ui/breadcrumb';
	import { breadcrumbs } from '$lib/context/breadcrumbs';

	let { children } = $props();

	const isLinkActive = (href: string) => {
		return page.url.pathname.includes(href);
	};
</script>

<div class="flex min-h-screen w-full flex-col">
	<header class="bg-background sticky top-0 flex h-16 items-center gap-4 border-b px-4 md:px-6">
		<nav
			class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6"
		>
			<a href="/home" class="flex items-center gap-2 text-lg font-semibold md:text-base">
				<div class="flex items-center justify-center gap-1 text-xl">
					<ShieldHalf class="h-6 w-6" />
					LicenseToGo
				</div>
				<span class="sr-only">Acme Inc</span>
			</a>
			<a
				href="/license-manager"
				class:text-foreground={isLinkActive('/license-manager')}
				class="hover:text-foreground w-16 transition-colors"
			>
				Licenses
			</a>
			<a
				href="/api-keys"
				class:text-foreground={isLinkActive('/api-keys')}
				class="text-muted-foreground hover:text-foreground w-16 transition-colors"
			>
				API Keys
			</a>
			<a
				href="/settings"
				class:text-foreground={isLinkActive('/settings')}
				class="text-muted-foreground hover:text-foreground w-16 transition-colors"
			>
				Settings
			</a>
		</nav>
		<div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
			<span class="flex-1"></span>
			<!-- <form class="ml-auto flex-1 sm:flex-initial"> -->
			<!-- 	<div class="relative"> -->
			<!-- 		<Search class="text-muted-foreground absolute top-2.5 left-2.5 h-4 w-4" /> -->
			<!-- 		<Input -->
			<!-- 			type="search" -->
			<!-- 			placeholder="Search license..." -->
			<!-- 			class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]" -->
			<!-- 		/> -->
			<!-- 	</div> -->
			<!-- </form> -->
			<Button
				variant="outline"
				onclick={() => {
					pb.authStore.clear();
					goto('/login');
				}}
			>
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
