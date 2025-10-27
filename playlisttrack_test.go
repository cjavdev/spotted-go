// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/internal/testutil"
	"github.com/stainless-sdks/spotted-go/option"
)

func TestPlaylistTrackUpdateWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Playlists.Tracks.Update(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackUpdateParams{
			QueryUris:    spotted.String("uris"),
			InsertBefore: spotted.Int(3),
			RangeLength:  spotted.Int(2),
			RangeStart:   spotted.Int(1),
			SnapshotID:   spotted.String("snapshot_id"),
			BodyUris:     []string{"string"},
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Playlists.Tracks.Add(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackAddParams{
			QueryPosition: spotted.Int(0),
			QueryUris:     spotted.String("spotify:track:4iV5W9uYEdYUVa79Axb7Rh,spotify:track:1301WleyT98MSxVHPZCA6M"),
			BodyPosition:  spotted.Int(0),
			BodyUris:      []string{"string"},
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Playlists.Tracks.Remove(
		context.TODO(),
		"3cEYpjA9oz9GiPac4AsH4n",
		spotted.PlaylistTrackRemoveParams{
			Tracks: []spotted.PlaylistTrackRemoveParamsTrack{{
				Uri: spotted.String("uri"),
			}},
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
