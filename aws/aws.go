package aws

import (
	"fmt"
	"os"
	"strings"
)

// DefaultRegion is used when AWS_REGION env variable is not set.
const DefaultRegion = "us-east-1"

// AWS represents amazon web services.
type AWS interface {
	URL() string
}

// New creates an AWS based on the given service.
func New(service string) (AWS, error) {
	if !validate(service) {
		return nil, fmt.Errorf("error: %v is invalid service", service)
	}
	var a AWS
	// TODO: use reflect package for clean codes.
	switch service {
	case "cloudwatch":
		a = &cloudwatch{}
	case "ec2":
		a = &ec2{}
	case "ecs":
		a = &ecs{}
	case "elasticache":
		a = &elasticache{}
	case "iam":
		a = &iam{}
	case "rds":
		a = &rds{}
	case "route53":
		a = &route53{}
	case "s3":
		a = &s3{}
	case "vpc":
		a = &vpc{}
	}
	return a, nil
}

// Supported returns supported service
func Supported() []string {
	return []string{
		"cloudwatch",
		"ec2",
		"ecs",
		"elasticache",
		"iam",
		"rds",
		"route53",
		"s3",
		"vpc",
	}
}

func url(base string) string {
	r := region()
	return "https://" + strings.Replace(base, "REGION", r, -1)
}

func validate(service string) bool {
	for _, s := range Supported() {
		if s == service {
			return true
		}
	}
	return false
}

func region() string {
	r := os.Getenv("AWS_REGION")
	if r == "" {
		r = DefaultRegion
	}
	return r
}
