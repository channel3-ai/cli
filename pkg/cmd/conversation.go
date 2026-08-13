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

var conversationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Run one conversation turn. Omit `conversation_id` to create the thread with this\nturn; pass it to continue an existing thread.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "message",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "context",
			Usage:    "Partner-supplied context pinned to the top of a conversation thread.",
			BodyPath: "context",
		},
		&requestflag.Flag[*string]{
			Name:     "conversation-id",
			Usage:    "Existing thread to continue. When omitted, a new thread is created and its id returned.",
			BodyPath: "conversation_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Search filters for the search API.",
			BodyPath: "filters",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "Stream turn events over SSE (default) or return the assembled turn as JSON.",
			Default:  true,
			BodyPath: "stream",
		},
		&requestflag.Flag[string]{
			Name:       "x-user-id",
			Usage:      "Optional user identifier to attribute clicks and sales to a user in your system. Channel3 appends it to buy URLs in the response.",
			HeaderPath: "x-user-id",
		},
	},
	Action:          handleConversationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"message": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "message.parts",
			InnerField: "parts",
		},
		&requestflag.InnerFlag[string]{
			Name:       "message.role",
			Usage:      `Allowed values: "user".`,
			InnerField: "role",
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
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.age",
			Usage:      "Filter by age group. Age-agnostic products are treated as adult products.",
			InnerField: "age",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.attributes",
			Usage:      "If provided, only products whose extracted attributes match these key/value constraints will be returned. Keys are attribute handles (e.g. 'color', 'material') and values are lists of allowed values (OR within a key, AND across keys). When a category filter is also supplied, all keys must be valid attributes of at least one of the requested categories. See `Category.attributes` for the valid keys/values per category.",
			InnerField: "attributes",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filters.availability",
			Usage:      "Offer availability statuses to match (OR). Defaults to ['InStock']. An offer with no availability data counts as 'InStock'. Pass every value to disable availability filtering.",
			InnerField: "availability",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.brand-ids",
			Usage:      "If provided, only products from these brands will be returned",
			InnerField: "brand_ids",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.category-ids",
			Usage:      "If provided, only products from these categories will be returned. Accepts category slugs.",
			InnerField: "category_ids",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.colors",
			Usage:      "[Beta] Color filter wrapper. Holds required colors and optional match mode.",
			InnerField: "colors",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filters.conditions",
			Usage:      "Offer conditions to match (OR). Defaults to ['new'], which also matches offers whose condition is unknown. Pass every value to disable condition filtering.",
			InnerField: "conditions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.dimensions",
			Usage:      "Physical-dimension range filters, matched against the same offer.\n\nMatching products have at least one offer satisfying every provided\nrange (alongside any locale/price/availability filters). Values are\ncompared with a small relative tolerance. An offer with no dimension data\nfor a filtered field does not match; note that when a single merchant on a\nproduct reports a dimension it is shared across that product's offers, so a\nmatching offer may not itself surface that dimension in the response.",
			InnerField: "dimensions",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.exclude-brand-ids",
			Usage:      "If provided, products from these brands will be excluded from the results",
			InnerField: "exclude_brand_ids",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.exclude-category-ids",
			Usage:      "If provided, products in these categories (or their descendants) will be excluded from the results. Accepts category slugs.",
			InnerField: "exclude_category_ids",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.exclude-website-ids",
			Usage:      `If provided, products from these websites will be excluded from the results. Accepts website IDs or domains (e.g. "nike.com").`,
			InnerField: "exclude_website_ids",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filters.gender",
			Usage:      "Product gender. 'unisex' is deprecated: coerced to None on input, never emitted.",
			InnerField: "gender",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.price",
			Usage:      "Price filter for search. Values are inclusive.",
			InnerField: "price",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filters.sale",
			Usage:      "If 'on_sale', only products with at least one on-sale offer (priced below its compare-at price) for the requested locale are returned. If omitted, no filter.",
			InnerField: "sale",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.website-ids",
			Usage:      `If provided, only products from these websites will be returned. Accepts website IDs or domains (e.g. "nike.com").`,
			InnerField: "website_ids",
		},
	},
})

var conversationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Thread metadata plus one page of its message history. Paginate `items` with\n`limit` and `cursor`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "conversation-id",
			Required:  true,
			PathParam: "conversation_id",
		},
		&requestflag.Flag[*string]{
			Name:      "cursor",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
	},
	Action:          handleConversationsRetrieve,
	HideHelpCommand: true,
}

func handleConversationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ConversationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Conversations.New(ctx, params, options...)
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
		Title:          "conversations create",
		Transform:      transform,
	})
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

	params := channel3go.ConversationGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Conversations.Get(
		ctx,
		cmd.Value("conversation-id").(string),
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
		Title:          "conversations retrieve",
		Transform:      transform,
	})
}
