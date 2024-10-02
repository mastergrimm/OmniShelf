<script lang="ts">
	import type { PageData } from "./$types";

	import Anime from "$lib/components/Anime.svelte";

	export let data: PageData;

	$: anime = data.anime;
</script>

<div class="books-page">
	<h1>My Anime</h1>
	<Anime {anime} />
</div>
