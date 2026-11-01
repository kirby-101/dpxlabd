package storage

/*
 * Cloudflare R2: S3 Bucket
 => sched storage
 => "pkg.dpxlab.de/${ABI}-LATEST/" -> symlink
 ==> "pkg.dpxlab.de/${osreldate}_${portscommit}/" -> REPO
*/

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/credentials"
	_ "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

/*
o.BaseEndpoint = aws.String(

	fmt.Sprintf(creds.endPoint, creds.accountId), //!

)
*/

// ! client.Close() ?
type StorageClient struct {
	client *s3.Client
	ctx    context.Context
}

func New() {}

/*
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
*/

func (s *StorageClient) OpenBucket(name string) (*Bucket, error) {
	//var b = &Bucket{}

	//s.client.GetBucket

	return nil, nil
}
