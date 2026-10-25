package store

/*
o.BaseEndpoint = aws.String(

	fmt.Sprintf(creds.endPoint, creds.accountId), //!

)
*/
type StoreCredentials struct {
	endPoint        string // https://%s.r2.cloudflarestorage.com
	accountId       string // Provide your Cloudflare account ID#
	accessKeyId     string // Retrieve your S3 API credentials for your R2 bucket via API tokens
	accessKeySecret string // (see: https://developers.cloudflare.com/r2/api/tokens)
}
