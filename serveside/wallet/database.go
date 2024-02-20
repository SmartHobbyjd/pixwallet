// backend/internal/wallet/database.go

package wallet

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Transaction represents a Pix transaction
type Transaction struct {
	gorm.Model
	Amount    string
	Recipient string
	Status    string
}

// ConnectToDatabase initializes the connection to the PostgreSQL database
func ConnectToDatabase() (*gorm.DB, error) {
	// Replace with your PostgreSQL connection details
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=your_username dbname=your_dbname password=your_password sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	// Auto-migrate the Transaction model
	db.AutoMigrate(&Transaction{})

	log.Println("Connected to database")
	return db, nil
}
