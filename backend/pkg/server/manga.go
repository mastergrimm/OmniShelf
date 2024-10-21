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

func (s *Server) getAllManga(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query("SELECT * FROM manga")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()
	var mangaList []models.Manga
	mangaType := reflect.TypeOf(models.Manga{})

	values := make([]interface{}, mangaType.NumField())
	for i := range values {
		values[i] = reflect.New(mangaType.Field(i).Type).Interface()
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		manga := models.Manga{}
		for i := 0; i < mangaType.NumField(); i++ {
			field := reflect.ValueOf(&manga).Elem().Field(i)
			val := reflect.ValueOf(values[i]).Elem()
			field.Set(val)
		}

		mangaList = append(mangaList, manga)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mangaList)
}

func (s *Server) importManga(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var inserted, updated int

	// Read and discard the header row
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

		mangaMangaDBID, _ := strconv.Atoi(record[0])
		myRating, _ := strconv.ParseFloat(record[10], 64)
		myTimesRead, _ := strconv.Atoi(record[15])

		// Convert priority string to integer
		var priorityInt int
		switch record[17] {
		case "Low":
			priorityInt = 1
		case "Medium":
			priorityInt = 2
		case "High":
			priorityInt = 3
		default:
			priorityInt = 0
		}

		myRereadValue, _ := strconv.Atoi(record[18])
		myRereading := record[19] == "YES"
		myDiscuss := record[20] == "YES"
		updateOnImport := record[22] == "1"

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM manga WHERE manga_mangadb_id = ?)", mangaMangaDBID).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if exists {
			_, err = s.DB.Exec(`UPDATE manga SET
                manga_title = ?, manga_volumes = ?, manga_chapters = ?, my_id = ?,
                my_read_volumes = ?, my_read_chapters = ?, my_start_date = ?, my_finish_date = ?,
                my_scanalation_group = ?, my_score = ?, my_storage = ?, my_retail_volumes = ?,
                my_status = ?, my_comments = ?, my_times_read = ?, my_tags = ?,
                my_priority = ?, my_reread_value = ?, my_rereading = ?, my_discuss = ?,
                my_sns = ?, update_on_import = ?
                WHERE manga_mangadb_id = ?`,
				record[1], record[2], record[3], record[4], record[5], record[6],
				record[7], record[8], record[9], myRating, record[11], record[12],
				record[13], record[14], myTimesRead, record[16],
				priorityInt, myRereadValue, myRereading, myDiscuss,
				record[21], updateOnImport, mangaMangaDBID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updated++
		} else {
			_, err = s.DB.Exec(`INSERT INTO manga (
                manga_mangadb_id, manga_title, manga_volumes, manga_chapters, my_id,
                my_read_volumes, my_read_chapters, my_start_date, my_finish_date,
                my_scanalation_group, my_score, my_storage, my_retail_volumes,
                my_status, my_comments, my_times_read, my_tags,
                my_priority, my_reread_value, my_rereading, my_discuss,
                my_sns, update_on_import
                ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				mangaMangaDBID, record[1], record[2], record[3], record[4], record[5],
				record[6], record[7], record[8], record[9], myRating, record[11],
				record[12], record[13], record[14], myTimesRead, record[16],
				priorityInt, myRereadValue, myRereading, myDiscuss,
				record[21], updateOnImport)
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
