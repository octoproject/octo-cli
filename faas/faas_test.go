package faas

import "testing"

func TestDeployFunction(t *testing.T) {
	c := New("", "", "")

	err := c.DeployFunction(&Function{
		ServiceName: "functions/nodeinfo-http:latest",
		Image:       "nodeinfo",
		Namespace:   "test",
	})

	if err != nil {
		t.Fatal(err)
	}
}
