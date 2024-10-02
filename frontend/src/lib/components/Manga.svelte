<script lang="ts">
	import Listing from "./Listing.svelte";
	export let manga: any[];
</script>

<Listing
	title="Manga"
	items={manga}
	formAction="?/mangaUpload"
	itemKey="manga_title"
	link="manga"
/>
