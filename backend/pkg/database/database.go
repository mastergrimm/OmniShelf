package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/mastergrimm/OmniShelf/pkg/models"
	_ "modernc.org/sqlite"
)

func InitDB(dbPath string) (*sql.DB, error) {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(dbPath), 0755)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	err := CreateTableFromStruct(db, "media", models.Media{})
	if err != nil {
		return err
	}

	err = CreateTableFromStruct(db, "books", models.Book{})
	if err != nil {
		return err
	}

	err = CreateTableFromStruct(db, "anime", models.Anime{})
	if err != nil {
		return err
	}

	err = CreateTableFromStruct(db, "manga", models.Manga{})
	if err != nil {
		return err
	}

	err = CreateTableFromStruct(db, "singleplayer", models.Game{})
	if err != nil {
		return err
	}

	err = CreateTableFromStruct(db, "multiplayer", models.Game{})
	if err != nil {
		return err
	}

	return nil
}

func CreateTableFromStruct(db *sql.DB, tableName string, s interface{}) error {
	createTableSQL := generateCreateTableSQL(tableName, s)
	_, err := db.Exec(createTableSQL)
	return err
}

func generateCreateTableSQL(tableName string, s interface{}) string {
	t := reflect.TypeOf(s)
	var columns []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnName := field.Name
		columnType := getSQLType(field.Type)

		columns = append(columns, fmt.Sprintf("%s %s", columnName, columnType))
	}

	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n\t%s\n);",
		tableName, strings.Join(columns, ",\n\t"))

	return createTableSQL
}

func getSQLType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "INTEGER"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	case reflect.Bool:
		return "BOOLEAN"
	case reflect.String:
		return "TEXT"
	default:
		return "TEXT"
	}
}
