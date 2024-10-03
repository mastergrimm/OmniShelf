<script lang="ts">
	import Listing from "./Listing.svelte";
	export let movies: any[];
</script>

<Listing
	title="Movies"
	items={movies}
	formUpload="?/moviesUpload"
	formDelete="?/mediaDelete"
	itemKey="title"
	link="movies"
	itemScore="your_rating"
/>
