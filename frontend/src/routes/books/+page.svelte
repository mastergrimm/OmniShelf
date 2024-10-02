<script lang="ts">
	import type { PageData } from "./$types";

	import Books from "$lib/components/Books.svelte";

	export let data: PageData;

	$: books = data.books;
</script>

<div class="page">
	<h1>My Books</h1>
	<Books {books} />
</div>
