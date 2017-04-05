package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	logger := NewLogger("s3dl")

	args := os.Args[1:]

	if err := validateArgs(args); err != nil {
		logger.Printf("%v", err)
		os.Exit(1)
	}

	region := args[0]
	s3Path := args[1]
	localPath := args[2]

	s3Bucket := getS3BucketName(s3Path)
	s3Prefix := getS3Key(s3Path)

	logger.Printf("Download first file from s3://%s/%s/ in %s to %s\n", s3Bucket, s3Prefix, region, localPath)

	session, err := createAwsSession(region)
	if err != nil {
		logger.Printf("Failed to create AWS Session: %v", err)
		os.Exit(1)
	}

	s3 := createS3Client(session)

	s3Objects, err := listFiles(s3, s3Bucket, s3Prefix)
	if err != nil {
		logger.Printf("%v", err)
		os.Exit(1)
	}

	firstS3ObjectKey := *s3Objects[0].Key

	downloadFile(s3, s3Bucket, firstS3ObjectKey, localPath)
}

func validateArgs(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("Error: you must have exactly 3 command line arguments.\n%s", getUsageText())
	}

	return nil
}

func getUsageText() string {
	return `Usage: s3dl <bucket-region> s3://<bucket-name>/<key> <local-path>`
}

func getS3BucketName(s3Path string) string {
	s3Path = strings.Trim(s3Path, "s3://")
	s3PathComponents := strings.Split(s3Path, "/")
	return s3PathComponents[0]
}

func getS3Key(s3Path string) string {
	s3Path = strings.Trim(s3Path, "s3://")
	s3PathComponents := strings.Split(s3Path, "/")
	return strings.Join(s3PathComponents[1:], "/")
}