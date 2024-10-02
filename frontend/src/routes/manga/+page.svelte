<script lang="ts">
	import type { PageData } from "./$types";

	import Manga from "$lib/components/Manga.svelte";

	export let data: PageData;

	$: manga = data.manga;
</script>

<div class="page">
	<h1>My Manga</h1>
	<Manga {manga} />
</div>
