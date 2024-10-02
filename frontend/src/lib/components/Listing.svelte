<script lang="ts">
	import { enhance } from "$app/forms";

	export let title: string;
	export let items: any[];
	export let formAction: string;
	export let itemKey: string;
	export let link: string;
</script>

<div class="grid">
	<a href="/{link}" class="grid__title">{title}</a>
	<div class="grid__upload">
		<form
			method="POST"
			action={formAction}
			enctype="multipart/form-data"
			use:enhance
		>
			<label>
				<input type="file" name={title.toLowerCase()} accept=".csv" />
			</label>
			<button type="submit">Submit</button>
		</form>
	</div>
	<div class="grid__list">
		{#if items.length === 0}
			<div class="grid__item">No {title.toLowerCase()} found</div>
		{:else}
			{#each items as item}
				<div class="grid__item">
					{item[itemKey]}
				</div>
			{/each}
		{/if}
	</div>
</div>

<style lang="scss">
	.grid {
		/* Grid styles */
	}

	.grid__title {
		font-size: 2rem;
		font-weight: bold;
		text-align: center;
	}

	.grid__upload {
		padding: var(--spacing-4);
		display: flex;
		justify-content: center;
		align-items: center;

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
			padding: var(--spacing-2);
			color: var(--clr-bg);

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
</style>
