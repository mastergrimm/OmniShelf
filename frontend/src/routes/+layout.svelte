<script lang="ts">
	import "$lib/styles/app.scss";
	import type { LayoutData } from "./$types";
	export let data: LayoutData;
</script>

<main>
	<div class="header">
		<div class="title">OmniShelf</div>
		<div class="caption">Collection of media</div>
	</div>
	<slot {data} />
</main>

<style lang="scss">
	main {
		padding: 5%;
	}
	.header {
		margin-bottom: var(--spacing-8);
	}
	.title {
		font-size: var(--text-5xl);
		font-weight: bold;
		text-align: center;
	}
	.caption {
		font-size: var(--text-2xl);
		text-align: center;
	}
</style>
