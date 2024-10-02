<script lang="ts">
	import Listing from "./Listing.svelte";
	export let anime: any[];
</script>

<Listing
	title="Anime"
	items={anime}
	formAction="?/animeUpload"
	itemKey="series_title"
	link="anime"
/>
