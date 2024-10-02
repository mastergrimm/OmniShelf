<script lang="ts">
	import Listing from "./Listing.svelte";
	export let singleplayer: any[];
</script>

<Listing
	title="Single Player Games"
	items={singleplayer}
	formAction="?/singleplayerUpload"
	itemKey="game"
	link="singleplayer"
/>
