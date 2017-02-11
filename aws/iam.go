package aws

type iam struct {
}

func (iam) URL() string {
	b := "console.aws.amazon.com/iam/home?region=REGION"
	return url(b)
}
