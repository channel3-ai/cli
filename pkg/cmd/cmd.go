// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/public-sdk-cli/internal/autocomplete"
	"github.com/stainless-sdks/public-sdk-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "public-sdk",
		Usage:     "CLI for the channel3 API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("CHANNEL3_API_KEY"),
			},
			&requestflag.Flag[string]{
				Name:    "language",
				Usage:   "Default ISO 639-1 language code applied to product calls (e.g. 'en'). Per-call config.language overrides this.",
				Sources: cli.EnvVars("CHANNEL3_LANGUAGE"),
			},
			&requestflag.Flag[string]{
				Name:    "country",
				Usage:   "Default ISO 3166-1 alpha-2 country code applied to product calls (e.g. 'GB'). Per-call values override this.",
				Sources: cli.EnvVars("CHANNEL3_COUNTRY"),
			},
			&requestflag.Flag[string]{
				Name:    "currency",
				Usage:   "Default ISO 4217 currency code applied to product calls (e.g. 'GBP'). Per-call values override this.",
				Sources: cli.EnvVars("CHANNEL3_CURRENCY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "products",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&productsRetrieve,
					&productsFindSimilar,
					&productsLookup,
					&productsSearch,
					&productsSearchByImage,
				},
			},
			{
				Name:     "brands",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&brandsRetrieve,
					&brandsList,
					&brandsFind,
					&brandsSearch,
				},
			},
			{
				Name:     "categories",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&categoriesRetrieve,
					&categoriesList,
					&categoriesSearch,
				},
			},
			{
				Name:     "websites",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&websitesRetrieve,
				},
			},
			{
				Name:     "price-tracking",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&priceTrackingListSubscriptions,
					&priceTrackingRetrieveHistory,
					&priceTrackingStart,
					&priceTrackingStop,
				},
			},
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchPerform,
				},
			},
			{
				Name:     "enrich",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&enrichEnrichURL,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "public-sdk @manpages [-o public-sdk.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "public-sdk.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "public-sdk.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
