// backend/internal/wallet/wallet.go

package wallet

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/your_username/wallet-service/internal/pixpb"
)

// WalletService implements the gRPC service
type WalletService struct {
	db *gorm.DB
}

// NewWalletService creates a new WalletService instance
func NewWalletService(db *gorm.DB) *WalletService {
	return &WalletService{db: db}
}

// SendPix handles the SendPix gRPC request
func (s *WalletService) SendPix(ctx context.Context, req *pixpb.SendPixRequest) (*pixpb.SendPixResponse, error) {
	// Your implementation for sending Pix transactions

	// Example: Store the transaction in the database
	transaction := Transaction{
		Amount:    req.Amount,
		Recipient: req.Recipient,
		Status:    "completed", // You can implement status handling logic
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(&transaction).Error; err != nil {
		log.Printf("Error storing transaction: %v", err)
		return nil, fmt.Errorf("failed to process transaction")
	}

	// Return transaction details
	return &pixpb.SendPixResponse{
		Id:        fmt.Sprint(transaction.ID),
		Amount:    transaction.Amount,
		Recipient: transaction.Recipient,
		Status:    transaction.Status,
	}, nil
}
