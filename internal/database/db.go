package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatal("❌ Error al conectar a SQLite:", err)
	}

	createTables()
	log.Println("✅ Conectado a SQLite y base lista")
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		log.Fatal("❌ Error al crear tabla:", err)
	}
}
