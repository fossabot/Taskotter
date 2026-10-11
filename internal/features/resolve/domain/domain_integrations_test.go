// Taskotter 2026.
// SPDX-License-Identifier: Apache-2.0.

package domain_test

import (
	"strings"
	"testing"

	"github.com/task-otter/Taskotter/internal/features/resolve/domain"
	"github.com/task-otter/Taskotter/internal/shared/consts"
)

const (
	notFoundMsg     = "not found in store"
	closeMatches    = "close matches"
	errorFmt        = "Error() = %q"
	attemptedModule = "go/node/npm"
)

// TestResolveErrorMessageIncludesAllDetails verifies attempted module and close matches appear.
func TestResolveErrorMessageIncludesAllDetails(t *testing.T) {
	t.Parallel()

	resolveErr := &domain.ResolveError{
		LogicalTask:  consts.Go,
		Attempted:    attemptedModule,
		Message:      notFoundMsg,
		CloseMatches: []string{consts.Go, "golang"},
	}

	msg := resolveErr.Error()
	assertContains(t, msg, "attempted source module")
	assertContains(t, msg, closeMatches)

	if resolveErr.Detail() != msg {
		t.Fatalf("Detail() = %q, want %q", resolveErr.Detail(), msg)
	}
}

// TestResolveErrorMessageOmitsOptionalDetails verifies the minimal message form.
func TestResolveErrorMessageOmitsOptionalDetails(t *testing.T) {
	t.Parallel()

	resolveErr := &domain.ResolveError{
		Attempted:    "",
		CloseMatches: nil,

		LogicalTask: consts.Go,
		Message:     notFoundMsg,
	}

	msg := resolveErr.Error()
	assertOmits(t, msg, "attempted")
	assertOmits(t, msg, closeMatches)
}

// TestResolveErrorEmptyMessageBranches covers Error paths after Message is empty.
func TestResolveErrorEmptyMessageBranches(t *testing.T) {
	t.Parallel()

	attempted := resolveErrorForTask(consts.Go)

	attempted.Attempted = attemptedModule
	assertContains(t, attempted.Error(), attemptedModule)

	matches := resolveErrorForTask(consts.Go)

	matches.CloseMatches = []string{consts.Go}
	assertContains(t, matches.Error(), closeMatches)
	assertContains(t, resolveErrorForTask(consts.Go).Error(), `task "go"`)
	assertContains(t, resolveErrorForTask(consts.Empty).Error(), `task ""`)
}

// TestResolveErrorTask covers Task().
func TestResolveErrorTask(t *testing.T) {
	t.Parallel()

	resolveErr := &domain.ResolveError{
		Attempted:    "",
		Message:      "",
		CloseMatches: nil,
		LogicalTask:  consts.Go,
	}

	if resolveErr.Task() != consts.Go {
		t.Fatalf("Task() = %q", resolveErr.Task())
	}
}

func assertContains(t *testing.T, msg, needle string) {
	t.Helper()

	if !strings.Contains(msg, needle) {
		t.Fatalf(errorFmt, msg)
	}
}

func assertOmits(t *testing.T, msg, needle string) {
	t.Helper()

	if strings.Contains(msg, needle) {
		t.Fatalf(errorFmt, msg)
	}
}

func resolveErrorForTask(task string) *domain.ResolveError {
	resolveErr := new(domain.ResolveError)

	resolveErr.LogicalTask = task

	return resolveErr
}
