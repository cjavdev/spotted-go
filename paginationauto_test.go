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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	iter := client.Shows.ListEpisodesAutoPaging(
		context.TODO(),
		"showid",
		spotted.ShowListEpisodesParams{
			Limit:  spotted.Int(10),
			Offset: spotted.Int(20),
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
