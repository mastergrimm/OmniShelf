package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type Server struct {
	Port int
	Mux  *http.ServeMux
	DB   *sql.DB
}

func NewServer(port int, db *sql.DB) *Server {
	return &Server{
		Port: port,
		Mux:  http.NewServeMux(),
		DB:   db,
	}
}

func (s *Server) StartServer() error {
	// Home route
	s.Mux.HandleFunc("GET /", s.getAll)

	// Books routes
	s.Mux.HandleFunc("GET /books", s.getAllBooks)
	s.Mux.HandleFunc("POST /books", s.importBooks)
	s.Mux.HandleFunc("DELETE /books", s.clearBooks)

	// Media routes
	s.Mux.HandleFunc("GET /media", s.getAllMedia)
	s.Mux.HandleFunc("POST /media", s.importMedia)
	s.Mux.HandleFunc("DELETE /media", s.clearMedia)

	// Anime routes
	s.Mux.HandleFunc("GET /anime", s.getAllAnime)
	s.Mux.HandleFunc("POST /anime", s.importAnime)
	s.Mux.HandleFunc("DELETE /anime", s.clearAnime)

	// Manga routes
	s.Mux.HandleFunc("GET /manga", s.getAllManga)
	s.Mux.HandleFunc("POST /manga", s.importManga)
	s.Mux.HandleFunc("DELETE /manga", s.clearManga)

	// Singleplayer games routes
	s.Mux.HandleFunc("GET /singleplayer", s.getAllSingleGames)
	s.Mux.HandleFunc("POST /singleplayer", s.importSingleGames)
	s.Mux.HandleFunc("DELETE /singleplayer", s.clearSingleGames)

	// Multiplayer games routes
	s.Mux.HandleFunc("GET /multiplayer", s.getAllMultiGames)
	s.Mux.HandleFunc("POST /multiplayer", s.importMultiGames)
	s.Mux.HandleFunc("DELETE /multiplayer", s.clearMultiGames)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(s.Mux)

	addr := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Server starting on port %d\n", s.Port)

	return http.ListenAndServe(addr, handler)
}

func (s *Server) getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to OmniShelf\n\n")

	fmt.Fprint(w, "Books\n")
	s.getAllBooks(w, r)
	fmt.Fprint(w, "\n")

	fmt.Fprint(w, "Media\n")
	s.getAllMedia(w, r)
	fmt.Fprint(w, "\n")

	fmt.Fprint(w, "Anime\n")
	s.getAllAnime(w, r)
	fmt.Fprint(w, "\n")

	fmt.Fprint(w, "Manga\n")
	s.getAllManga(w, r)
	fmt.Fprint(w, "\n")

	fmt.Fprint(w, "Games\n")
	s.getAllSingleGames(w, r)
	fmt.Fprint(w, "\n")
	s.getAllMultiGames(w, r)
	fmt.Fprint(w, "\n")
}

func (s *Server) clearTable(w http.ResponseWriter, r *http.Request, tableName string) {
	tx, err := s.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to begin transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer tx.Rollback()

	result, err := tx.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	if err != nil {
		http.Error(w, "Failed to clear table: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to get rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":       fmt.Sprintf("Successfully cleared %s table", tableName),
		"rows_affected": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) clearBooks(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "books")
}

func (s *Server) clearMedia(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "media")
}

func (s *Server) clearAnime(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "anime")
}

func (s *Server) clearManga(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "manga")
}

func (s *Server) clearSingleGames(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "singleplayer")
}

func (s *Server) clearMultiGames(w http.ResponseWriter, r *http.Request) {
	s.clearTable(w, r, "multiplayer")
}
