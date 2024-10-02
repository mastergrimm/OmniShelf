<script lang="ts">
	import type { PageData } from "./$types";

	import Movies from "$lib/components/Movies.svelte";

	export let data: PageData;

	$: movies = data.movies;
</script>

<div class="page">
	<h1>My Books</h1>
	<Movies {movies} />
</div>
