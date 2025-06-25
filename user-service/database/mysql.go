package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rizalherniawan/99-backend-test/user-service/config"
)

func Init() *sql.DB {
	// Load environment variables

	username := config.GetEnv("DB_USERNAME")
	password := config.GetEnv("DB_PASSWORD")
	name := config.GetEnv("DB_NAME")
	port := config.GetEnv("DB_PORT")
	host := config.GetEnv("DB_HOST")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Println("Waiting for DB to be ready:", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to DB after retries: %v", err)
	}

	log.Println("Connected to DB successfully")
	return db
}

func RunMigration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Failed to run migration....", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create migration instance: %v", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(fmt.Sprintf("failed to apply migrations: %v", err))
	}

	log.Println("Migration run successfully")
}
