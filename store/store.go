package store

/*
 * Cloudflare R2: S3 Bucket
 => sched storage
 => "pkg.dpxlab.de/${ABI}-LATEST/" -> symlink
 ==> "pkg.dpxlab.de/${osreldate}_${portscommit}/" -> REPO
*/

import (
	"context"

	"cloud.google.com/go/storage"
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
