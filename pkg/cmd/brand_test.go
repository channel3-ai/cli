// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/public-sdk-cli/internal/mocktest"
)

func TestBrandsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brands", "retrieve",
			"--brand-id", "brand_id",
		)
	})
}

func TestBrandsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brands", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestBrandsFind(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brands", "find",
			"--query", "query",
		)
	})
}

func TestBrandsSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"brands", "search",
			"--query", "x",
			"--limit", "1",
		)
	})
}
