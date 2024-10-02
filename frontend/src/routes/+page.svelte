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

	$: movies = data.movies;
	$: tvShows = data.tvShows;
	$: books = data.books;
	$: anime = data.anime;
	$: manga = data.manga;
	$: singleplayer = data.singleplayer;
	$: multiplayer = data.multiplayer;
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
		grid-template-columns: repeat(5, 1fr);
		grid-gap: 1rem;
	}
</style>

