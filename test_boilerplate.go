//+build mage

package main

import (
	"errors"
	"io"
	"regexp"
	"testing"
)

func RunGoTests(tests []testing.InternalTest) error {
	m := testing.MainStart(&nullTestDeps{}, tests, nil, nil)
	if x := m.Run(); x > 0 {
		return errors.New("testing.MainStart returned non-zero")
	}
	return nil
}

// nullTestDeps implements testing.testDeps to stub dependencies for testing.MainStart()
//
// See https://stackoverflow.com/questions/59064256/is-it-possible-to-call-a-test-func-from-another-file-to-start-the-testing
//     https://golang.org/src/testing/testing.go?s=38015:38129#L1146
//     https://golang.org/src/testing/internal/testdeps/deps.go
type nullTestDeps struct{}

func (d *nullTestDeps) MatchString(pat, str string) (bool, error)         { return regexp.MatchString(pat, str) }
func (d *nullTestDeps) StartCPUProfile(_ io.Writer) error                 { return nil }
func (d *nullTestDeps) StopCPUProfile()                                   {}
func (d *nullTestDeps) WriteHeapProfile(_ io.Writer) error                { return nil }
func (d *nullTestDeps) WriteProfileTo(_ string, _ io.Writer, _ int) error { return nil }
func (d *nullTestDeps) ImportPath() string                                { return "" }
func (d *nullTestDeps) StartTestLog(io.Writer)                            { return }
func (d *nullTestDeps) StopTestLog() error                                { return nil }
