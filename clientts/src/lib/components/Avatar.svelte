<script lang="ts">
	import * as Avatar from './ui/avatar/index.js';

	type Props = { class?: string; name: string | undefined; url: string };

	let { name = '', url = '', class: className = '' }: Props = $props();
	let initials = $state('...');

	const getInitials = () => {
		const names = name?.split(' ');
		if (names) {
			return names[0].charAt(0) + names[1].charAt(0);
		}
		return '';
	};

	$effect(() => {
		if (name) {
			initials = getInitials();
		}
	});
</script>

<Avatar.Root class={className}>
	<Avatar.Image src={url} alt={name} />
	<Avatar.Fallback>
		{initials}
	</Avatar.Fallback>
</Avatar.Root>
