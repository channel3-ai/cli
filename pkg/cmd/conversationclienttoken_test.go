// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/channel3-ai/cli/internal/mocktest"
)

func TestConversationsClientTokensCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations:client-tokens", "create",
			"--conversation-id", "conversation_id",
			"--session-id", "session_id",
			"--ttl-seconds", "60",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"conversation_id: conversation_id\n" +
			"session_id: session_id\n" +
			"ttl_seconds: 60\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conversations:client-tokens", "create",
		)
	})
}

func TestConversationsClientTokensRevoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"conversations:client-tokens", "revoke",
			"--token", "token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("token: token")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"conversations:client-tokens", "revoke",
		)
	})
}
