package main

import (
	"fmt"
	"os"
	"strings"
)

// DefaultRegion is used when AWS_REGION env variable is not set.
const DefaultRegion = "us-east-1"

// An AWS represents amazon web services.
type AWS struct {
	service string
}

func newAWS(s string) (*AWS, error) {
	a := &AWS{service: s}
	if !a.validate() {
		return nil, fmt.Errorf("error: %v is invalid service", s)
	}
	return a, nil
}

// URL returns the AWS management console service URL.
func (a *AWS) URL() string {
	m := urlmap()
	r := region()
	s := a.service
	return "https://" + strings.Replace(m[s], "REGION", r, -1)
}

// Validate checks whether the given service is valid.
func (a *AWS) validate() bool {
	m := urlmap()
	_, ok := m[a.service]
	return ok
}

func urlmap() map[string]string {
	return map[string]string{
		"ec2":         "REGION.console.aws.amazon.com/ec2/v2/home?REGION&region=REGION",
		"ecs":         "REGION.console.aws.amazon.com/ecs/home?region=REGION",
		"elasticache": "REGION.console.aws.amazon.com/elasticache/home?region=REGION",
		"iam":         "console.aws.amazon.com/iam/home?region=REGION",
		"rds":         "REGION.console.aws.amazon.com/rds/home?region=REGION",
		"route53":     "console.aws.amazon.com/route53/home?region=REGION",
		"s3":          "console.aws.amazon.com/s3/home?region=REGION",
		"vpc":         "REGION.console.aws.amazon.com/vpc/home?region=REGION",
	}
}

func (*AWS) supported() []string {
	m := urlmap()
	s := []string{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

func region() string {
	r := os.Getenv("AWS_REGION")
	if r == "" {
		r = DefaultRegion
	}
	return r
}
