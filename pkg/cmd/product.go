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
		&requestflag.Flag[*string]{
			Name:      "length-unit",
			Usage:     "Preferred unit for length dimensions (length/width/height). When unset, dimensions are returned in the unit the merchant stated.",
			QueryPath: "length_unit",
		},
		&requestflag.Flag[any]{
			Name:      "website-id",
			Usage:     `Optional list of website IDs to constrain the buy URL to, relevant if multiple merchants exist. Accepts website IDs or domains (e.g. "nike.com").`,
			QueryPath: "website_ids",
		},
		&requestflag.Flag[*string]{
			Name:      "weight-unit",
			Usage:     "Preferred unit for weight dimensions. When unset, weight is returned in the unit the merchant stated.",
			QueryPath: "weight_unit",
		},
		&requestflag.Flag[string]{
			Name:       "x-user-id",
			Usage:      "Optional user identifier to attribute clicks and sales to a user in your system. Channel3 appends it to buy URLs in the response.",
			HeaderPath: "x-user-id",
		},
	},
	Action:          handleProductsRetrieve,
	HideHelpCommand: true,
}

var productsBrowse = requestflag.WithInnerFlags(cli.Command{
	Name:    "browse",
	Usage:   "List and page through products for a set of filters.",
	Suggest: true,
	Flags: []cli.Flag{
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
			Usage:    "Opaque token from a previous browse response to fetch the next page.",
			BodyPath: "page_token",
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
	Action:          handleProductsBrowse,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filters": {
		&requestflag.InnerFlag[any]{
			Name:       "filters.age",
			Usage:      "Filter by age group. Age-agnostic products are treated as adult products.",
			InnerField: "age",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "filters.attributes",
			Usage:      "If provided, only products matching these key/value constraints will be returned. Keys are attribute handles (e.g. 'color', 'material') and values are lists of allowed values (OR within a key, AND across keys). When a category filter is also supplied, all keys must be valid attributes of at least one of the requested categories. See `Category.attributes` for the valid keys and values per category.",
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
	Action:          handleProductsFindSimilar,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code (plus the pan-region ``EU``).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.length-unit",
			Usage:      "Preferred unit for length dimensions (length/width/height) in responses. A request dimension filter's unit for the field takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "length_unit",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.weight-unit",
			Usage:      "Preferred unit for weight dimensions in responses. A request dimension filter's weight unit takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "weight_unit",
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
			Usage:      "If provided, only products matching these key/value constraints will be returned. Keys are attribute handles (e.g. 'color', 'material') and values are lists of allowed values (OR within a key, AND across keys). When a category filter is also supplied, all keys must be valid attributes of at least one of the requested categories. See `Category.attributes` for the valid keys and values per category.",
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
		&requestflag.Flag[string]{
			Name:       "x-user-id",
			Usage:      "Optional user identifier to attribute clicks and sales to a user in your system. Channel3 appends it to buy URLs in the response.",
			HeaderPath: "x-user-id",
		},
	},
	Action:          handleProductsLookup,
	HideHelpCommand: true,
}

var productsMonetize = cli.Command{
	Name:    "monetize",
	Usage:   "Return monetizable offers (with max commission rate) for a product URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "The URL of the product to monetize",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:       "x-user-id",
			Usage:      "Optional user identifier to attribute clicks and sales to a user in your system. Channel3 appends it to buy URLs in the response.",
			HeaderPath: "x-user-id",
		},
	},
	Action:          handleProductsMonetize,
	HideHelpCommand: true,
}

var productsSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Search for products with pagination support.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "base64-image",
			Usage:    "Base64 encoded image. At least one of `query`, `image_url`, `base64_image`, or `page_token` must be provided.",
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
			Usage:    "Image URL. At least one of `query`, `image_url`, `base64_image`, or `page_token` must be provided.",
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
			Usage:    "Search query. At least one of `query`, `image_url`, `base64_image`, or `page_token` must be provided.",
			BodyPath: "query",
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
	Action:          handleProductsSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code (plus the pan-region ``EU``).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.length-unit",
			Usage:      "Preferred unit for length dimensions (length/width/height) in responses. A request dimension filter's unit for the field takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "length_unit",
		},
		&requestflag.InnerFlag[string]{
			Name:       "config.mode",
			Usage:      "Search strategy. `default` (recommended) combines lexical + semantic search and is right for most use cases. `keyword` is lexical only — use it for real-time, low-latency needs like ad targeting. `agentic` uses an LLM to plan multiple structured sub-searches for complex queries, with higher latency than the other modes.",
			InnerField: "mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.weight-unit",
			Usage:      "Preferred unit for weight dimensions in responses. A request dimension filter's weight unit takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "weight_unit",
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
			Usage:      "If provided, only products matching these key/value constraints will be returned. Keys are attribute handles (e.g. 'color', 'material') and values are lists of allowed values (OR within a key, AND across keys). When a category filter is also supplied, all keys must be valid attributes of at least one of the requested categories. See `Category.attributes` for the valid keys and values per category.",
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
		&requestflag.Flag[*string]{
			Name:     "segment",
			Usage:    `Image segmentation mode. None (default) disables segmentation. "AUTO" segments and crops the main product automatically. A custom string (e.g. "shoe", "mug") segments the specified object.`,
			BodyPath: "segment",
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
	Action:          handleProductsSearchByImage,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.country",
			Usage:      "ISO 3166-1 alpha-2 country code (plus the pan-region ``EU``).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.currency",
			Usage:      "ISO 4217 currency code.",
			InnerField: "currency",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.language",
			Usage:      "ISO 639-1 language code.",
			InnerField: "language",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.length-unit",
			Usage:      "Preferred unit for length dimensions (length/width/height) in responses. A request dimension filter's unit for the field takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "length_unit",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.weight-unit",
			Usage:      "Preferred unit for weight dimensions in responses. A request dimension filter's weight unit takes precedence; when neither is set, the merchant's stated unit is returned.",
			InnerField: "weight_unit",
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
			Usage:      "If provided, only products matching these key/value constraints will be returned. Keys are attribute handles (e.g. 'color', 'material') and values are lists of allowed values (OR within a key, AND across keys). When a category filter is also supplied, all keys must be valid attributes of at least one of the requested categories. See `Category.attributes` for the valid keys and values per category.",
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

func handleProductsBrowse(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductBrowseParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Products.Browse(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products browse",
			Transform:      transform,
		})
	} else {
		iter := client.Products.BrowseAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "products browse",
			Transform:      transform,
		})
	}
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

func handleProductsMonetize(ctx context.Context, cmd *cli.Command) error {
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

	params := channel3go.ProductMonetizeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Products.Monetize(ctx, params, options...)
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
		Title:          "products monetize",
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
