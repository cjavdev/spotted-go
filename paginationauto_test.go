// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/internal/testutil"
	"github.com/stainless-sdks/spotted-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Shows.ListEpisodesAutoPaging(
		context.TODO(),
		"showid",
		spotted.ShowListEpisodesParams{
			Limit:  spotted.Int(5),
			Offset: spotted.Int(10),
		},
	)
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		show := iter.Current()
		t.Logf("%+v\n", show.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
