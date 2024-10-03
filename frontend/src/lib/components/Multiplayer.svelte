<script lang="ts">
	import Listing from "./Listing.svelte";
	export let multiplayer: any[];
</script>

<Listing
	title="Multiplayer Games"
	items={multiplayer}
	formUpload="?/multiplayerUpload"
	formDelete="?/multiplayerDelete"
	itemKey="name"
	itemScore="rating"
	link="multiplayer"
/>
