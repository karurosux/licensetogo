<script>
	import { goto } from '$app/navigation';
	import Breadcrumbs from '$lib/components/breadcrumbs/Breadcrumbs.svelte';
	import { APP_NAME } from '$lib/constants';
	import { Eye, KeyRound, Plus } from 'lucide-svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>API Key Generated | {APP_NAME}</title>
</svelte:head>

<Breadcrumbs
	items={[
		{ label: 'API Keys Manager', icon: KeyRound },
		{ label: 'Create API Key', icon: Plus },
		{ label: 'API Key Created', icon: Eye }
	]}
></Breadcrumbs>

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
