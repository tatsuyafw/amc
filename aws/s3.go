package aws

type s3 struct {
	query string
}

func (s3) URL() string {
	b := "console.aws.amazon.com/s3/home?region=REGION"
	return url(b)
}

func (a s3) Validate() bool {
	return true
}
