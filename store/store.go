package store

/*
 * Cloudflare R2: S3 Bucket
 => pkgsched storage
 => "https://pkg.dpxlab.de/${ABI}/latest" -> MANIFEST
 => html/css schmücken?

https://docs.cloud.google.com/storage/docs/samples/storage-transfer-manager-download-chunks-concurrently
*/

import (
	"context"

	"cloud.google.com/go/storage"
	_ "cloud.google.com/go/storage/transfermanager"
	"google.golang.org/api/option"
)

// ! client.Close()
type Store struct {
	client *storage.Client
	ctx    context.Context
}

func t() {
	option.WithAuthCredentials()
}

func New(options ...option.ClientOption) (*Store, error) {
	var ctx = context.Background()

	client, err := storage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &Store{
		client: client,
		ctx:    ctx,
	}, nil
}

//func (s *Store) Start() {}

//func (s *Store) Stop() {}
