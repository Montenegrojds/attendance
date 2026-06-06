package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func Init() (*sql.DB, error) {
	godotenv.Load()
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			DNI TEXT PRIMARY KEY,
			Name TEXT,
			Email TEXT,
			StudentCode TEXT
		);
		
		CREATE TABLE IF NOT EXISTS users(
			ID	SERIAL PRIMARY KEY,
			Name TEXT,
			Email	TEXT,
			Password	TEXT,
			Role	TEXT 
		);

		CREATE TABLE IF NOT EXISTS courses(
		ID	SERIAL PRIMARY KEY,
		Name TEXT,
		Code TEXT,
		Group_number INT,
		Professor TEXT
		);

		CREATE TABLE IF NOT EXISTS sessions(
		ID SERIAL PRIMARY KEY,
		CourseID INT REFERENCES courses(id),
		Date	DATE,
		Start_time TIME,
		End_time	TIME

		);

		CREATE TABLE IF NOT EXISTS attendances(
		ID SERIAL PRIMARY KEY,
		SessionID INT REFERENCES sessions(id),
		Present BOOL,
		TakenBy INT REFERENCES users(id),
		DNI	TEXT REFERENCES students(dni)

		);

		CREATE TABLE IF NOT EXISTS points(
		ID SERIAL PRIMARY KEY,
		StudentDNI	TEXT REFERENCES students(dni),
		Amount	INT,
		SessionID INT REFERENCES sessions(id)
		);
	
	`)

	if err != nil {
		return err
	}
	return nil
}
