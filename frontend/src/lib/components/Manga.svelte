<script lang="ts">
	import Listing from "./Listing.svelte";
	export let manga: any[];
</script>

<Listing
	title="Manga"
	items={manga}
	formUpload="?/mangaUpload"
	formDelete="?/mangaDelete"
	itemKey="manga_title"
	itemScore="my_score"
	link="manga"
/>
