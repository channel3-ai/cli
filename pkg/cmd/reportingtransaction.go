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

var reportingTransactionsList = cli.Command{
	Name:    "list",
	Usage:   "List affiliate transactions for your account over a datetime window.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "Inclusive end of the window (ISO 8601 datetime with optional offset, e.g. 2026-08-01T23:59:59-04:00). Offset-aware values are converted to UTC; naive values are treated as UTC.",
			QueryPath: "end_date",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page (max 100).",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number (1-indexed).",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Inclusive start of the window (ISO 8601 datetime with optional offset, e.g. 2026-08-01T00:00:00-04:00). Offset-aware values are converted to UTC; naive values are treated as UTC.",
			QueryPath: "start_date",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleReportingTransactionsList,
	HideHelpCommand: true,
}

func handleReportingTransactionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ReportingTransactionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Reporting.Transactions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "reporting:transactions list",
			Transform:      transform,
		})
	} else {
		iter := client.Reporting.Transactions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "reporting:transactions list",
			Transform:      transform,
		})
	}
}
