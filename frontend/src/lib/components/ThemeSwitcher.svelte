<script lang="ts">
	import Icon from "@iconify/svelte";

	let themeIcon = $state("light");
</script>

<button class="theme">
	<Icon icon={themeIcon} />
</button>

