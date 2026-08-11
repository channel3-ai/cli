// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/channel3-ai/cli/internal/apiquery"
	"github.com/channel3-ai/cli/internal/requestflag"
	"github.com/channel3-ai/sdk-go"
	"github.com/urfave/cli/v3"
)

var responsesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Run a shopping conversation turn.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "attachment",
			BodyPath: "attachments",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "context",
			Usage:    "Partner-supplied context pinned to the top of a conversation thread.",
			BodyPath: "context",
		},
		&requestflag.Flag[*string]{
			Name:     "conversation-id",
			BodyPath: "conversation_id",
		},
		&requestflag.Flag[bool]{
			Name:     "debug",
			Default:  false,
			BodyPath: "debug",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "image",
			BodyPath: "image",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "message",
			BodyPath: "message",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "message",
			BodyPath: "messages",
		},
		&requestflag.Flag[string]{
			Name:       "x-user-id",
			Usage:      "Optional user identifier to attribute clicks and sales to a user in your system. Channel3 appends it to buy URLs in the response.",
			HeaderPath: "x-user-id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleResponsesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"attachment": {
		&requestflag.InnerFlag[string]{
			Name:                  "attachment.url",
			InnerField:            "url",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "attachment.key",
			InnerField:            "key",
			OuterIsArrayOfObjects: true,
		},
	},
	"context": {
		&requestflag.InnerFlag[*string]{
			Name:       "context.application-context",
			Usage:      "What platform or surface is hosting this conversation.",
			InnerField: "application_context",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "context.user-context",
			Usage:      "Who the conversation is with (profile, preferences, session facts).",
			InnerField: "user_context",
		},
	},
	"image": {
		&requestflag.InnerFlag[*string]{
			Name:       "image.base64",
			InnerField: "base64",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "image.url",
			InnerField: "url",
		},
	},
	"message": {
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			InnerField: "role",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "message.parts",
			InnerField: "parts",
		},
	},
})

func handleResponsesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ResponseNewParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	stream := client.Responses.NewStreaming(ctx, params, options...)
	maxItems := int64(-1)
	if cmd.IsSet("max-items") {
		maxItems = cmd.Value("max-items").(int64)
	}
	return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "responses create",
		Transform:      transform,
	})
}
