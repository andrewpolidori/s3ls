package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func createAwsSession(awsRegion string) (*session.Session, error) {
	awsConfig := defaults.Get().Config.WithRegion(awsRegion)

	_, err := awsConfig.Credentials.Get()
	if err != nil {
		return nil, fmt.Errorf("Error finding AWS credentials (did you set the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables?): %v", err)
	}

	return session.New( awsConfig ), nil
}

// Create a client the SDK can use to perform operations on the EC2 service.
func createS3Client(session *session.Session) *s3.S3 {
	return s3.New(session)
}

func listFiles(s3Session *s3.S3, bucketName string, prefix string) ([]*s3.Object, error) {
	output, err := s3Session.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to list files from S3 Bucket: %v", err)
	}

	return output.Contents, nil
}