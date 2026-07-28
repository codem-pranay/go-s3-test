package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	fmt.Printf("Hello, Go %s on %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Executable: %s\n", os.Args[0])

	// Initialize AWS session with region from environment variable or default
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "ap-south-1" // Default region
		fmt.Printf("AWS_REGION not set, using default: %s\n", region)
	} else {
		fmt.Printf("Using AWS_REGION: %s\n", region)
	}

	// Create AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Printf("Error creating AWS session: %v\n", err)
		return
	}

	fmt.Printf("AWS S3 client initialized successfully in region: %s\n", *sess.Config.Region)
}
