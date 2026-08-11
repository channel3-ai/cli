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
