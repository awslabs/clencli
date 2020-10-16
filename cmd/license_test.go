package cmd

import "testing"

func TestLicenseCmdWithNoArgsAndNoFlags(t *testing.T) {
	cmd := LicenseCmd()
	cmd.Execute()
}
