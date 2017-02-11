package aws

type route53 struct {
	query string
}

func (route53) URL() string {
	b := "REGION.console.aws.amazon.com/cloudwatch/home?region=REGION"
	return url(b)
}

func (a route53) Validate() bool {
	return true
}
