<script lang="ts">
	import Listing from "./Listing.svelte";
	import type { Media } from "$lib/types/types";
	export let tvShows: Media[];

	function getExternalLink(item: Media | any) {
		return `https://www.imdb.com/title/${item.const}`;
	}
</script>

<Listing
	title="TV Shows"
	items={tvShows}
	formUpload="?/tvShowsUpload"
	formDelete="?/mediaDelete"
	itemKey="title"
	link="tvshows"
	itemScore="your_rating"
	{getExternalLink}
/>
