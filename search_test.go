// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/internal/testutil"
	"github.com/cjavdev/spotted-go/option"
)

func TestSearchQueryWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := spotted.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
	)
	_, err := client.Search.Query(context.TODO(), spotted.SearchQueryParams{
		Q:               "remaster%20track:Doxy%20artist:Miles%20Davis",
		Type:            []string{"album"},
		IncludeExternal: spotted.SearchQueryParamsIncludeExternalAudio,
		Limit:           spotted.Int(10),
		Market:          spotted.String("ES"),
		Offset:          spotted.Int(5),
	})
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
