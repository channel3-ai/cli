// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/public-sdk-cli/internal/mocktest"
)

func TestCategoriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"categories", "retrieve",
			"--slug", "slug",
		)
	})
}

func TestCategoriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"categories", "list",
			"--max-items", "10",
			"--page", "1",
			"--page-size", "1",
			"--roots-only=true",
		)
	})
}

func TestCategoriesSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"categories", "search",
			"--query", "x",
			"--limit", "1",
		)
	})
}
