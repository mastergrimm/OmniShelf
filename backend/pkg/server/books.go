package server

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/mastergrimm/OmniShelf/pkg/models"
)

func (s *Server) getAllBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query("SELECT * FROM books")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var books []models.Book
	bookType := reflect.TypeOf(models.Book{})

	values := make([]interface{}, bookType.NumField())
	for i := range values {
		values[i] = reflect.New(bookType.Field(i).Type).Interface()
	}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		book := models.Book{}
		for i := 0; i < bookType.NumField(); i++ {
			field := reflect.ValueOf(&book).Elem().Field(i)
			val := reflect.ValueOf(values[i]).Elem()
			field.Set(val)
		}

		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func printStruct(w http.ResponseWriter, v interface{}) {
	val := reflect.ValueOf(v)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Fprintf(w, "%s: %v\n", fieldName, fieldValue)
	}
}

func (s *Server) importBooks(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var inserted, updated int

	// Skip the header row
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

		myRating, _ := strconv.Atoi(record[7])
		averageRating, _ := strconv.ParseFloat(record[8], 64)
		numberOfPages, _ := strconv.Atoi(record[11])
		yearPublished, _ := strconv.Atoi(record[12])
		originalPublicationYear, _ := strconv.Atoi(record[13])
		readCount, _ := strconv.Atoi(record[22])
		ownedCopies, _ := strconv.Atoi(record[23])

		book := models.Book{
			Book_ID:                    record[0],
			Title:                      record[1],
			Author:                     record[2],
			Author_LF:                  record[3],
			Additional_Authors:         record[4],
			ISBN:                       record[5],
			ISBN13:                     record[6],
			My_Rating:                  myRating,
			Average_Rating:             averageRating,
			Publisher:                  record[9],
			Binding:                    record[10],
			Number_Of_Pages:            numberOfPages,
			Year_Published:             yearPublished,
			Original_Publication_Year:  originalPublicationYear,
			Date_Read:                  record[14],
			Date_Added:                 record[15],
			Bookshelves:                record[16],
			Bookshelves_With_Positions: record[17],
			Exclusive_Shelf:            record[18],
			My_Review:                  record[19],
			Spoiler:                    record[20],
			Private_Notes:              record[21],
			Read_Count:                 readCount,
			Owned_Copies:               ownedCopies,
		}

		var exists bool
		err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE Book_ID = ?)", book.Book_ID).Scan(&exists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if exists {
			_, err = s.DB.Exec(`UPDATE books SET
				Title = ?, Author = ?, Author_LF = ?, Additional_Authors = ?,
				ISBN = ?, ISBN13 = ?, My_Rating = ?, Average_Rating = ?,
				Publisher = ?, Binding = ?, Number_Of_Pages = ?, Year_Published = ?,
				Original_Publication_Year = ?, Date_Read = ?, Date_Added = ?,
				Bookshelves = ?, Bookshelves_With_Positions = ?, Exclusive_Shelf = ?,
				My_Review = ?, Spoiler = ?, Private_Notes = ?, Read_Count = ?,
				Owned_Copies = ?
				WHERE Book_ID = ?`,
				book.Title, book.Author, book.Author_LF, book.Additional_Authors,
				book.ISBN, book.ISBN13, book.My_Rating, book.Average_Rating,
				book.Publisher, book.Binding, book.Number_Of_Pages, book.Year_Published,
				book.Original_Publication_Year, book.Date_Read, book.Date_Added,
				book.Bookshelves, book.Bookshelves_With_Positions, book.Exclusive_Shelf,
				book.My_Review, book.Spoiler, book.Private_Notes, book.Read_Count,
				book.Owned_Copies, book.Book_ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updated++
		} else {
			_, err = s.DB.Exec(`INSERT INTO books (
				Book_ID, Title, Author, Author_LF, Additional_Authors,
				ISBN, ISBN13, My_Rating, Average_Rating, Publisher, Binding,
				Number_Of_Pages, Year_Published, Original_Publication_Year,
				Date_Read, Date_Added, Bookshelves, Bookshelves_With_Positions,
				Exclusive_Shelf, My_Review, Spoiler, Private_Notes, Read_Count,
				Owned_Copies
				) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				book.Book_ID, book.Title, book.Author, book.Author_LF, book.Additional_Authors,
				book.ISBN, book.ISBN13, book.My_Rating, book.Average_Rating, book.Publisher,
				book.Binding, book.Number_Of_Pages, book.Year_Published, book.Original_Publication_Year,
				book.Date_Read, book.Date_Added, book.Bookshelves, book.Bookshelves_With_Positions,
				book.Exclusive_Shelf, book.My_Review, book.Spoiler, book.Private_Notes,
				book.Read_Count, book.Owned_Copies)
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
