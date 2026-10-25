package store

/*
 * Cloudflare R2: S3 Bucket
 => sched storage
 => "pkg.dpxlab.de/${ABI}-LATEST/" -> symlink
 ==> "pkg.dpxlab.de/${osreldate}_${portscommit}/" -> REPO
*/

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ! client.Close() ?
type Store struct {
	creds  StoreCredentials
	client *s3.Client
	ctx    context.Context
}

func New(ctx context.Context, creds StoreCredentials) (*Store, error) {
	var s = &Store{
		creds: creds,
		ctx:   ctx,
	}

	cfg, err := config.LoadDefaultConfig(
		s.ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				creds.accesKeyId,
				creds.accesKeySecret,
				"",
			),
		),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	s.client = s3.NewFromConfig(
		cfg,
		func(o *s3.Options) {
			o.BaseEndpoint = aws.String(creds.endPoint)
		},
	)

	return s, nil
}

func (s *Store) OpenBucket(name string) (any, error) {
	var bucket = &Bucket{}

	//s.client.GetBucket
}
