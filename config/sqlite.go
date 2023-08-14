package config

import (
	"os"

	"github.com/viniihds/api-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Criando DB e conectando
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("Erro de conexao com o s: %v", err)
		return nil, err
	}
	// Migrando o Schema
	err = db.AutoMigrate(&schemas.Vaga{})
	if err != nil {
		logger.Errf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}