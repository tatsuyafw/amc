package aws

type dynamodb struct {
	query string
}

var dynamodbQueries = map[string]string{
	"tables": "tables:",
}

func (a dynamodb) URL() string {
	b := "REGION.console.aws.amazon.com/dynamodb/home?region=REGION"
	if a.query != "" {
		b += "#" + dynamodbQueries[a.query]
	}
	return url(b)
}

func (a dynamodb) Validate() bool {
	return true
}
