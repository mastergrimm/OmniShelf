package server

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/mastergrimm/OmniShelf/pkg/models"
)

func (s *Server) getAllAnime(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query("SELECT * FROM anime")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var animeList []models.Anime
	animeType := reflect.TypeOf(models.Anime{})
	values := make([]interface{}, animeType.NumField())
	for i := range values {
		values[i] = reflect.New(animeType.Field(i).Type).Interface()
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		anime := models.Anime{}
		for i := 0; i < animeType.NumField(); i++ {
			field := reflect.ValueOf(&anime).Elem().Field(i)
			val := reflect.ValueOf(values[i]).Elem()
			field.Set(val)
		}
		animeList = append(animeList, anime)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animeList)
}

func (s *Server) importAnime(w http.ResponseWriter, r *http.Request) {
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

		seriesAnimeDBID, _ := strconv.Atoi(record[0])
		seriesEpisodes, _ := strconv.Atoi(record[3])
		myID, _ := strconv.Atoi(record[4])
		myWatchedEpisodes, _ := strconv.Atoi(record[5])
		myScore, _ := strconv.ParseFloat(record[9], 64)
		myStorageValue, _ := strconv.ParseFloat(record[11], 64)
		myTimesWatched, _ := strconv.Atoi(record[14])
		myRewatchValue, _ := strconv.Atoi(record[15])
		myPriority, _ := strconv.Atoi(record[16])
		myRewatching, _ := strconv.ParseBool(record[18])
		myRewatchingEp, _ := strconv.Atoi(record[19])
		myDiscuss, _ := strconv.ParseBool(record[20])
		updateOnImport, _ := strconv.ParseBool(record[22])

		anime := models.Anime{
			Series_AnimeDB_ID:   seriesAnimeDBID,
			Series_Title:        record[1],
			Series_Type:         record[2],
			Series_Episodes:     seriesEpisodes,
			My_ID:               myID,
			My_Watched_Episodes: myWatchedEpisodes,
			My_Start_Date:       record[6],
			My_Finish_Date:      record[7],
			My_Rated:            record[8],
			My_Score:            myScore,
			My_Storage:          record[10],
			My_Storage_Value:    myStorageValue,
			My_Status:           record[12],
			My_Comments:         record[13],
			My_Times_Watched:    myTimesWatched,
			My_Rewatch_Value:    myRewatchValue,
			My_Priority:         myPriority,
			My_Tags:             record[17],
			My_Rewatching:       myRewatching,
			My_Rewatching_Ep:    myRewatchingEp,
			My_Discuss:          myDiscuss,
			My_SNS:              record[21],
			Update_On_Import:    updateOnImport,
		}

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM anime WHERE series_animedb_id = ?)", anime.Series_AnimeDB_ID).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if exists {
			_, err = s.DB.Exec(`UPDATE anime SET
				series_title = ?, series_type = ?, series_episodes = ?, my_id = ?,
				my_watched_episodes = ?, my_start_date = ?, my_finish_date = ?, my_rated = ?,
				my_score = ?, my_storage = ?, my_storage_value = ?, my_status = ?,
				my_comments = ?, my_times_watched = ?, my_rewatch_value = ?, my_priority = ?,
				my_tags = ?, my_rewatching = ?, my_rewatching_ep = ?, my_discuss = ?,
				my_sns = ?, update_on_import = ?
				WHERE series_animedb_id = ?`,
				anime.Series_Title, anime.Series_Type, anime.Series_Episodes, anime.My_ID,
				anime.My_Watched_Episodes, anime.My_Start_Date, anime.My_Finish_Date, anime.My_Rated,
				anime.My_Score, anime.My_Storage, anime.My_Storage_Value, anime.My_Status,
				anime.My_Comments, anime.My_Times_Watched, anime.My_Rewatch_Value, anime.My_Priority,
				anime.My_Tags, anime.My_Rewatching, anime.My_Rewatching_Ep, anime.My_Discuss,
				anime.My_SNS, anime.Update_On_Import, anime.Series_AnimeDB_ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updated++
		} else {
			_, err = s.DB.Exec(`INSERT INTO anime (
				series_animedb_id, series_title, series_type, series_episodes, my_id,
				my_watched_episodes, my_start_date, my_finish_date, my_rated, my_score,
				my_storage, my_storage_value, my_status, my_comments, my_times_watched,
				my_rewatch_value, my_priority, my_tags, my_rewatching, my_rewatching_ep,
				my_discuss, my_sns, update_on_import
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				anime.Series_AnimeDB_ID, anime.Series_Title, anime.Series_Type, anime.Series_Episodes, anime.My_ID,
				anime.My_Watched_Episodes, anime.My_Start_Date, anime.My_Finish_Date, anime.My_Rated, anime.My_Score,
				anime.My_Storage, anime.My_Storage_Value, anime.My_Status, anime.My_Comments, anime.My_Times_Watched,
				anime.My_Rewatch_Value, anime.My_Priority, anime.My_Tags, anime.My_Rewatching, anime.My_Rewatching_Ep,
				anime.My_Discuss, anime.My_SNS, anime.Update_On_Import)
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
