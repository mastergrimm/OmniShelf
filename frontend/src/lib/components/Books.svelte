<script lang="ts">
	import Listing from "./Listing.svelte";

	import type { Book } from "$lib/types/types";
	export let books: Book[];

	function getExternalLink(item: Book | any) {
		return `https://www.goodreads.com/book/show/${item.book_id}`;
	}
</script>

<Listing
	title="Books"
	items={books}
	formUpload="?/booksUpload"
	formDelete="?/booksDelete"
	itemKey="title"
	itemScore="my_rating"
	link="books"
	{getExternalLink}
/>
