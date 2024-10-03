<script lang="ts">
	import Listing from "./Listing.svelte";
	export let books: any[];
</script>

<Listing
	title="Books"
	items={books}
	formUpload="?/booksUpload"
	formDelete="?/booksDelete"
	itemKey="title"
	itemScore="my_rating"
	link="books"
/>
