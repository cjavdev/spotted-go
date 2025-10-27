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

func TestRecommendationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Recommendations.Get(context.TODO(), spotted.RecommendationGetParams{
		Limit:                  spotted.Int(10),
		Market:                 spotted.String("ES"),
		MaxAcousticness:        spotted.Float(0),
		MaxDanceability:        spotted.Float(0),
		MaxDurationMs:          spotted.Int(0),
		MaxEnergy:              spotted.Float(0),
		MaxInstrumentalness:    spotted.Float(0),
		MaxKey:                 spotted.Int(0),
		MaxLiveness:            spotted.Float(0),
		MaxLoudness:            spotted.Float(0),
		MaxMode:                spotted.Int(0),
		MaxPopularity:          spotted.Int(0),
		MaxSpeechiness:         spotted.Float(0),
		MaxTempo:               spotted.Float(0),
		MaxTimeSignature:       spotted.Int(0),
		MaxValence:             spotted.Float(0),
		MinAcousticness:        spotted.Float(0),
		MinDanceability:        spotted.Float(0),
		MinDurationMs:          spotted.Int(0),
		MinEnergy:              spotted.Float(0),
		MinInstrumentalness:    spotted.Float(0),
		MinKey:                 spotted.Int(0),
		MinLiveness:            spotted.Float(0),
		MinLoudness:            spotted.Float(0),
		MinMode:                spotted.Int(0),
		MinPopularity:          spotted.Int(0),
		MinSpeechiness:         spotted.Float(0),
		MinTempo:               spotted.Float(0),
		MinTimeSignature:       spotted.Int(11),
		MinValence:             spotted.Float(0),
		SeedArtists:            spotted.String("4NHQUGzhtTLFvgF5SZesLK"),
		SeedGenres:             spotted.String("classical,country"),
		SeedTracks:             spotted.String("0c6xIDDpzE81m2q797ordA"),
		TargetAcousticness:     spotted.Float(0),
		TargetDanceability:     spotted.Float(0),
		TargetDurationMs:       spotted.Int(0),
		TargetEnergy:           spotted.Float(0),
		TargetInstrumentalness: spotted.Float(0),
		TargetKey:              spotted.Int(0),
		TargetLiveness:         spotted.Float(0),
		TargetLoudness:         spotted.Float(0),
		TargetMode:             spotted.Int(0),
		TargetPopularity:       spotted.Int(0),
		TargetSpeechiness:      spotted.Float(0),
		TargetTempo:            spotted.Float(0),
		TargetTimeSignature:    spotted.Int(0),
		TargetValence:          spotted.Float(0),
	})
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecommendationListAvailableGenreSeeds(t *testing.T) {
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
	_, err := client.Recommendations.ListAvailableGenreSeeds(context.TODO())
	if err != nil {
		var apierr *spotted.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
