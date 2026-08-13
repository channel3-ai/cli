// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/channel3-ai/cli/internal/apiquery"
	"github.com/channel3-ai/cli/internal/requestflag"
	"github.com/channel3-ai/sdk-go"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var conversationsClientTokensCreate = cli.Command{
	Name:    "create",
	Usage:   "Mint a short-lived, browser-safe token. With `conversation_id` the token\ncontinues and reads that thread; without it, the token's first turn creates the\nthread and binds the token to it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "conversation-id",
			BodyPath: "conversation_id",
		},
		&requestflag.Flag[int64]{
			Name:     "ttl-seconds",
			Default:  1800,
			BodyPath: "ttl_seconds",
		},
	},
	Action:          handleConversationsClientTokensCreate,
	HideHelpCommand: true,
}

var conversationsClientTokensRevoke = cli.Command{
	Name:    "revoke",
	Usage:   "Revoke a client token immediately. The token travels in the request body, not\nthe URL, so it stays out of access logs; only the minting vendor can revoke it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "token",
			Required: true,
			BodyPath: "token",
		},
	},
	Action:          handleConversationsClientTokensRevoke,
	HideHelpCommand: true,
}

func handleConversationsClientTokensCreate(ctx context.Context, cmd *cli.Command) error {
	client := channel3go.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := channel3go.ConversationClientTokenNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Conversations.ClientTokens.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "conversations:client-tokens create",
		Transform:      transform,
	})
}

func handleConversationsClientTokensRevoke(ctx context.Context, cmd *cli.Command) error {
	client := channel3go.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := channel3go.ConversationClientTokenRevokeParams{}

	return client.Conversations.ClientTokens.Revoke(ctx, params, options...)
}
