<script>
	import { goto } from '$app/navigation';
	import { APP_NAME } from '$lib/constants';
	import { Eye, KeyRound, Plus } from 'lucide-svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>API Key Generated | {APP_NAME}</title>
</svelte:head>

<div class="breadcrumbs bg-base-200 border-b-base-300 w-full border-b p-6 text-sm">
	<ul>
		<li>
			<a>
				<KeyRound />
				API Keys Manager
			</a>
		</li>
		<li>
			<a>
				<Plus />
				Create API Key
			</a>
		</li>
		<li>
			<a>
				<Eye />
				API Key Viewer
			</a>
		</li>
	</ul>
</div>
<div class="flex w-full items-center justify-center">
	<div class="flex flex-col items-center justify-center gap-4 p-4">
		<h1 class="text-lg font-bold">API Key Generated</h1>
		<p>Make sure to save this API Key, as it could not be displayed again.</p>
		<div class="mockup-code my-4 w-full">
			<pre data-prefix="$"><code class="text-white">{data.license}</code></pre>
		</div>
		<button class="btn btn-outline" onclick={() => goto('/apikeys')}>
			Go To API Key Manager
		</button>
	</div>
</div>
