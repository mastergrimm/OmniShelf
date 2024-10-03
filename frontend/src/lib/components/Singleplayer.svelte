<script lang="ts">
	import Listing from "./Listing.svelte";

	export let singleplayer: any[];
</script>

<Listing
	title="Single Player Games"
	items={singleplayer}
	formUpload="?/singleplayerUpload"
	formDelete="?/singleplayerDelete"
	itemKey="name"
	itemScore="rating"
	link="singleplayer"
/>
