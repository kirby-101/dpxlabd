package store

import (
	"bytes"
	"io"
)

type Bucket struct{}

func (b *Bucket) CreateObject(name string, buf io.Reader) error {}
func (b *Bucket) RenameObject(name, newname string) error       {}
func (b *Bucket) RemoveObject(name string)                      {}
func (b *Bucket) ReadObject(name string) (*bytes.Buffer, error) {}
func (b *Bucket) WriteObject(name string, buf io.Reader) error  {}

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
