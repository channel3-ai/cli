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

var websitesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Resolve a website URL to its ID and best_commission_rate. Tip: website_ids\nfilters accept domains directly, so this lookup is most useful for retrieving\ncommission rates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "query",
			Required:  true,
			QueryPath: "query",
		},
	},
	Action:          handleWebsitesRetrieve,
	HideHelpCommand: true,
}

func handleWebsitesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := channel3go.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := channel3go.WebsiteGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Websites.Get(ctx, params, options...)
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
		Title:          "websites retrieve",
		Transform:      transform,
	})
}
