package aws

type vpc struct {
	query string
}

func (vpc) URL() string {
	b := "REGION.console.aws.amazon.com/vpc/home?region=REGION"
	return url(b)
}

func (a vpc) Validate() bool {
	return true
}
