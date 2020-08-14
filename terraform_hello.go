//+build mage

package main

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/stretchr/testify/assert"
)

// TerraformHello demonstrates launching a null Terraform module.
// It is exported by Mage thanks to its function signature.
func TerraformHello() error {
	return RunGoTests([]testing.InternalTest{
		{Name: "TestTerraformHello", F: TestTerraformHello},
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
