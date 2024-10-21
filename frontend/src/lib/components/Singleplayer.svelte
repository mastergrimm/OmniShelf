<script lang="ts">
	import Listing from "./Listing.svelte";
	import type { Game } from "$lib/types/types";

	export let singleplayer: Game[];

	function getExternalLink(item: Game | any) {
		return `https://store.steampowered.com/app/${item.appid}`;
	}
</script>

<Listing
	title="Single Player Games"
	items={singleplayer}
	formUpload="?/singleplayerUpload"
	formDelete="?/singleplayerDelete"
	itemKey="name"
	itemScore="rating"
	link="singleplayer"
	{getExternalLink}
/>
