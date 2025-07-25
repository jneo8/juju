// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package commands

import (
	"fmt"

	"github.com/juju/juju/internal/cmd"
)

// NewCommandByName creates a new command instance by name.
// This function exposes private command constructors for external use.
func NewCommandByName(name string) (cmd.Command, error) {
	switch name {
	case "migrate":
		return newMigrateCommand(), nil
	case "bootstrap":
		return newBootstrapCommand(), nil
	case "switch":
		return newSwitchCommand(), nil
	case "version":
		return newVersionCommand(), nil
	case "sync-agent-binary":
		return newSyncAgentBinaryCommand(), nil
	case "upgrade-model":
		return newUpgradeModelCommand(), nil
	case "upgrade-controller":
		return newUpgradeControllerCommand(), nil
	case "help-hooks":
		return newhelpHookCmdsCommand(), nil
	case "help-actions":
		return newHelpActionCmdsCommand(), nil
	case "debug-log":
		return newDebugLogCommand(nil), nil
	default:
		return nil, fmt.Errorf("unknown command: %s", name)
	}
}