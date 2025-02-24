<script>
	import { goto } from '$app/navigation';

	let { data } = $props();
</script>

<div class="flex w-full items-center justify-center">
	<div class="flex flex-col items-center justify-center gap-4 p-4">
		<h1 class="text-lg font-bold">License Generated</h1>
		<p>Make sure to save this license, as it will not be displayed again.</p>
		<div class="mockup-code my-4 w-full">
			<pre data-prefix="$"><code class="text-white">{data.license}</code></pre>
		</div>
		<button class="btn btn-outline" onclick={() => goto('/license-manager')}>
			Go To License Manager</button
		>
	</div>
</div>
