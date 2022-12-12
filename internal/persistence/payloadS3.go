package persistence

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"test/internal/config"
	"test/internal/model"
)

type PayloadS3 struct {
	s3Client *s3.S3
}

func (p PayloadS3) StorePayload(data model.Payload) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	input := &s3.PutObjectInput{
		Body:          aws.ReadSeekCloser(bytes.NewReader(payload)),
		Bucket:        aws.String(config.Get().S3Bucket),
		Key:           aws.String(config.Get().S3Key),
		ACL:           aws.String("private"),
		ContentLength: aws.Int64(int64(len(payload))),
		ContentType:   aws.String(http.DetectContentType(payload)),
	}
	_, err = p.s3Client.PutObject(input)

	if err != nil {
		return err
	}
	return nil
}

func newPayLoadRepoS3(ctx context.Context) (repo *PayloadS3, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Get().S3Region)},
	)

	if err != nil {
		return nil, err
	}

	return &PayloadS3{s3Client: s3.New(sess)}, nil
}
