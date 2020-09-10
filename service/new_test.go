package service

import (
	"testing"

	"github.com/octoproject/octo-cli/config"
)

func TestNewFunction(t *testing.T) {

	s := &config.Service{
		Query:       "select * from users where id = $1",
		ServiceName: "get-users",
		ServiceType: "read",
		Params: config.Params{
			Type: "path",
			Schema: []map[string]config.Validator{
				{
					"id": config.Validator{
						Required: true,
					},
				},
			},
		},
	}

	err := NewFunction(s)
	if err != nil {
		t.Fatal(err)
	}

}
