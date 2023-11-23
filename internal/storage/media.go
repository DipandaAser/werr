package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
)

func (st *Storage) UploadMedia(reader io.Reader, userID, mediaId string) error {
	// We init an uploader
	uploader := manager.NewUploader(st.Client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(st.config.Bucket),
		Key:    aws.String(fmt.Sprintf("%s/medias/%s/%s.raw", userID, mediaId, mediaId)),
		Body:   reader,
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		return err
	}

	return nil
}
