package aws

type s3 struct {
}

func (s3) URL() string {
	b := "console.aws.amazon.com/s3/home?region=REGION"
	return url(b)
}
