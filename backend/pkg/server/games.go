package server

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/mastergrimm/OmniShelf/pkg/models"
)

func (s *Server) getAllGames(w http.ResponseWriter, r *http.Request, tableName string) {
	rows, err := s.DB.Query("SELECT * FROM " + tableName)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.Name, &game.Edition, &game.Platform, &game.Format,
			&game.Region, &game.NowPlaying, &game.Backlogged, &game.OwnershipStatus,
			&game.ProgressStatus, &game.Rating, &game.InitialReleaseDate, &game.ItemReleaseDate,
			&game.AddedOn, &game.Genre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		games = append(games, game)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func (s *Server) getAllSingleGames(w http.ResponseWriter, r *http.Request) {
	s.getAllGames(w, r, "singleplayer")
}

func (s *Server) getAllMultiGames(w http.ResponseWriter, r *http.Request) {
	s.getAllGames(w, r, "multiplayer")
}

func (s *Server) importGames(w http.ResponseWriter, r *http.Request, tableName string) {
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

		nowPlaying, _ := strconv.ParseBool(record[5])
		backlogged, _ := strconv.ParseBool(record[6])
		rating, _ := strconv.ParseFloat(record[9], 64)

		game := models.Game{
			Name:               record[0],
			Edition:            record[1],
			Platform:           record[2],
			Format:             record[3],
			Region:             record[4],
			NowPlaying:         nowPlaying,
			Backlogged:         backlogged,
			OwnershipStatus:    record[7],
			ProgressStatus:     record[8],
			Rating:             rating,
			InitialReleaseDate: record[10],
			ItemReleaseDate:    record[11],
			AddedOn:            record[12],
			Genre:              record[13],
		}

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM "+tableName+" WHERE name = ? AND platform = ?)", game.Name, game.Platform).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if exists {
			_, err = s.DB.Exec(`UPDATE `+tableName+` SET
				edition = ?, format = ?, region = ?, now_playing = ?,
				backlogged = ?, ownership_status = ?, progress_status = ?, rating = ?,
				initial_release_date = ?, item_release_date = ?, added_on = ?, genre = ?
				WHERE name = ? AND platform = ?`,
				game.Edition, game.Format, game.Region, game.NowPlaying, game.Backlogged,
				game.OwnershipStatus, game.ProgressStatus, game.Rating, game.InitialReleaseDate,
				game.ItemReleaseDate, game.AddedOn, game.Genre, game.Name, game.Platform)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updated++
		} else {
			_, err = s.DB.Exec(`INSERT INTO `+tableName+` (
				name, edition, platform, format, region, now_playing, backlogged,
				ownership_status, progress_status, rating, initial_release_date,
				item_release_date, added_on, genre
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				game.Name, game.Edition, game.Platform, game.Format, game.Region,
				game.NowPlaying, game.Backlogged, game.OwnershipStatus, game.ProgressStatus,
				game.Rating, game.InitialReleaseDate, game.ItemReleaseDate, game.AddedOn,
				game.Genre)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			inserted++
		}
	}

	log.Printf("Import completed. Inserted: %d, Updated: %d", inserted, updated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"inserted": inserted,
		"updated":  updated,
	})
}

func (s *Server) importSingleGames(w http.ResponseWriter, r *http.Request) {
	s.importGames(w, r, "singleplayer")
}

func (s *Server) importMultiGames(w http.ResponseWriter, r *http.Request) {
	s.importGames(w, r, "multiplayer")
}

