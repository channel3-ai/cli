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

var conversationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Return metadata for a conversation thread.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "conversation-id",
			Required:  true,
			PathParam: "conversation_id",
		},
	},
	Action:          handleConversationsRetrieve,
	HideHelpCommand: true,
}

func handleConversationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := channel3go.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("conversation-id") && len(unusedArgs) > 0 {
		cmd.Set("conversation-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Conversations.Get(ctx, cmd.Value("conversation-id").(string), options...)
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
		Title:          "conversations retrieve",
		Transform:      transform,
	})
}
