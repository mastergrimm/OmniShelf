import type { Actions } from './$types';

async function handleRequest(endpoint: string, method: 'GET' | 'POST' | 'DELETE', file?: File) {
	const url = `http://omnishelf-backend-1:8080${endpoint}`;
	const options: RequestInit = { method };

	if (method === 'POST' && file) {
		const formData = new FormData();
		formData.append('file', file);
		options.body = formData;
	}

	if (method === 'DELETE') {
		options.method = 'DELETE';
	}

	try {
		const response = await fetch(url, options);
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const result = await response.json();
		return { status: 200, body: result };
	} catch (error) {
		console.error('Error:', error);
		return {
			status: 500,
			body: `An error occurred while ${method === 'DELETE' ? 'deleting' : 'uploading'} the data`
		};
	}
}

export const actions: Actions = {
	booksUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('books');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/books', 'POST', file);
	},

	booksDelete: async () => {
		return handleRequest('/books', 'DELETE');
	},

	moviesUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('movies');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/media', 'POST', file);
	},

	tvShowsUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('tvShows');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/media', 'POST', file);
	},

	mediaDelete: async () => {
		return handleRequest('/media', 'DELETE');
	},

	animeUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('anime');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/anime', 'POST', file);
	},

	animeDelete: async () => {
		return handleRequest('/anime', 'DELETE');
	},

	mangaUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('manga');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/manga', 'POST', file);
	},

	mangaDelete: async () => {
		return handleRequest('/manga', 'DELETE');
	},

	singleplayerUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('singleplayer');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/singleplayer', 'POST', file);
	},

	singleplayerDelete: async () => {
		return handleRequest('/singleplayer', 'DELETE');
	},

	multiplayerUpload: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('multiplayer');
		if (!(file instanceof File)) {
			return { status: 400, body: 'No file uploaded' };
		}
		return handleRequest('/multiplayer', 'POST', file);
	},

	multiplayerDelete: async () => {
		return handleRequest('/multiplayer', 'DELETE');
	}
};
