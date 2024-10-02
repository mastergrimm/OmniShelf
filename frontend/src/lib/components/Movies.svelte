<script lang="ts">
	import Listing from "./Listing.svelte";
	export let movies: any[];
</script>

<Listing
	title="Movies"
	items={movies}
	formAction="?/moviesUpload"
	itemKey="title"
	link="movies"
/>
