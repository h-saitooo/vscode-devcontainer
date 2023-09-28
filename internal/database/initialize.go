package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	// DB is the global database connection pool
	DB *sql.DB
)

func Initialize() error {
	var err error
	// jst, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		return fmt.Errorf("failed to load location: %w", err)
	}

	fmt.Printf("POSTGRES_DB: %s\n", os.Getenv("POSTGRES_DB"))
	fmt.Printf("POSTGRES_USER: %s\n", os.Getenv("POSTGRES_USER"))
	fmt.Printf("POSTGRES_PASSWORD: %s\n", os.Getenv("POSTGRES_PASSWORD"))

	var connStr string
	connStr += "dbname=" + os.Getenv("POSTGRES_DB")
	connStr += " user=" + os.Getenv("POSTGRES_USER")
	connStr += " password=" + os.Getenv("POSTGRES_PASSWORD")
	connStr += " host=" + os.Getenv("POSTGRES_HOSTNAME")
	connStr += " sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database %w", err)
	}

	// Set connection pool parameters
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxLifetime(time.Hour)

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")

	return nil
}

type User struct {
	name  string
	email string
}

func Query() *sql.Rows {
	query := `SELECT name, email FROM users;`

	rows, err := DB.Query(query)

	if err != nil {
		log.Fatalf("Error occured db query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.name, &user.email); err != nil {
			log.Fatalf("rows.Scan error: %v", err)
		}
		fmt.Println(user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("rows.Err occur error: %v", err)
	}

	log.Println("Success")

	return rows
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
