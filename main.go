package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Hello, Go %s on %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Executable: %s\n", os.Args[0])

	// Check for DATABASE_URL environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("Warning: DATABASE_URL environment variable is not set")
		os.Exit(1)
	}

	fmt.Printf("Database URL configured: %s\n", maskDatabaseURL(dbURL))
	fmt.Println("Application initialized successfully")
}

func maskDatabaseURL(url string) string {
	// Mask the password in the database URL for logging purposes
	if len(url) > 20 {
		return url[:10] + "..." + url[len(url)-10:]
	}
	return "***"
}
