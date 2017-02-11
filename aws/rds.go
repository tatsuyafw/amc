package aws

type rds struct {
	query string
}

func (rds) URL() string {
	b := "REGION.console.aws.amazon.com/rds/home?region=REGION"
	return url(b)
}

func (a rds) Validate() bool {
	return true
}
