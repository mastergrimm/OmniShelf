<script lang="ts">
	import type { PageData } from "./$types";

	import Movies from "$lib/components/Movies.svelte";
	import TvShows from "$lib/components/TvShows.svelte";
	import Books from "$lib/components/Books.svelte";
	import Anime from "$lib/components/Anime.svelte";
	import Manga from "$lib/components/Manga.svelte";
	import Singleplayer from "$lib/components/Singleplayer.svelte";
	import Multiplayer from "$lib/components/Multiplayer.svelte";

	export let data: PageData;

	let rankingLimit = 20;

	$: movies = data.movies.slice(0, rankingLimit);
	$: tvShows = data.tvShows.slice(0, rankingLimit);
	$: books = data.books.slice(0, rankingLimit);
	$: anime = data.anime.slice(0, rankingLimit);
	$: manga = data.manga.slice(0, rankingLimit);
	$: singleplayer = data.singleplayer.slice(0, rankingLimit);
	$: multiplayer = data.multiplayer.slice(0, rankingLimit);
</script>

<div class="home">
	<Movies {movies} />
	<TvShows {tvShows} />
	<Books {books} />
	<Anime {anime} />
	<Manga {manga} />
	<Singleplayer {singleplayer} />
	<Multiplayer {multiplayer} />
</div>

<style lang="scss">
	.home {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		grid-gap: var(--spacing-8);
	}
</style>

