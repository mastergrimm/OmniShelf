export interface Media {
	const: string;
	your_rating: string;
	date_rated: string;
	title: string;
	original_title: string;
	url: string;
	title_type: string;
	imdb_rating: string;
	runtime: string;
	year: string;
	genres: string;
	num_votes: string;
	release_date: string;
	directors: string;
}

export interface Book {
	book_id: string;
	title: string;
	author: string;
	author_lf: string;
	additional_authors: string;
	isbn: string;
	isbn13: string;
	my_rating: number;
	average_rating: number;
	publisher: string;
	binding: string;
	number_of_pages: number;
	year_published: number;
	original_publication_year: number;
	date_read: string;
	date_added: string;
	bookshelves: string;
	bookshelves_with_positions: string;
	exclusive_shelf: string;
	my_review: string;
	spoiler: string;
	private_notes: string;
	read_count: number;
	owned_copies: number;
}

export interface Anime {
	series_animedb_id: number;
	series_title: string;
	series_type: string;
	series_episodes: number;
	my_id: number;
	my_watched_episodes: number;
	my_start_date: string;
	my_finish_date: string;
	my_rated: string;
	my_score: number;
	my_storage: string;
	my_storage_value: number;
	my_status: string;
	my_comments: string;
	my_times_watched: number;
	my_rewatch_value: number;
	my_priority: number;
	my_tags: string;
	my_rewatching: boolean;
	my_rewatching_ep: number;
	my_discuss: boolean;
	my_sns: string;
	update_on_import: boolean;
}

export interface Manga {
	manga_mangadb_id: number;
	manga_title: string;
	manga_volumes: number;
	manga_chapters: number;
	my_id: number;
	my_read_volumes: number;
	my_read_chapters: number;
	my_start_date: string;
	my_finish_date: string;
	my_scanalation_group: string;
	my_score: number;
	my_storage: string;
	my_retail_volumes: number;
	my_status: string;
	my_comments: string;
	my_times_read: number;
	my_tags: string;
	my_priority: number;
	my_reread_value: number;
	my_rereading: boolean;
	my_discuss: boolean;
	my_sns: string;
	update_on_import: boolean;
}

export interface Game {
	id: number;
	name: string;
	edition: string;
	platform: string;
	format: string;
	region: string;
	nowPlaying: boolean;
	backlogged: boolean;
	ownershipStatus: string;
	progressStatus: string;
	rating: number;
	initialReleaseDate: string;
	itemReleaseDate: string;
	addedOn: string;
	genre: string;
}
