<script lang="ts">
	import Listing from "./Listing.svelte";
	export let books: any[];
</script>

<Listing
	title="Books"
	items={books}
	formAction="?/booksUpload"
	itemKey="title"
	link="books"
/>
