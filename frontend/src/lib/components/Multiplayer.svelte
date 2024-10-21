<script lang="ts">
	import Listing from "./Listing.svelte";
	import type { Game } from "$lib/types/types";
	export let multiplayer: Game[];

	function getExternalLink(item: Game | any) {
		return `https://store.steampowered.com/app/${item.appid}`;
	}
</script>

<Listing
	title="Multiplayer Games"
	items={multiplayer}
	formUpload="?/multiplayerUpload"
	formDelete="?/multiplayerDelete"
	itemKey="name"
	itemScore="rating"
	link="multiplayer"
	{getExternalLink}
/>
