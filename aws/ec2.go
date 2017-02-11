package aws

type ec2 struct {
}

func (ec2) URL() string {
	b := "REGION.console.aws.amazon.com/ec2/v2/home?REGION&region=REGION"
	return url(b)
}
