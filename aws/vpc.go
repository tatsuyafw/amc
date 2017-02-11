package aws

type vpc struct {
}

func (vpc) URL() string {
	b := "REGION.console.aws.amazon.com/vpc/home?region=REGION"
	return url(b)
}
