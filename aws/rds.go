package aws

type rds struct {
}

func (rds) URL() string {
	b := "REGION.console.aws.amazon.com/rds/home?region=REGION"
	return url(b)
}
