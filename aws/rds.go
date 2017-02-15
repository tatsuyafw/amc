package aws

import "github.com/tatsuyafw/amc/util"

type rds struct {
	query string
}

var rdsQueries = map[string]string{
	"parameter-groups": "parameter-groups:",
}

func (a rds) URL() string {
	b := "REGION.console.aws.amazon.com/rds/home?region=REGION"
	if a.query != "" {
		b += "#" + rdsQueries[a.query]
	}
	return url(b)
}

func (a rds) Validate() bool {
	if a.query == "" {
		return true
	}
	return util.IncludeStr(util.KeysStr(rdsQueries), a.query)
}
