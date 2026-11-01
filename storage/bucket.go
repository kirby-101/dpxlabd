package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Bucket struct {
	name   string
	client *s3.Client
}

func (b *Bucket) ListObjects(ctx context.Context) ([]types.Object, error) {
	lsO, err := b.client.ListObjectsV2(
		ctx,
		&s3.ListObjectsV2Input{
			Bucket: aws.String(b.name),
		},
	)
	if err != nil {
		return nil, err
	}

	return lsO.Contents, nil
}

func (b *Bucket) CreateObject(name string, buf io.Reader) error {
	return nil
}

func (b *Bucket) RenameObject(name, newname string) error {
	return nil
}

func (b *Bucket) RemoveObject(name string) {}

// *(s3.GetObjectOutput).Body => io.ReadCloser
func (b *Bucket) ReadObject(ctx context.Context, name string) (*s3.GetObjectOutput, error) {
	rO, err := b.client.GetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(b.name),
		},
	)
	if err != nil {
		return nil, err
	}

	return rO, nil
}

func (b *Bucket) WriteObject(name string, buf io.Reader) error {
	return nil
}

/*
// Read the object1 from bucket.
rc, err := client.Bucket("bucket").Object("object1").NewReader(ctx)
if err != nil {
	log.Fatal(err)
}
defer rc.Close()
body, err := io.ReadAll(rc)
if err != nil {
	log.Fatal(err)
}
*/
