<script lang="ts">
	import Listing from "./Listing.svelte";
	export let anime: any[];
</script>

<Listing
	title="Anime"
	items={anime}
	formUpload="?/animeUpload"
	formDelete="?/animeDelete"
	itemKey="series_title"
	itemScore="my_score"
	link="anime"
/>
