package main

import (
	"os"
	"testing"
)

func TestURL(t *testing.T) {
	reset := setTestEnv("AWS_REGION", "ap-northeast-1")
	defer reset()

	a := AWS{}
	s := "ec2"

	actual := a.URL(s)
	expected := "https://ap-northeast-1.console.aws.amazon.com/ec2/v2/home?ap-northeast-1&region=ap-northeast-1"

	if actual != expected {
		t.Errorf("expected: %v, but got %v", expected, actual)
	}

}

func setTestEnv(key, val string) func() {
	preVal := os.Getenv(key)
	os.Setenv(key, val)
	return func() {
		os.Setenv(key, preVal)
	}
}
