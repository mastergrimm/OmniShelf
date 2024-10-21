<script lang="ts">
	import Listing from "./Listing.svelte";

	import type { Manga } from "$lib/types/types";
	export let manga: Manga[];

	function getExternalLink(item: Manga | any) {
		return `https://myanimelist.net/manga/${item.manga_mangadb_id}`;
	}
</script>

<Listing
	title="Manga"
	items={manga}
	formUpload="?/mangaUpload"
	formDelete="?/mangaDelete"
	itemKey="manga_title"
	itemScore="my_score"
	link="manga"
	{getExternalLink}
/>
