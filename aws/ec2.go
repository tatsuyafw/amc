package aws

import "github.com/tatsuyafw/amc/util"

type ec2 struct {
	query string
}

var queries = map[string]string{
	"addresses":       "Addresses",
	"events":          "Events",
	"images":          "Images",
	"instances":       "Instances",
	"limits":          "Limits",
	"reports":         "Reports",
	"security-groups": "SecurityGroups",
	"snapshots":       "Snapshots",
	"tags":            "Tags",
	"volumes":         "Volumes",
}

func (a ec2) URL() string {
	b := "REGION.console.aws.amazon.com/ec2/v2/home?REGION&region=REGION"
	if a.query != "" {
		b += "#" + queries[a.query]
	}
	return url(b)
}

func (a ec2) Validate() bool {
	if a.query == "" {
		return true
	}
	return util.IncludeStr(util.KeysStr(queries), a.query)
}
