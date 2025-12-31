package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start all servers in goroutines
	go setupNetHTTPServer()
	go setupGinServer()
	go setupEchoServer()
	go setupFiberServer()
	go setupChiServer()

	fmt.Println("All API servers started!")
	fmt.Println("Net/HTTP: http://localhost:8080")
	fmt.Println("Gin:      http://localhost:8081")
	fmt.Println("Echo:     http://localhost:8082")
	fmt.Println("Fiber:    http://localhost:8083")
	fmt.Println("Chi:      http://localhost:8084")
	fmt.Println("\nPress Ctrl+C to stop all servers")

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down servers...")
}

