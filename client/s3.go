package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/slham/toolbelt/l"
	"time"
)

var sess *session.Session

// Create aws session
func InitializeSession() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
}

// Returns S3 object from given bucket
func GetS3Object(bucket, key string) ([]byte, error) {
	// Create an downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)
	l.Debug(nil, "downloader created")

	// Store object in buffer
	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		l.Error(nil, "could not fetch object from S3: %v", err)
		return make([]byte, 0), err
	}

	return buff.Bytes(), nil
}

// Return the Key of the latest S3 object from the given bucket with the provided prefix
func GetLatestS3Key(bucket, prefix string) (string, error) {
	cli := s3.New(sess)
	l.Debug(nil, "S3 client created")

	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}
	objs, err := cli.ListObjectsV2(params)
	if err != nil || objs == nil {
		l.Error(nil, "unable to list S3 bucket objects %s::%s", bucket, prefix)
		return "", err
	}

	var holder *time.Time
	var key string

	// Find most recent object and return its Key
	for _, obj := range objs.Contents {
		if holder == nil {
			key = *obj.Key
			holder = obj.LastModified
		} else {
			if obj.LastModified.After(*holder) {
				key = *obj.Key
				holder = obj.LastModified
			}
		}
	}

	return key, nil
}
