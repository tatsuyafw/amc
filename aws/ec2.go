package aws

import "github.com/tatsuyafw/amc/util"

type ec2 struct {
	query string
}

var querys = map[string]string{
	"instances": "Instances",
}

func (a ec2) URL() string {
	b := "REGION.console.aws.amazon.com/ec2/v2/home?REGION&region=REGION"
	if a.query != "" {
		b += "#" + querys[a.query]
	}
	return url(b)
}

func (a ec2) Validate() bool {
	if a.query == "" {
		return true
	}
	return util.IncludeStr(util.KeysStr(querys), a.query)
}
