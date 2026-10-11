// Taskotter 2026.
// SPDX-License-Identifier: Apache-2.0.

package taskfile

import (
	"errors"

	"github.com/task-otter/Taskotter/internal/features/root/ports"
)

var (
	_ ports.RootTemplateProvider = (*Ops)(nil)
	_ ports.IncludeRewriter      = (*Ops)(nil)
	_ ports.RootTaskfileUpdater  = (*Ops)(nil)

	errNoModuleVars = errors.New("module Taskfile has no vars")
)
