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

func TestPlaylistTrackUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Playlists.Tracks.Update(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackUpdateParams{
			InsertBefore: spotted.Int(3),
			Published:    spotted.Bool(true),
			RangeLength:  spotted.Int(2),
			RangeStart:   spotted.Int(1),
			SnapshotID:   spotted.String("snapshot_id"),
			Uris:         []string{"string"},
		},
	)
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlaylistTrackListWithOptionalParams(t *testing.T) {
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
	_, err := client.Playlists.Tracks.List(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackListParams{
			AdditionalTypes: spotted.String("additional_types"),
			Fields:          spotted.String("items(added_by.id,track(name,href,album(name,href)))"),
			Limit:           spotted.Int(10),
			Market:          spotted.String("ES"),
			Offset:          spotted.Int(5),
		},
	)
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlaylistTrackAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Playlists.Tracks.Add(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackAddParams{
			Position:  spotted.Int(0),
			Published: spotted.Bool(true),
			Uris:      []string{"string"},
		},
	)
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlaylistTrackRemoveWithOptionalParams(t *testing.T) {
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
	_, err := client.Playlists.Tracks.Remove(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackRemoveParams{
			Tracks: []spotted.PlaylistTrackRemoveParamsTrack{{
				Uri: spotted.String("uri"),
			}},
			Published:  spotted.Bool(true),
			SnapshotID: spotted.String("snapshot_id"),
		},
	)
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
