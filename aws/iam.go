package aws

type iam struct {
	query string
}

func (iam) URL() string {
	b := "console.aws.amazon.com/iam/home?region=REGION"
	return url(b)
}

func (a iam) Validate() bool {
	return true
}
