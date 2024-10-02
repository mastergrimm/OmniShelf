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
		mangaVolumes, _ := strconv.Atoi(record[2])
		mangaChapters, _ := strconv.Atoi(record[3])
		myID, _ := strconv.Atoi(record[4])
		myReadVolumes, _ := strconv.Atoi(record[5])
		myReadChapters, _ := strconv.Atoi(record[6])
		myScore, _ := strconv.ParseFloat(record[10], 64)
		myRetailVolumes, _ := strconv.Atoi(record[12])
		myTimesRead, _ := strconv.Atoi(record[15])
		myPriority, _ := strconv.Atoi(record[17])
		myRereadValue, _ := strconv.Atoi(record[18])
		myRereading, _ := strconv.ParseBool(record[19])
		myDiscuss, _ := strconv.ParseBool(record[20])
		updateOnImport, _ := strconv.ParseBool(record[22])

		manga := models.Manga{
			Manga_MangaDB_ID:     mangaMangaDBID,
			Manga_Title:          record[1],
			Manga_Volumes:        mangaVolumes,
			Manga_Chapters:       mangaChapters,
			My_ID:                myID,
			My_Read_Volumes:      myReadVolumes,
			My_Read_Chapters:     myReadChapters,
			My_Start_Date:        record[7],
			My_Finish_Date:       record[8],
			My_Scanalation_Group: record[9],
			My_Score:             myScore,
			My_Storage:           record[11],
			My_Retail_Volumes:    myRetailVolumes,
			My_Status:            record[13],
			My_Comments:          record[14],
			My_Times_Read:        myTimesRead,
			My_Tags:              record[16],
			My_Priority:          myPriority,
			My_Reread_Value:      myRereadValue,
			My_Rereading:         myRereading,
			My_Discuss:           myDiscuss,
			My_SNS:               record[21],
			Update_On_Import:     updateOnImport,
		}

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM manga WHERE manga_mangadb_id = ?)", manga.Manga_MangaDB_ID).Scan(&exists)
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
				manga.Manga_Title, manga.Manga_Volumes, manga.Manga_Chapters, manga.My_ID,
				manga.My_Read_Volumes, manga.My_Read_Chapters, manga.My_Start_Date, manga.My_Finish_Date,
				manga.My_Scanalation_Group, manga.My_Score, manga.My_Storage, manga.My_Retail_Volumes,
				manga.My_Status, manga.My_Comments, manga.My_Times_Read, manga.My_Tags,
				manga.My_Priority, manga.My_Reread_Value, manga.My_Rereading, manga.My_Discuss,
				manga.My_SNS, manga.Update_On_Import, manga.Manga_MangaDB_ID)
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
				manga.Manga_MangaDB_ID, manga.Manga_Title, manga.Manga_Volumes, manga.Manga_Chapters, manga.My_ID,
				manga.My_Read_Volumes, manga.My_Read_Chapters, manga.My_Start_Date, manga.My_Finish_Date,
				manga.My_Scanalation_Group, manga.My_Score, manga.My_Storage, manga.My_Retail_Volumes,
				manga.My_Status, manga.My_Comments, manga.My_Times_Read, manga.My_Tags,
				manga.My_Priority, manga.My_Reread_Value, manga.My_Rereading, manga.My_Discuss,
				manga.My_SNS, manga.Update_On_Import)
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
