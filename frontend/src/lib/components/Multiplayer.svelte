<script lang="ts">
	import Listing from "./Listing.svelte";
	export let multiplayer: any[];
</script>

<Listing
	title="Multiplayer Games"
	items={multiplayer}
	formAction="?/multiplayerUpload"
	itemKey="game"
	link="multiplayer"
/>
