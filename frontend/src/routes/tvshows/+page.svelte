<script lang="ts">
	import type { PageData } from "./$types";

	import TvShows from "$lib/components/TvShows.svelte";

	export let data: PageData;

	$: tvshows = data.tvShows;
</script>

<div class="page">
	<h1>My Books</h1>
	<TvShows tvShows={tvshows} />
</div>
