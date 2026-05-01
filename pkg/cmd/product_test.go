// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/channel3-ai/cli/internal/mocktest"
	"github.com/channel3-ai/cli/internal/requestflag"
)

func TestProductsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "retrieve",
			"--product-id", "product_id",
			"--country", "US",
			"--currency", "USD",
			"--language", "en",
			"--website-id", "[string]",
		)
	})
}

func TestProductsFindSimilar(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "find-similar",
			"--max-items", "10",
			"--product-id", "product_id",
			"--config", "{country: US, currency: USD, language: en}",
			"--filters", "{age: [newborn], availability: [InStock], brand_ids: [string], category_ids: [string], condition: new, exclude_brand_ids: [string], exclude_category_ids: [string], exclude_website_ids: [string], gender: male, price: {max_price: 0, min_price: 0}, website_ids: [string]}",
			"--limit", "1",
			"--page-token", "page_token",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productsFindSimilar)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "find-similar",
			"--max-items", "10",
			"--product-id", "product_id",
			"--config.country", "US",
			"--config.currency", "USD",
			"--config.language", "en",
			"--filters.age", "[newborn]",
			"--filters.availability", "[InStock]",
			"--filters.brand-ids", "[string]",
			"--filters.category-ids", "[string]",
			"--filters.condition", "new",
			"--filters.exclude-brand-ids", "[string]",
			"--filters.exclude-category-ids", "[string]",
			"--filters.exclude-website-ids", "[string]",
			"--filters.gender", "male",
			"--filters.price", "{max_price: 0, min_price: 0}",
			"--filters.website-ids", "[string]",
			"--limit", "1",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"product_id: product_id\n" +
			"config:\n" +
			"  country: US\n" +
			"  currency: USD\n" +
			"  language: en\n" +
			"filters:\n" +
			"  age:\n" +
			"    - newborn\n" +
			"  availability:\n" +
			"    - InStock\n" +
			"  brand_ids:\n" +
			"    - string\n" +
			"  category_ids:\n" +
			"    - string\n" +
			"  condition: new\n" +
			"  exclude_brand_ids:\n" +
			"    - string\n" +
			"  exclude_category_ids:\n" +
			"    - string\n" +
			"  exclude_website_ids:\n" +
			"    - string\n" +
			"  gender: male\n" +
			"  price:\n" +
			"    max_price: 0\n" +
			"    min_price: 0\n" +
			"  website_ids:\n" +
			"    - string\n" +
			"limit: 1\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"products", "find-similar",
			"--max-items", "10",
		)
	})
}

func TestProductsLookup(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "lookup",
			"--url", "url",
			"--max-staleness-hours", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"url: url\n" +
			"max_staleness_hours: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"products", "lookup",
		)
	})
}

func TestProductsSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "search",
			"--max-items", "10",
			"--base64-image", "base64_image",
			"--config", "{country: US, currency: USD, keyword_search_only: true, language: en}",
			"--filters", "{age: [newborn], availability: [InStock], brand_ids: [string], category_ids: [string], condition: new, exclude_brand_ids: [string], exclude_category_ids: [string], exclude_website_ids: [string], gender: male, price: {max_price: 0, min_price: 0}, website_ids: [string]}",
			"--image-url", "image_url",
			"--limit", "1",
			"--page-token", "page_token",
			"--query", "query",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productsSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "search",
			"--max-items", "10",
			"--base64-image", "base64_image",
			"--config.country", "US",
			"--config.currency", "USD",
			"--config.keyword-search-only=true",
			"--config.language", "en",
			"--filters.age", "[newborn]",
			"--filters.availability", "[InStock]",
			"--filters.brand-ids", "[string]",
			"--filters.category-ids", "[string]",
			"--filters.condition", "new",
			"--filters.exclude-brand-ids", "[string]",
			"--filters.exclude-category-ids", "[string]",
			"--filters.exclude-website-ids", "[string]",
			"--filters.gender", "male",
			"--filters.price", "{max_price: 0, min_price: 0}",
			"--filters.website-ids", "[string]",
			"--image-url", "image_url",
			"--limit", "1",
			"--page-token", "page_token",
			"--query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"base64_image: base64_image\n" +
			"config:\n" +
			"  country: US\n" +
			"  currency: USD\n" +
			"  keyword_search_only: true\n" +
			"  language: en\n" +
			"filters:\n" +
			"  age:\n" +
			"    - newborn\n" +
			"  availability:\n" +
			"    - InStock\n" +
			"  brand_ids:\n" +
			"    - string\n" +
			"  category_ids:\n" +
			"    - string\n" +
			"  condition: new\n" +
			"  exclude_brand_ids:\n" +
			"    - string\n" +
			"  exclude_category_ids:\n" +
			"    - string\n" +
			"  exclude_website_ids:\n" +
			"    - string\n" +
			"  gender: male\n" +
			"  price:\n" +
			"    max_price: 0\n" +
			"    min_price: 0\n" +
			"  website_ids:\n" +
			"    - string\n" +
			"image_url: image_url\n" +
			"limit: 1\n" +
			"page_token: page_token\n" +
			"query: query\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"products", "search",
			"--max-items", "10",
		)
	})
}

func TestProductsSearchByImage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "search-by-image",
			"--max-items", "10",
			"--base64-image", "base64_image",
			"--config", "{country: US, currency: USD, language: en}",
			"--filters", "{age: [newborn], availability: [InStock], brand_ids: [string], category_ids: [string], condition: new, exclude_brand_ids: [string], exclude_category_ids: [string], exclude_website_ids: [string], gender: male, price: {max_price: 0, min_price: 0}, website_ids: [string]}",
			"--image-url", "image_url",
			"--limit", "1",
			"--page-token", "page_token",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(productsSearchByImage)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"products", "search-by-image",
			"--max-items", "10",
			"--base64-image", "base64_image",
			"--config.country", "US",
			"--config.currency", "USD",
			"--config.language", "en",
			"--filters.age", "[newborn]",
			"--filters.availability", "[InStock]",
			"--filters.brand-ids", "[string]",
			"--filters.category-ids", "[string]",
			"--filters.condition", "new",
			"--filters.exclude-brand-ids", "[string]",
			"--filters.exclude-category-ids", "[string]",
			"--filters.exclude-website-ids", "[string]",
			"--filters.gender", "male",
			"--filters.price", "{max_price: 0, min_price: 0}",
			"--filters.website-ids", "[string]",
			"--image-url", "image_url",
			"--limit", "1",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"base64_image: base64_image\n" +
			"config:\n" +
			"  country: US\n" +
			"  currency: USD\n" +
			"  language: en\n" +
			"filters:\n" +
			"  age:\n" +
			"    - newborn\n" +
			"  availability:\n" +
			"    - InStock\n" +
			"  brand_ids:\n" +
			"    - string\n" +
			"  category_ids:\n" +
			"    - string\n" +
			"  condition: new\n" +
			"  exclude_brand_ids:\n" +
			"    - string\n" +
			"  exclude_category_ids:\n" +
			"    - string\n" +
			"  exclude_website_ids:\n" +
			"    - string\n" +
			"  gender: male\n" +
			"  price:\n" +
			"    max_price: 0\n" +
			"    min_price: 0\n" +
			"  website_ids:\n" +
			"    - string\n" +
			"image_url: image_url\n" +
			"limit: 1\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"products", "search-by-image",
			"--max-items", "10",
		)
	})
}
