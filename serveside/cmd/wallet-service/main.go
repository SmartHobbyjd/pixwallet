// backend/cmd/wallet-service/main.go

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/your_username/wallet-service/internal/pixpb"
	"github.com/your_username/wallet-service/internal/wallet"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := wallet.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database")
		return
	}
	defer db.Close()

	// Create a new WalletService instance
	walletService := wallet.NewWalletService(db)

	// Start gRPC server
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pixpb.RegisterWalletServiceServer(grpcServer, walletService)

	go func() {
		log.Printf("Starting gRPC server on port %d", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Handle graceful shutdown
	waitForShutdown()
}

func waitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	fmt.Println("\nShutting down the server...")

	// Gracefully stop the gRPC server
	grpcServer.GracefulStop()

	fmt.Println("Server gracefully stopped")
}
