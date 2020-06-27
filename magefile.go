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
		{
			Name: "TestRunTerratest",
			F:     TestRunTerratest,
		},
	})
}

func TestRunTerratest(t *testing.T) {
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

