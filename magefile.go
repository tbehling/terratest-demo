//+build mage

package main

import (
    "github.com/magefile/mage/sh"
)

func Build() error {
	if err := sh.Run("go", "version"); err != nil {
		return err
	}
	return nil
}