<script lang="ts">
	import type { PageData } from "./$types";

	import Singleplayer from "$lib/components/Singleplayer.svelte";

	export let data: PageData;

	$: singleplayer = data.singleplayer;
</script>

<div class="page">
	<h1>My Singleplayer Games</h1>
	<Singleplayer {singleplayer} />
</div>
