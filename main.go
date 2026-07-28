package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// getConfig retrieves and validates DATABASE_URL from environment
func getConfig() (string, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return "", fmt.Errorf("DATABASE_URL environment variable is not set. Please configure it in CodeBuild environment variables or buildspec.yml")
	}
	return databaseURL, nil
}

// maskSensitiveData masks credentials in connection strings for logging
func maskSensitiveData(url string) string {
	return "[CONFIGURED]"
}

func main() {
	fmt.Printf("Hello, Go %s on %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Executable: %s\n", os.Args[0])

	// Validate DATABASE_URL is configured
	databaseURL, err := getConfig()
	if err != nil {
		log.Printf("Warning: %v\n", err)
		log.Println("Application will continue, but database operations will not be available.")
		// In production, you might want to exit here:
		// log.Fatal(err)
	} else {
		// Mask sensitive information in logs
		maskedURL := maskSensitiveData(databaseURL)
		fmt.Printf("Database configured: %s\n", maskedURL)
	}
}
