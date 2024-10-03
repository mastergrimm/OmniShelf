<script lang="ts">
	import Listing from "./Listing.svelte";
	export let tvShows: any[];
</script>

<Listing
	title="TV Shows"
	items={tvShows}
	formUpload="?/tvShowsUpload"
	formDelete="?/mediaDelete"
	itemKey="title"
	link="tvshows"
	itemScore="your_rating"
/>
