package aws

type elb struct {
	query string
}

func (elb) URL() string {
	b := "REGION.console.aws.amazon.com/ec2/v2/home?region=REGION#LoadBalancers:"
	return url(b)
}

func (a elb) Validate() bool {
	return true
}
