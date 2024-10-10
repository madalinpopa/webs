package examples

import (
	"github.com/madalinpopa/webs"
	"time"
)

var client = getClient()

func getClient() *webs.Client {
	client := webs.NewClientBuilder().
		SetConnectTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()
	return client
}
