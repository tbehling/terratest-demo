//+build mage

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

	"github.com/stretchr/testify/assert"
)

// TerratestHttp demonstrates testing an Nginx server, as an ad-hoc Docker container.
func TerratestHttp() error {
	return RunGoTests([]testing.InternalTest{
		{Name: "TestHttp", F: TestHttp},
	})
}

func TestHttp(t *testing.T) {
	withDockerContainer(t, "nginx", func(ci *docker.ContainerInspect) {
		url := fmt.Sprintf("http://%s:%d", "localhost", ci.Ports[0].HostPort)

		http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, 30, time.Second, func(code int, body string) bool {
			return assert.Equal(t, code, 200) && assert.Contains(t, body, "Welcome to nginx!")
		})
	})
}
