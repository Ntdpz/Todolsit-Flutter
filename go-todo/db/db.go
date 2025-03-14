package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// ดึงค่าตัวแปรจาก .env
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "admin"),
		getEnv("DB_PASSWORD", "secret"),
		getEnv("DB_NAME", "tododb"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("ไม่สามารถเชื่อมต่อ Database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database ไม่ตอบสนอง:", err)
	}

	fmt.Println("✅ Database เชื่อมต่อสำเร็จ!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
