// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/public-sdk-cli/internal/apiquery"
	"github.com/stainless-sdks/public-sdk-cli/internal/requestflag"
	"github.com/stainless-sdks/public-sdk-go"
	"github.com/stainless-sdks/public-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var priceTrackingListSubscriptions = cli.Command{
	Name:    "list-subscriptions",
	Usage:   "List your active price tracking subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "Pagination cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Max results (1-100)",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePriceTrackingListSubscriptions,
	HideHelpCommand: true,
}

var priceTrackingRetrieveHistory = cli.Command{
	Name:    "retrieve-history",
	Usage:   "Get price history for a canonical product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "canonical-product-id",
			Required:  true,
			PathParam: "canonical_product_id",
		},
		&requestflag.Flag[int64]{
			Name:      "days",
			Usage:     "Number of days of history to fetch (max 30)",
			Default:   30,
			QueryPath: "days",
		},
	},
	Action:          handlePriceTrackingRetrieveHistory,
	HideHelpCommand: true,
}

var priceTrackingStart = cli.Command{
	Name:    "start",
	Usage:   "Start tracking prices for a canonical product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "canonical-product-id",
			Required: true,
			BodyPath: "canonical_product_id",
		},
	},
	Action:          handlePriceTrackingStart,
	HideHelpCommand: true,
}

var priceTrackingStop = cli.Command{
	Name:    "stop",
	Usage:   "Stop tracking prices for a canonical product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "canonical-product-id",
			Required: true,
			BodyPath: "canonical_product_id",
		},
	},
	Action:          handlePriceTrackingStop,
	HideHelpCommand: true,
}

func handlePriceTrackingListSubscriptions(ctx context.Context, cmd *cli.Command) error {
	client := publicsdk.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := publicsdk.PriceTrackingListSubscriptionsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.PriceTracking.ListSubscriptions(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "price-tracking list-subscriptions",
			Transform:      transform,
		})
	} else {
		iter := client.PriceTracking.ListSubscriptionsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "price-tracking list-subscriptions",
			Transform:      transform,
		})
	}
}

func handlePriceTrackingRetrieveHistory(ctx context.Context, cmd *cli.Command) error {
	client := publicsdk.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("canonical-product-id") && len(unusedArgs) > 0 {
		cmd.Set("canonical-product-id", unusedArgs[0])
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

	params := publicsdk.PriceTrackingGetHistoryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PriceTracking.GetHistory(
		ctx,
		cmd.Value("canonical-product-id").(string),
		params,
		options...,
	)
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
		Title:          "price-tracking retrieve-history",
		Transform:      transform,
	})
}

func handlePriceTrackingStart(ctx context.Context, cmd *cli.Command) error {
	client := publicsdk.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := publicsdk.PriceTrackingStartParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PriceTracking.Start(ctx, params, options...)
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
		Title:          "price-tracking start",
		Transform:      transform,
	})
}

func handlePriceTrackingStop(ctx context.Context, cmd *cli.Command) error {
	client := publicsdk.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := publicsdk.PriceTrackingStopParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PriceTracking.Stop(ctx, params, options...)
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
		Title:          "price-tracking stop",
		Transform:      transform,
	})
}
