package aws

type elasticache struct {
	query string
}

func (elasticache) URL() string {
	b := "REGION.console.aws.amazon.com/elasticache/home?region=REGION"
	return url(b)
}

func (a elasticache) Validate() bool {
	return true
}
