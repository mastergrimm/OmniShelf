package server

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/mastergrimm/OmniShelf/pkg/models"
)

func (s *Server) getAllMedia(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query("SELECT * FROM media")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var movies []models.Media
	var tvShows []models.Media

	mediaType := reflect.TypeOf(models.Media{})
	values := make([]interface{}, mediaType.NumField())
	for i := range values {
		values[i] = reflect.New(mediaType.Field(i).Type).Interface()
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		media := models.Media{}
		for i := 0; i < mediaType.NumField(); i++ {
			field := reflect.ValueOf(&media).Elem().Field(i)
			val := reflect.ValueOf(values[i]).Elem()
			field.Set(val)
		}

		if media.Title_Type == "Movie" {
			movies = append(movies, media)
		} else if media.Title_Type == "TV Series" {
			tvShows = append(tvShows, media)
		}
	}

	response := map[string]interface{}{
		"movies":  movies,
		"tvShows": tvShows,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) importMedia(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var inserted, updated int

	if _, err := reader.Read(); err != nil {
		http.Error(w, "Failed to read CSV header: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		media := models.Media{
			Const:          record[0],
			Your_Rating:    record[1],
			Date_Rated:     record[2],
			Title:          record[3],
			Original_Title: record[4],
			URL:            record[5],
			Title_Type:     record[6],
			IMDb_Rating:    record[7],
			Runtime:        record[8],
			Year:           record[9],
			Genres:         record[10],
			Num_Votes:      record[11],
			Release_Date:   record[12],
			Directors:      record[13],
		}

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM media WHERE Const = ?)", media.Const).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if exists {
			_, err = s.DB.Exec(`UPDATE media SET
				Your_Rating = ?, Date_Rated = ?, Title = ?, Original_Title = ?,
				URL = ?, Title_Type = ?, IMDb_Rating = ?, Runtime = ?, Year = ?,
				Genres = ?, Num_Votes = ?, Release_Date = ?, Directors = ?
				WHERE Const = ?`,
				media.Your_Rating, media.Date_Rated, media.Title, media.Original_Title,
				media.URL, media.Title_Type, media.IMDb_Rating, media.Runtime, media.Year,
				media.Genres, media.Num_Votes, media.Release_Date, media.Directors,
				media.Const)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updated++
		} else {
			_, err = s.DB.Exec(`INSERT INTO media (
				Const, Your_Rating, Date_Rated, Title, Original_Title,
				URL, Title_Type, IMDb_Rating, Runtime, Year,
				Genres, Num_Votes, Release_Date, Directors
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				media.Const, media.Your_Rating, media.Date_Rated, media.Title, media.Original_Title,
				media.URL, media.Title_Type, media.IMDb_Rating, media.Runtime, media.Year,
				media.Genres, media.Num_Votes, media.Release_Date, media.Directors)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			inserted++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"inserted": inserted,
		"updated":  updated,
	})
}
