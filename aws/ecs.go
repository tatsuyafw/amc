package aws

type ecs struct {
}

func (ecs) URL() string {
	b := "REGION.console.aws.amazon.com/ecs/home?region=REGION"
	return url(b)
}
