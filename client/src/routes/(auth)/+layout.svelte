<script>
	import { goto } from '$app/navigation';
	import { pb } from '$lib/utils/pb';
	import { ShieldHalf, LogOut } from 'lucide-svelte';

	let { children } = $props();

	const handleLogout = () => {
		pb.authStore.clear();
		goto('/login');
	};
</script>

<div class="navbar bg-base-100 container m-auto">
	<div class="navbar-start">
		<div class="dropdown">
			<div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h8m-8 6h16"
					/>
				</svg>
			</div>
			<ul
				tabindex="0"
				class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
			>
				<li><a href="/license-manager">Licenses</a></li>
				<li><a href="/apikeys">API Keys</a></li>
			</ul>
		</div>
		<a class="btn btn-ghost text-xl" href="/license-manager">
			<ShieldHalf class="h-6 w-6" />
			LicenseToGo
		</a>
	</div>
	<div class="navbar-center hidden lg:flex">
		<ul class="menu menu-horizontal px-1">
			<li><a href="/license-manager">Licenses</a></li>
			<li><a href="/apikeys">API Keys</a></li>
		</ul>
	</div>
	<div class="navbar-end">
		<button class="btn btn-ghost" onclick={handleLogout}>
			<LogOut class="h-4 w-4" />
			Logout
		</button>
	</div>
</div>
<div class="container m-auto p-4">{@render children?.()}</div>
