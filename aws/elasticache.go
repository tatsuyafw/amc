package aws

type elasticache struct {
}

func (elasticache) URL() string {
	b := "REGION.console.aws.amazon.com/elasticache/home?region=REGION"
	return url(b)
}
