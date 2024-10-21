<script lang="ts">
	import Listing from "./Listing.svelte";
	import type { Anime } from "$lib/types/types";
	export let anime: Anime[];

	function getExternalLink(item: Anime | any) {
		return `https://myanimelist.net/anime/${item.series_animedb_id}`;
	}
</script>

<Listing
	title="Anime"
	items={anime}
	formUpload="?/animeUpload"
	formDelete="?/animeDelete"
	itemKey="series_title"
	itemScore="my_score"
	link="anime"
	{getExternalLink}
/>
