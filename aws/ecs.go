package aws

type ecs struct {
	query string
}

func (ecs) URL() string {
	b := "REGION.console.aws.amazon.com/ecs/home?region=REGION"
	return url(b)
}

func (a ecs) Validate() bool {
	return true
}
