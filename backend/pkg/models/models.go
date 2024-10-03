package models

type Media struct {
	Const          string `json:"const" db:"const"`
	Your_Rating    string `json:"your_rating" db:"your_rating"`
	Date_Rated     string `json:"date_rated" db:"date_rated"`
	Title          string `json:"title" db:"title"`
	Original_Title string `json:"original_title" db:"original_title"`
	URL            string `json:"url" db:"url"`
	Title_Type     string `json:"title_type" db:"title_type"`
	IMDb_Rating    string `json:"imdb_rating" db:"imdb_rating"`
	Runtime        string `json:"runtime" db:"runtime"`
	Year           string `json:"year" db:"year"`
	Genres         string `json:"genres" db:"genres"`
	Num_Votes      string `json:"num_votes" db:"num_votes"`
	Release_Date   string `json:"release_date" db:"release_date"`
	Directors      string `json:"directors" db:"directors"`
}

type Book struct {
	Book_ID                    string  `json:"book_id" db:"book_id"`
	Title                      string  `json:"title" db:"title"`
	Author                     string  `json:"author" db:"author"`
	Author_LF                  string  `json:"author_lf" db:"author_lf"`
	Additional_Authors         string  `json:"additional_authors" db:"additional_authors"`
	ISBN                       string  `json:"isbn" db:"isbn"`
	ISBN13                     string  `json:"isbn13" db:"isbn13"`
	My_Rating                  int     `json:"my_rating" db:"my_rating"`
	Average_Rating             float64 `json:"average_rating" db:"average_rating"`
	Publisher                  string  `json:"publisher" db:"publisher"`
	Binding                    string  `json:"binding" db:"binding"`
	Number_Of_Pages            int     `json:"number_of_pages" db:"number_of_pages"`
	Year_Published             int     `json:"year_published" db:"year_published"`
	Original_Publication_Year  int     `json:"original_publication_year" db:"original_publication_year"`
	Date_Read                  string  `json:"date_read" db:"date_read"`
	Date_Added                 string  `json:"date_added" db:"date_added"`
	Bookshelves                string  `json:"bookshelves" db:"bookshelves"`
	Bookshelves_With_Positions string  `json:"bookshelves_with_positions" db:"bookshelves_with_positions"`
	Exclusive_Shelf            string  `json:"exclusive_shelf" db:"exclusive_shelf"`
	My_Review                  string  `json:"my_review" db:"my_review"`
	Spoiler                    string  `json:"spoiler" db:"spoiler"`
	Private_Notes              string  `json:"private_notes" db:"private_notes"`
	Read_Count                 int     `json:"read_count" db:"read_count"`
	Owned_Copies               int     `json:"owned_copies" db:"owned_copies"`
}

type Anime struct {
	Series_AnimeDB_ID   int     `json:"series_animedb_id" db:"series_animedb_id"`
	Series_Title        string  `json:"series_title" db:"series_title"`
	Series_Type         string  `json:"series_type" db:"series_type"`
	Series_Episodes     int     `json:"series_episodes" db:"series_episodes"`
	My_ID               int     `json:"my_id" db:"my_id"`
	My_Watched_Episodes int     `json:"my_watched_episodes" db:"my_watched_episodes"`
	My_Start_Date       string  `json:"my_start_date" db:"my_start_date"`
	My_Finish_Date      string  `json:"my_finish_date" db:"my_finish_date"`
	My_Rated            string  `json:"my_rated" db:"my_rated"`
	My_Score            float64 `json:"my_score" db:"my_score"`
	My_Storage          string  `json:"my_storage" db:"my_storage"`
	My_Storage_Value    float64 `json:"my_storage_value" db:"my_storage_value"`
	My_Status           string  `json:"my_status" db:"my_status"`
	My_Comments         string  `json:"my_comments" db:"my_comments"`
	My_Times_Watched    int     `json:"my_times_watched" db:"my_times_watched"`
	My_Rewatch_Value    int     `json:"my_rewatch_value" db:"my_rewatch_value"`
	My_Priority         int     `json:"my_priority" db:"my_priority"`
	My_Tags             string  `json:"my_tags" db:"my_tags"`
	My_Rewatching       bool    `json:"my_rewatching" db:"my_rewatching"`
	My_Rewatching_Ep    int     `json:"my_rewatching_ep" db:"my_rewatching_ep"`
	My_Discuss          bool    `json:"my_discuss" db:"my_discuss"`
	My_SNS              string  `json:"my_sns" db:"my_sns"`
	Update_On_Import    bool    `json:"update_on_import" db:"update_on_import"`
}

type Manga struct {
	Manga_MangaDB_ID     int     `json:"manga_mangadb_id" db:"manga_manga_db_id"`
	Manga_Title          string  `json:"manga_title" db:"manga_title"`
	Manga_Volumes        int     `json:"manga_volumes" db:"manga_volumes"`
	Manga_Chapters       int     `json:"manga_chapters" db:"manga_chapters"`
	My_ID                int     `json:"my_id" db:"my_id"`
	My_Read_Volumes      int     `json:"my_read_volumes" db:"my_read_volumes"`
	My_Read_Chapters     int     `json:"my_read_chapters" db:"my_read_chapters"`
	My_Start_Date        string  `json:"my_start_date" db:"my_start_date"`
	My_Finish_Date       string  `json:"my_finish_date" db:"my_finish_date"`
	My_Scanalation_Group string  `json:"my_scanalation_group" db:"my_scanalation_group"`
	My_Score             float64 `json:"my_score" db:"my_score"`
	My_Storage           string  `json:"my_storage" db:"my_storage"`
	My_Retail_Volumes    int     `json:"my_retail_volumes" db:"my_retail_volumes"`
	My_Status            string  `json:"my_status" db:"my_status"`
	My_Comments          string  `json:"my_comments" db:"my_comments"`
	My_Times_Read        int     `json:"my_times_read" db:"my_times_read"`
	My_Tags              string  `json:"my_tags" db:"my_tags"`
	My_Priority          int     `json:"my_priority" db:"my_priority"`
	My_Reread_Value      int     `json:"my_reread_value" db:"my_reread_value"`
	My_Rereading         bool    `json:"my_rereading" db:"my_rereading"`
	My_Discuss           bool    `json:"my_discuss" db:"my_discuss"`
	My_SNS               string  `json:"my_sns" db:"my_sns"`
	Update_On_Import     bool    `json:"update_on_import" db:"update_on_import"`
}

type Game struct {
	Name               string  `json:"name" db:"name"`
	Edition            string  `json:"edition" db:"edition"`
	Platform           string  `json:"platform" db:"platform"`
	Format             string  `json:"format" db:"format"`
	Region             string  `json:"region" db:"region"`
	NowPlaying         bool    `json:"now_playing" db:"now_playing"`
	Backlogged         bool    `json:"backlogged" db:"backlogged"`
	OwnershipStatus    string  `json:"ownership_status" db:"ownership_status"`
	ProgressStatus     string  `json:"progress_status" db:"progress_status"`
	Rating             float64 `json:"rating" db:"rating"`
	InitialReleaseDate string  `json:"initial_release_date" db:"initial_release_date"`
	ItemReleaseDate    string  `json:"item_release_date" db:"item_release_date"`
	AddedOn            string  `json:"added_on" db:"added_on"`
	Genre              string  `json:"genre" db:"genre"`
}
