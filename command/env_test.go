package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestEnvCommand_implement(t *testing.T) {
	var _ cli.Command = &EnvCommand{}
}
