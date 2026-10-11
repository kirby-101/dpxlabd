package store

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

type Bucket struct{}
