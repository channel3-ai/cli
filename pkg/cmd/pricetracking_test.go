// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/channel3-ai/cli/internal/mocktest"
)

func TestPriceTrackingListSubscriptions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"price-tracking", "list-subscriptions",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}

func TestPriceTrackingRetrieveHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"price-tracking", "retrieve-history",
			"--canonical-product-id", "canonical_product_id",
			"--days", "1",
		)
	})
}

func TestPriceTrackingStart(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"price-tracking", "start",
			"--canonical-product-id", "canonical_product_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("canonical_product_id: canonical_product_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"price-tracking", "start",
		)
	})
}

func TestPriceTrackingStop(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"price-tracking", "stop",
			"--canonical-product-id", "canonical_product_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("canonical_product_id: canonical_product_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"price-tracking", "stop",
		)
	})
}
