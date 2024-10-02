import type { Media, Book, Anime, Manga, Game } from '$lib/types';

declare global {
	namespace App {
		interface Locals { }
		interface PageData {
			books: Book[];
			movies: Media[];
			tvShows: Media[];
			anime: Anime[];
			manga: Manga[];
			singleplayer: Game[];
			multiplayer: Game[];
		}
	}
}

export { };
