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

var productsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get detailed information about a specific product by its ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "product-id",
			Required:  true,
			PathParam: "product_id",
		},
		&requestflag.Flag[*string]{
			Name:      "country",
			Usage:     "ISO 3166-1 alpha-2 country code. Matches any country when unset; defaults to 'US' only when language and currency are also unset.",
			QueryPath: "country",
		},
		&requestflag.Flag[*string]{
			Name:      "currency",
			Usage:     "ISO 4217 currency code. When unset, inferred from `country` (e.g. GB -> GBP); falls back to 'USD' only when all three locale fields are unset.",
			QueryPath: "currency",
		},
		&requestflag.Flag[*string]{
			Name:      "language",
			Usage:     "ISO 639-1 language code. Matches any language when unset; defaults to 'en' only when country and currency are also unset.",
			QueryPath: "language",
		},
		&requestflag.Flag[any]{
			Name:      "website-id",
			Usage:     `Optional list of website IDs to constrain the buy URL to, relevant if multiple merchants exist. Accepts website IDs or domains (e.g. "nike.com").`,
			QueryPath: "website_ids",
		},
	},
	Action:          handleProductsRetrieve,
	HideHelpCommand: true,
}

var productsFindSimilar = requestflag.WithInnerFlags(cli.Command{
	Name:    "find-similar",
	Usage:   "Find products similar to a given product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "Canonical product ID to find similar products for.",
			Required: true,
			BodyPath: "product_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Locale options for API requests.\n\nLocale fields are optional; the server infers missing values. Details are\non ``language``, ``country``, and ``currency`` below.",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Search filters for the search API.",
			BodyPath: "filters",
		},
		&requestflag.Flag[*int64]{
			Name:     "limit",
			Usage:    "Optional limit on the number of results. Default is 20, max is 30.",
			Default:  requestflag.Ptr[int64](20),
			BodyPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "Opaque token from a previous similar response to fetch the next page of results.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleProductsFindSimilar,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g. ``currency=EUR`` with no specific country).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code. When unset, inferred from ``country`` (e.g. ``GB`` → ``GBP``), defaulting to ``USD``.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code. When unset, inferred from ``country`` (preferred) then ``currency``, defaulting to ``en``.",
			InnerField: "language",
		},
	},
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.age",
			Usage:      "Filter by age group. Age-agnostic products are treated as adult products.",
			InnerField: "age",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.availability",
			Usage:      "If provided, only products with these availability statuses will be returned",
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
		&requestflag.InnerFlag[*string]{
			Name:       "filters.condition",
			Usage:      "Filter by product condition. Incubating: condition data is currently incomplete; products without condition data will be included in all condition filter results.",
			InnerField: "condition",
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
			Usage:      `Allowed values: "male", "female".`,
			InnerField: "gender",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.price",
			Usage:      "Price filter for search. Values are inclusive.",
			InnerField: "price",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.website-ids",
			Usage:      `If provided, only products from these websites will be returned. Accepts website IDs or domains (e.g. "nike.com").`,
			InnerField: "website_ids",
		},
	},
})

var productsLookup = cli.Command{
	Name:    "lookup",
	Usage:   "Retrieve product information for any supported product URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL of the product to look up",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[int64]{
			Name:     "max-staleness-hours",
			Usage:    "Maximum age (in hours) of cached product data before forcing a fresh lookup. Defaults to 3 hours.",
			Default:  3,
			BodyPath: "max_staleness_hours",
		},
	},
	Action:          handleProductsLookup,
	HideHelpCommand: true,
}

var productsSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search for products with pagination support.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "base64-image",
			Usage:    "Base64 encoded image. At least one of `query`, `image_url`, or `base64_image` must be provided.",
			BodyPath: "base64_image",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Search and locale options for a search request.",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Search filters for the search API.",
			BodyPath: "filters",
		},
		&requestflag.Flag[*string]{
			Name:     "image-url",
			Usage:    "Image URL. At least one of `query`, `image_url`, or `base64_image` must be provided.",
			BodyPath: "image_url",
		},
		&requestflag.Flag[*int64]{
			Name:     "limit",
			Usage:    "Optional limit on the number of results. Default is 20, max is 30.",
			Default:  requestflag.Ptr[int64](20),
			BodyPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "Opaque token from a previous search response to fetch the next page of results.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:     "query",
			Usage:    "Search query. At least one of `query`, `image_url`, or `base64_image` must be provided.",
			BodyPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleProductsSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g. ``currency=EUR`` with no specific country).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code. When unset, inferred from ``country`` (e.g. ``GB`` → ``GBP``), defaulting to ``USD``.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.keyword-search-only",
			Usage:      "If True, search will only use keyword search and not vector search. Keyword-only search is not supported with image input.",
			InnerField: "keyword_search_only",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code. When unset, inferred from ``country`` (preferred) then ``currency``, defaulting to ``en``.",
			InnerField: "language",
		},
	},
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.age",
			Usage:      "Filter by age group. Age-agnostic products are treated as adult products.",
			InnerField: "age",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.availability",
			Usage:      "If provided, only products with these availability statuses will be returned",
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
		&requestflag.InnerFlag[*string]{
			Name:       "filters.condition",
			Usage:      "Filter by product condition. Incubating: condition data is currently incomplete; products without condition data will be included in all condition filter results.",
			InnerField: "condition",
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
			Usage:      `Allowed values: "male", "female".`,
			InnerField: "gender",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.price",
			Usage:      "Price filter for search. Values are inclusive.",
			InnerField: "price",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.website-ids",
			Usage:      `If provided, only products from these websites will be returned. Accepts website IDs or domains (e.g. "nike.com").`,
			InnerField: "website_ids",
		},
	},
})

var productsSearchByImage = requestflag.WithInnerFlags(cli.Command{
	Name:    "search-by-image",
	Usage:   "Search the catalog by image (URL or base64), with pagination support.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "base64-image",
			Usage:    "Base64 encoded image bytes (no data URI prefix).",
			BodyPath: "base64_image",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Locale options for API requests.\n\nLocale fields are optional; the server infers missing values. Details are\non ``language``, ``country``, and ``currency`` below.",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filters",
			Usage:    "Search filters for the search API.",
			BodyPath: "filters",
		},
		&requestflag.Flag[*string]{
			Name:     "image-url",
			Usage:    "Publicly accessible URL of the image to search with.",
			BodyPath: "image_url",
		},
		&requestflag.Flag[*int64]{
			Name:     "limit",
			Usage:    "Optional limit on the number of results. Default is 20, max is 30.",
			Default:  requestflag.Ptr[int64](20),
			BodyPath: "limit",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "Opaque token from a previous image-search response to fetch the next page of results.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleProductsSearchByImage,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g. ``currency=EUR`` with no specific country).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code. When unset, inferred from ``country`` (e.g. ``GB`` → ``GBP``), defaulting to ``USD``.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code. When unset, inferred from ``country`` (preferred) then ``currency``, defaulting to ``en``.",
			InnerField: "language",
		},
	},
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.age",
			Usage:      "Filter by age group. Age-agnostic products are treated as adult products.",
			InnerField: "age",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.availability",
			Usage:      "If provided, only products with these availability statuses will be returned",
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
		&requestflag.InnerFlag[*string]{
			Name:       "filters.condition",
			Usage:      "Filter by product condition. Incubating: condition data is currently incomplete; products without condition data will be included in all condition filter results.",
			InnerField: "condition",
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
			Usage:      `Allowed values: "male", "female".`,
			InnerField: "gender",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.price",
			Usage:      "Price filter for search. Values are inclusive.",
			InnerField: "price",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filters.website-ids",
			Usage:      `If provided, only products from these websites will be returned. Accepts website IDs or domains (e.g. "nike.com").`,
			InnerField: "website_ids",
		},
	},
})

func handleProductsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := channel3go.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("product-id") && len(unusedArgs) > 0 {
		cmd.Set("product-id", unusedArgs[0])
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

	params := channel3go.ProductGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Products.Get(
		ctx,
		cmd.Value("product-id").(string),
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
		Title:          "products retrieve",
		Transform:      transform,
	})
}

func handleProductsFindSimilar(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductFindSimilarParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Products.FindSimilar(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products find-similar",
			Transform:      transform,
		})
	} else {
		iter := client.Products.FindSimilarAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products find-similar",
			Transform:      transform,
		})
	}
}

func handleProductsLookup(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductLookupParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Products.Lookup(ctx, params, options...)
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
		Title:          "products lookup",
		Transform:      transform,
	})
}

func handleProductsSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductSearchParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Products.Search(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products search",
			Transform:      transform,
		})
	} else {
		iter := client.Products.SearchAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products search",
			Transform:      transform,
		})
	}
}

func handleProductsSearchByImage(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductSearchByImageParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Products.SearchByImage(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products search-by-image",
			Transform:      transform,
		})
	} else {
		iter := client.Products.SearchByImageAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products search-by-image",
			Transform:      transform,
		})
	}
}
