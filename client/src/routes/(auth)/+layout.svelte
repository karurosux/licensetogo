<script>
	import { goto } from '$app/navigation';
	import { pb } from '$lib/utils/pb';
	import { KeyRound, LogOut, Scroll, ShieldHalf } from 'lucide-svelte';

	let { children } = $props();

	const handleLogout = () => {
		pb.authStore.clear();
		goto('/login');
	};
</script>

<div class="drawer lg:drawer-open">
	<input id="main-drawer" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content bg-base-200 flex h-screen flex-col">
		<div class="navbar bg-base-100 border-b-base-300 border-b">
			<div class="navbar-start pl-4"></div>
			<div class="navbar-end">
				<button class="btn btn-ghost" onclick={handleLogout}>
					<LogOut class="h-4 w-4" />
					Logout
				</button>
			</div>
		</div>
		<!-- Page content here -->
		<div class="flex-1">
			{@render children?.()}
		</div>
		<!-- <div class="border-base-300 bg-base-200 min-h-16 border-t"></div> -->
	</div>
	<div class="drawer-side bg-base-100 border-r-base-300 border-r">
		<label for="main-drawer" aria-label="close sidebar" class="drawer-overlay"></label>

		<div class="navbar border-b-base-300 border-b">
			<div class="navbar-center">
				<a
					class="btn btn-ghost bg-base-100 rounded-none border-none text-xl shadow-none"
					href="/license-manager"
				>
					<ShieldHalf class="h-6 w-6" />
					LicenseToGo
				</a>
			</div>
		</div>
		<ul class="menu menu-lg bg-base-100 text-base-content w-56 p-4">
			<!-- Sidebar content here -->
			<li>
				<a href="/license-manager">
					<Scroll />
					Licenses
				</a>
			</li>
			<li>
				<a href="/apikeys">
					<KeyRound />
					API Keys
				</a>
			</li>
		</ul>
	</div>
</div>
