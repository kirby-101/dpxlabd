package store

/*
 * Cloudflare R2: S3 Bucket
 => sched storage
 => "pkg.dpxlab.de/${ABI}-LATEST/" -> symlink
 ==> "pkg.dpxlab.de/${osreldate}_${portscommit}/" -> REPO
*/

import (
	_ "github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/credentials"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
)

/*

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
*/

//func (s *Store) Start() {}

//func (s *Store) Stop() {}

/*
func main() {
	var bucketName = "sdk-example"
	// Provide your Cloudflare account ID
	var accountId = "<ACCOUNT_ID>"
	// Retrieve your S3 API credentials for your R2 bucket via API tokens
	// (see: https://developers.cloudflare.com/r2/api/tokens)
	var accessKeyId = "<ACCESS_KEY_ID>"
	var accessKeySecret = "<SECRET_ACCESS_KEY>"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"), // Required by SDK but not used by R2
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})

	listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}

	//  {
	//  	"ChecksumAlgorithm": null,
	//  	"ETag": "\"eb2b891dc67b81755d2b726d9110af16\"",
	//  	"Key": "ferriswasm.png",
	//  	"LastModified": "2022-05-18T17:20:21.67Z",
	//  	"Owner": null,
	//  	"Size": 87671,
	//  	"StorageClass": "STANDARD"
	//  }

	listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range listBucketsOutput.Buckets {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}

	// {
	// 		"CreationDate": "2022-05-18T17:19:59.645Z",
	// 		"Name": "sdk-example"
	// }
}
*/
