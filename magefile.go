//+build mage

package main

import (
	"testing"

    "github.com/magefile/mage/sh"
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
	t.Error("placeholder failure")
}

