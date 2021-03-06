package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/keptn/keptn/cli/pkg/credentialmanager"
	"github.com/keptn/keptn/cli/pkg/logging"
)

func init() {
	logging.InitLoggers(os.Stdout, os.Stdout, os.Stderr)
}

// TestEvaluationDoneGetEvent tests the evaluation-done command
func TestGetStage(t *testing.T) {

	credentialmanager.MockAuthCreds = true

	cmd := fmt.Sprintf("get stage hardening --project=sockshop --mock")
	_, err := executeActionCommandC(cmd)
	if err != nil {
		t.Errorf(unexpectedErrMsg, err)
	}
}
