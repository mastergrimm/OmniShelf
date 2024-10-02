<script lang="ts">
	import type { PageData } from "./$types";

	import Multiplayer from "$lib/components/Movies.svelte";

	export let data: PageData;

	$: multiplayer = data.multiplayer;
</script>

<div class="page">
	<h1>My Multiplayer Games</h1>
	<Multiplayer movies={multiplayer} />
</div>
