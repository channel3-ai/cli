// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/channel3-ai/cli/internal/mocktest"
	"github.com/channel3-ai/cli/internal/requestflag"
)

func TestConversationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "create",
			"--message", "{parts: [{text: text, type: text}], role: user}",
			"--context", "{application_context: application_context, user_context: user_context}",
			"--conversation-id", "conversation_id",
			"--filters", "{age: [newborn], attributes: {foo: [string]}, availability: [InStock], brand_ids: [string], category_ids: [string], colors: {palette: [{hex: hex, percentage: 0}], match: strict}, conditions: [new], dimensions: {height: {unit: mm, max: 0, min: 0}, length: {unit: mm, max: 0, min: 0}, weight: {unit: mg, max: 0, min: 0}, width: {unit: mm, max: 0, min: 0}}, exclude_brand_ids: [string], exclude_category_ids: [string], exclude_website_ids: [string], gender: male, price: {max_price: 0, min_price: 0}, sale: on_sale, website_ids: [string]}",
			"--stream=true",
			"--x-user-id", "x-user-id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(conversationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "create",
			"--message.parts", "[{text: text, type: text}]",
			"--message.role", "user",
			"--context.application-context", "application_context",
			"--context.user-context", "user_context",
			"--conversation-id", "conversation_id",
			"--filters.age", "[newborn]",
			"--filters.attributes", "{foo: [string]}",
			"--filters.availability", "[InStock]",
			"--filters.brand-ids", "[string]",
			"--filters.category-ids", "[string]",
			"--filters.colors", "{palette: [{hex: hex, percentage: 0}], match: strict}",
			"--filters.conditions", "[new]",
			"--filters.dimensions", "{height: {unit: mm, max: 0, min: 0}, length: {unit: mm, max: 0, min: 0}, weight: {unit: mg, max: 0, min: 0}, width: {unit: mm, max: 0, min: 0}}",
			"--filters.exclude-brand-ids", "[string]",
			"--filters.exclude-category-ids", "[string]",
			"--filters.exclude-website-ids", "[string]",
			"--filters.gender", "male",
			"--filters.price", "{max_price: 0, min_price: 0}",
			"--filters.sale", "on_sale",
			"--filters.website-ids", "[string]",
			"--stream=true",
			"--x-user-id", "x-user-id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message:\n" +
			"  parts:\n" +
			"    - text: text\n" +
			"      type: text\n" +
			"  role: user\n" +
			"context:\n" +
			"  application_context: application_context\n" +
			"  user_context: user_context\n" +
			"conversation_id: conversation_id\n" +
			"filters:\n" +
			"  age:\n" +
			"    - newborn\n" +
			"  attributes:\n" +
			"    foo:\n" +
			"      - string\n" +
			"  availability:\n" +
			"    - InStock\n" +
			"  brand_ids:\n" +
			"    - string\n" +
			"  category_ids:\n" +
			"    - string\n" +
			"  colors:\n" +
			"    palette:\n" +
			"      - hex: hex\n" +
			"        percentage: 0\n" +
			"    match: strict\n" +
			"  conditions:\n" +
			"    - new\n" +
			"  dimensions:\n" +
			"    height:\n" +
			"      unit: mm\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"    length:\n" +
			"      unit: mm\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"    weight:\n" +
			"      unit: mg\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"    width:\n" +
			"      unit: mm\n" +
			"      max: 0\n" +
			"      min: 0\n" +
			"  exclude_brand_ids:\n" +
			"    - string\n" +
			"  exclude_category_ids:\n" +
			"    - string\n" +
			"  exclude_website_ids:\n" +
			"    - string\n" +
			"  gender: male\n" +
			"  price:\n" +
			"    max_price: 0\n" +
			"    min_price: 0\n" +
			"  sale: on_sale\n" +
			"  website_ids:\n" +
			"    - string\n" +
			"stream: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conversations", "create",
			"--x-user-id", "x-user-id",
		)
	})
}

func TestConversationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations", "retrieve",
			"--conversation-id", "conversation_id",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
