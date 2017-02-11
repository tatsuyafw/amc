package aws

type route53 struct {
}

func (route53) URL() string {
	b := "REGION.console.aws.amazon.com/cloudwatch/home?region=REGION"
	return url(b)
}
