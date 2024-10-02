package main

import (
	"log"
	"path/filepath"

	"github.com/mastergrimm/OmniShelf/pkg/database"
	"github.com/mastergrimm/OmniShelf/pkg/server"
)

func main() {
	dbPath := filepath.Join("data", "collections.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(8080, db)
	log.Fatal(s.StartServer())
}

