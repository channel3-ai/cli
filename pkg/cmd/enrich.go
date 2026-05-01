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

var enrichEnrichURL = cli.Command{
	Name:    "enrich-url",
	Usage:   "**Deprecated** — use POST /v1/lookup instead.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL of the product to enrich",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleEnrichEnrichURL,
	HideHelpCommand: true,
}

func handleEnrichEnrichURL(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.EnrichEnrichURLParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Enrich.EnrichURL(ctx, params, options...)
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
		Title:          "enrich enrich-url",
		Transform:      transform,
	})
}
