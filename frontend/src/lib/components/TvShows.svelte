<script lang="ts">
	import Listing from "./Listing.svelte";
	export let tvShows: any[];
</script>

<Listing
	title="TV Shows"
	items={tvShows}
	formAction="?/tvShowsUpload"
	itemKey="title"
	link="tvshows"
/>
