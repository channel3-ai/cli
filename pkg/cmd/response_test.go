// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/channel3-ai/cli/internal/mocktest"
	"github.com/channel3-ai/cli/internal/requestflag"
)

func TestResponsesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"responses", "create",
			"--max-items", "10",
			"--attachment", "[{url: url, key: key}]",
			"--context", "{application_context: application_context, user_context: user_context}",
			"--conversation-id", "conversation_id",
			"--debug=true",
			"--filters", "{age: [newborn], attributes: {foo: [string]}, availability: [InStock], brand_ids: [string], category_ids: [string], colors: {palette: [{hex: hex, percentage: 0}], match: strict}, conditions: [new], dimensions: {height: {unit: mm, max: 0, min: 0}, length: {unit: mm, max: 0, min: 0}, weight: {unit: mg, max: 0, min: 0}, width: {unit: mm, max: 0, min: 0}}, exclude_brand_ids: [string], exclude_category_ids: [string], exclude_website_ids: [string], gender: male, price: {max_price: 0, min_price: 0}, sale: on_sale, website_ids: [string]}",
			"--image", "{base64: base64, url: url}",
			"--message", "{role: role, parts: [{type: text, input: {foo: bar}, modelOnly: true, output: {foo: bar}, suggestedReplies: [string], text: text, toolCallId: toolCallId, toolName: toolName, url: url}]}",
			"--message", "{role: role, parts: [{type: text, input: {foo: bar}, modelOnly: true, output: {foo: bar}, suggestedReplies: [string], text: text, toolCallId: toolCallId, toolName: toolName, url: url}]}",
			"--x-user-id", "x-user-id",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(responsesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"responses", "create",
			"--max-items", "10",
			"--attachment.url", "url",
			"--attachment.key", "key",
			"--context.application-context", "application_context",
			"--context.user-context", "user_context",
			"--conversation-id", "conversation_id",
			"--debug=true",
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
			"--image.base64", "base64",
			"--image.url", "url",
			"--message.role", "role",
			"--message.parts", "[{type: text, input: {foo: bar}, modelOnly: true, output: {foo: bar}, suggestedReplies: [string], text: text, toolCallId: toolCallId, toolName: toolName, url: url}]",
			"--message.role", "role",
			"--message.parts", "[{type: text, input: {foo: bar}, modelOnly: true, output: {foo: bar}, suggestedReplies: [string], text: text, toolCallId: toolCallId, toolName: toolName, url: url}]",
			"--x-user-id", "x-user-id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attachments:\n" +
			"  - url: url\n" +
			"    key: key\n" +
			"context:\n" +
			"  application_context: application_context\n" +
			"  user_context: user_context\n" +
			"conversation_id: conversation_id\n" +
			"debug: true\n" +
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
			"image:\n" +
			"  base64: base64\n" +
			"  url: url\n" +
			"message:\n" +
			"  role: role\n" +
			"  parts:\n" +
			"    - type: text\n" +
			"      input:\n" +
			"        foo: bar\n" +
			"      modelOnly: true\n" +
			"      output:\n" +
			"        foo: bar\n" +
			"      suggestedReplies:\n" +
			"        - string\n" +
			"      text: text\n" +
			"      toolCallId: toolCallId\n" +
			"      toolName: toolName\n" +
			"      url: url\n" +
			"messages:\n" +
			"  - role: role\n" +
			"    parts:\n" +
			"      - type: text\n" +
			"        input:\n" +
			"          foo: bar\n" +
			"        modelOnly: true\n" +
			"        output:\n" +
			"          foo: bar\n" +
			"        suggestedReplies:\n" +
			"          - string\n" +
			"        text: text\n" +
			"        toolCallId: toolCallId\n" +
			"        toolName: toolName\n" +
			"        url: url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"responses", "create",
			"--max-items", "10",
			"--x-user-id", "x-user-id",
		)
	})
}
