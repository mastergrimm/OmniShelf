import type { Actions } from './$types';
import type { Media, Book, Anime, Manga, Game } from '$lib/types/types';

export const load = async () => {
	let url = "omnishelf-backend-1:8080";

	const fetchData = async (endpoint: string) => {
		try {
			let response = await fetch(`http://${url}${endpoint}`);
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			let data = await response.json();
			return data;
		} catch (error) {
			console.error(`Error fetching data from ${endpoint}:`, error);
			return null;
		}
	};

	const sortData = (data: any[], key: string) => {
		return data ? [...data].sort((a: any, b: any) => b[key] - a[key]) : [];
	};

	try {
		const [booksData, mediaData, animeData, mangaData, singleplayerData, multiplayerData] = await Promise.all([
			fetchData('/books'),
			fetchData('/media'),
			fetchData('/anime'),
			fetchData('/manga'),
			fetchData('/singleplayer'),
			fetchData('/multiplayer')
		]);

		return {
			books: sortData(booksData || [], 'my_rating'),
			movies: sortData(mediaData?.movies || [], 'your_rating'),
			tvShows: sortData(mediaData?.tvShows || [], 'your_rating'),
			anime: sortData(animeData || [], 'my_score'),
			manga: sortData(mangaData || [], 'my_score'),
			singleplayer: sortData(singleplayerData || [], 'rating'),
			multiplayer: sortData(multiplayerData || [], 'rating')
		};
	} catch (error) {
		console.error('Error in load function:', error);
		return {
			books: [],
			movies: [],
			tvShows: [],
			anime: [],
			manga: [],
			singleplayer: [],
			multiplayer: []
		};
	}
};
