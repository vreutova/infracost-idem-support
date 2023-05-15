package aws_test

import (
	"testing"

	"github.com/infracost/infracost/internal/providers/idem/idemtest"
)

func TestConfigConfigRule(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	opts := idemtest.DefaultGoldenFileOptions()
	opts.CaptureLogs = true
	idemtest.GoldenFileResourceTestsWithOpts(t, "config_config_rule_test", opts)
}