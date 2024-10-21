<script lang="ts">
	import Listing from "./Listing.svelte";
	import type { Media } from "$lib/types/types";
	export let movies: Media[];

	function getExternalLink(item: Media | any) {
		return `https://www.imdb.com/title/${item.const}`;
	}
</script>

<Listing
	title="Movies"
	items={movies}
	formUpload="?/moviesUpload"
	formDelete="?/mediaDelete"
	itemKey="title"
	link="movies"
	itemScore="your_rating"
	{getExternalLink}
/>
