package aws

type cloudwatch struct {
}

func (cloudwatch) URL() string {
	b := "REGION.console.aws.amazon.com/cloudwatch/home?region=REGION"
	return url(b)
}
