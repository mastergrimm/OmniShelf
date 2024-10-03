<script lang="ts">
	import { enhance } from "$app/forms";
	import { invalidateAll } from "$app/navigation";

	export let title: string;
	export let items: any[];
	export let formUpload: string;
	export let formDelete: string;
	export let itemKey: string;
	export let itemScore: string;
	export let link: string;

	let loading = false;
</script>

<div class="grid">
	<a href="/{link}" class="grid__title">{title}</a>
	<div class="grid__upload">
		<form
			method="POST"
			action={formUpload}
			enctype="multipart/form-data"
			use:enhance={() => {
				loading = true;

				return async ({ update }) => {
					await update();
					loading = false;
				};
			}}
		>
			<button type="submit">Submit</button>
			<label>
				<input type="file" name={link} accept=".csv" />
			</label>
		</form>
	</div>
	<div class="grid__list">
		{#if loading}
			<div class="grid__item">Loading...</div>
		{/if}
		{#if items.length === 0}
			<div class="grid__item">No {title.toLowerCase()} found</div>
		{:else}
			{#key items}
				{#each items as item, index}
					<div class="grid__item">
						<div class="item__score">
							{item[itemScore]}
						</div>
						<div class="item__title">{item[itemKey]}</div>
					</div>
				{/each}
			{/key}

			<form
				method="POST"
				action={formDelete}
				enctype="multipart/form-data"
				use:enhance
			>
				<button type="submit">CLEAR</button>
			</form>
		{/if}
	</div>
</div>

<style lang="scss">
	.grid {
		border: 2px solid var(--clr-info-200);
		border-radius: var(--rounded-lg);
		padding: var(--spacing-4);
	}

	.grid__title {
		font-size: 2rem;
		font-weight: bold;
		text-align: center;
	}

	.grid__upload {
		display: flex;
		align-items: center;

		margin: var(--spacing-4) 0;
		width: 100%;

		input[type="file"] {
			width: 100%;
		}

		form {
			display: flex;
			border-radius: var(--rounded-lg);
			overflow: hidden;
		}

		label {
			background-color: var(--clr-info-200);
			padding: var(--spacing-2) var(--spacing-4);
			color: var(--clr-info-900);

			&:hover {
				background-color: var(--clr-success-300);
				cursor: pointer;
			}
		}

		button {
			font-weight: bold;
			text-transform: uppercase;
			letter-spacing: 0.04em;
			background-color: var(--clr-info-800);
			padding: var(--spacing-2) var(--spacing-4);
			color: var(--clr-text);

			&:hover {
				background-color: var(--clr-success-900);
			}
		}
	}

	.grid__list {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.grid__item {
		padding: var(--spacing-2) var(--spacing-4);
		border-radius: var(--rounded-lg);
	}
</style>
