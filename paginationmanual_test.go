// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted_test

import (
	"context"
	"os"
	"testing"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/internal/testutil"
	"github.com/cjavdev/spotted-go/option"
)

func TestManualPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	page, err := client.Shows.ListEpisodes(
		context.TODO(),
		"showid",
		spotted.ShowListEpisodesParams{
			Limit:  spotted.Int(10),
			Offset: spotted.Int(20),
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, show := range page.Items {
		t.Logf("%+v\n", show.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, show := range page.Items {
			t.Logf("%+v\n", show.ID)
		}
	}
}
