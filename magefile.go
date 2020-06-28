//+build mage

package main

import (
	"testing"

	"github.com/magefile/mage/sh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"

)

func Build() error {
	if err := sh.Run("go", "version"); err != nil {
		return err
	}
	return nil
}

func Terratest() error {
	return RunGoTests([]testing.InternalTest{
		{Name: "TestTerraformHello", F: TestTerraformHello},
	})
}

func TerratestHttp() error {
	return RunGoTests([]testing.InternalTest{
		{Name: "TestHttp", F: TestHttp},
	})
}

func TestTerraformHello(t *testing.T) {
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: ".",
	}

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "Hello, World!", output)
}

func TestHttp(t *testing.T) {
	withDockerContainer(t, "nginx", func(ci *docker.ContainerInspect) {
		url := fmt.Sprintf("http://%s:%d", "localhost", ci.Ports[0].HostPort)

		http_helper.HttpGetWithRetryWithCustomValidation(t, url, nil, 30, time.Second, func(code int, body string) bool {
			return assert.Equal(t, code, 200) && assert.Contains(t, body, "Welcome to nginx!")
		})
	})
}

func withDockerContainer(t *testing.T, tag string, f func(*docker.ContainerInspect)) {
	opts := &docker.RunOptions{Detach: true, Remove: true, OtherOptions: []string{"-P"}}
	docker_id := docker.RunAndGetID(t, tag, opts)
	defer docker.Stop(t, []string{docker_id}, &docker.StopOptions{})
	container_info := docker.Inspect(t, docker_id)

	f(container_info)
}
