import type { Actions } from './$types';

const handleUpload = async (file: File, endpoint: string) => {
	const url = `http://omnishelf-backend-1:8080${endpoint}`;
	const uploadFormData = new FormData();
	uploadFormData.append('file', file);

	try {
		const response = await fetch(url, {
			method: 'POST',
			body: uploadFormData
		});

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}

		const result = await response.text();
		return {
			status: 200,
			body: result
		};
	} catch (error) {
		console.error('Error:', error);
		return {
			status: 500,
			body: 'An error occurred while uploading the file'
		};
	}
}

export const actions: Actions = {
	moviesUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('movies');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/media');
	},

	tvShowsUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('tvShows');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/media');
	},

	booksUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('books');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/books');
	},

	animeUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('anime');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/anime');
	},

	mangaUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('manga');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/manga');
	},

	singleplayerUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('singleplayer');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/singleplayer');
	},

	multiplayerUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('multiplayer');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleUpload(file, '/multiplayer');
	}
};
