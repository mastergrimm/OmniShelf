package server

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"io"
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
		var id sql.NullInt64
		err := rows.Scan(&id, &game.Game, &game.URL, &game.Rating, &game.Category,
			&game.Release_Date, &game.Platforms, &game.Genres, &game.Themes, &game.Companies, &game.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if id.Valid {
			game.ID = int(id.Int64)
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

func (s *Server) importGames(w http.ResponseWriter, r *http.Request, tableName string) error {
	file, _, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var inserted, updated int

	if _, err := reader.Read(); err != nil {
		return err
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO ` + tableName + ` (
		id, game, url, rating, category, release_date, platforms,
		genres, themes, companies, description
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	updateStmt, err := tx.Prepare(`UPDATE ` + tableName + ` SET
		game = ?, url = ?, rating = ?, category = ?, release_date = ?, platforms = ?,
		genres = ?, themes = ?, companies = ?, description = ?
		WHERE id = ?`)
	if err != nil {
		return err
	}
	defer updateStmt.Close()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		id, _ := strconv.Atoi(record[0])
		rating, _ := strconv.ParseFloat(record[3], 64)
		game := models.Game{
			ID:           id,
			Game:         record[1],
			URL:          record[2],
			Rating:       rating,
			Category:     record[4],
			Release_Date: record[5],
			Platforms:    record[6],
			Genres:       record[7],
			Themes:       record[8],
			Companies:    record[9],
			Description:  record[10],
		}

		var existingID int
		err = tx.QueryRow("SELECT id FROM "+tableName+" WHERE id = ?", game.ID).Scan(&existingID)
		if err == nil {
			_, err = updateStmt.Exec(
				game.Game, game.URL, game.Rating, game.Category, game.Release_Date, game.Platforms,
				game.Genres, game.Themes, game.Companies, game.Description, game.ID)
			if err != nil {
				return err
			}
			updated++
		} else if err == sql.ErrNoRows {
			_, err := stmt.Exec(
				game.ID, game.Game, game.URL, game.Rating, game.Category, game.Release_Date,
				game.Platforms, game.Genres, game.Themes, game.Companies, game.Description)
			if err != nil {
				return err
			}
			inserted++
		} else {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"inserted": inserted,
		"updated":  updated,
	})

	return nil
}

func (s *Server) importSingleGames(w http.ResponseWriter, r *http.Request) {
	err := s.importGames(w, r, "singleplayer")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) importMultiGames(w http.ResponseWriter, r *http.Request) {
	err := s.importGames(w, r, "multiplayer")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
