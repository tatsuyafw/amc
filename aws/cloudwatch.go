package aws

type cloudwatch struct {
	query string
}

func (cloudwatch) URL() string {
	b := "REGION.console.aws.amazon.com/cloudwatch/home?region=REGION"
	return url(b)
}

func (a cloudwatch) Validate() bool {
	return true
}
