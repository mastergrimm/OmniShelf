<script lang="ts">
	import "$lib/styles/app.scss";
	import "$lib/styles/app.css";

	import type { LayoutData } from "./$types";
	import ThemeSwitcher from "$lib/components/ThemeSwitcher.svelte";

	export let data: LayoutData;
</script>

<main>
	<a href="/" class="header">
		<div class="title">OmniShelf</div>
		<div class="caption">Collection of media</div>
		<div class="theme">
			<ThemeSwitcher />
		</div>
	</a>
	<slot {data} />
</main>

<style lang="postcss">
	main {
		@apply p-[5%];
	}
</style>
