// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// RecommendationService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecommendationService] method instead.
type RecommendationService struct {
	Options []option.RequestOption
}

// NewRecommendationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecommendationService(opts ...option.RequestOption) (r RecommendationService) {
	r = RecommendationService{}
	r.Options = opts
	return
}

// Recommendations are generated based on the available information for a given
// seed entity and matched against similar artists and tracks. If there is
// sufficient information about the provided seeds, a list of tracks will be
// returned together with pool size details.
//
// For artists and tracks that are very new or obscure there might not be enough
// data to generate a list of tracks.
//
// Deprecated: deprecated
func (r *RecommendationService) Get(ctx context.Context, query RecommendationGetParams, opts ...option.RequestOption) (res *RecommendationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "recommendations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a list of available genres seed parameter values for
// [recommendations](/documentation/web-api/reference/get-recommendations).
//
// Deprecated: deprecated
func (r *RecommendationService) ListAvailableGenreSeeds(ctx context.Context, opts ...option.RequestOption) (res *RecommendationListAvailableGenreSeedsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "recommendations/available-genre-seeds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RecommendationGetResponse struct {
	// An array of recommendation seed objects.
	Seeds []RecommendationGetResponseSeed `json:"seeds,required"`
	// An array of track objects ordered according to the parameters supplied.
	Tracks []shared.TrackObject `json:"tracks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Seeds       respjson.Field
		Tracks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecommendationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RecommendationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecommendationGetResponseSeed struct {
	// The id used to select this seed. This will be the same as the string used in the
	// `seed_artists`, `seed_tracks` or `seed_genres` parameter.
	ID string `json:"id"`
	// The number of tracks available after min\_\* and max\_\* filters have been
	// applied.
	AfterFilteringSize int64 `json:"afterFilteringSize"`
	// The number of tracks available after relinking for regional availability.
	AfterRelinkingSize int64 `json:"afterRelinkingSize"`
	// A link to the full track or artist data for this seed. For tracks this will be a
	// link to a Track Object. For artists a link to an Artist Object. For genre seeds,
	// this value will be `null`.
	Href string `json:"href"`
	// The number of recommended tracks available for this seed.
	InitialPoolSize int64 `json:"initialPoolSize"`
	// The entity type of this seed. One of `artist`, `track` or `genre`.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AfterFilteringSize respjson.Field
		AfterRelinkingSize respjson.Field
		Href               respjson.Field
		InitialPoolSize    respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecommendationGetResponseSeed) RawJSON() string { return r.JSON.raw }
func (r *RecommendationGetResponseSeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecommendationListAvailableGenreSeedsResponse struct {
	Genres []string `json:"genres,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Genres      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecommendationListAvailableGenreSeedsResponse) RawJSON() string { return r.JSON.raw }
func (r *RecommendationListAvailableGenreSeedsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecommendationGetParams struct {
	// The target size of the list of recommended tracks. For seeds with unusually
	// small pools or when highly restrictive filtering is applied, it may be
	// impossible to generate the requested number of recommended tracks. Debugging
	// information for such cases is available in the response. Default: 20\. Minimum:
	// 1\. Maximum: 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// An
	// [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// If a country code is specified, only content that is available in that market
	// will be returned.<br/> If a valid user access token is specified in the request
	// header, the country associated with the user account will take priority over
	// this parameter.<br/> _**Note**: If neither market or user country are provided,
	// the content is considered unavailable for the client._<br/> Users can view the
	// country that is associated with their account in the
	// [account settings](https://www.spotify.com/account/overview/).
	Market param.Opt[string] `query:"market,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxAcousticness param.Opt[float64] `query:"max_acousticness,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxDanceability param.Opt[float64] `query:"max_danceability,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxDurationMs param.Opt[int64] `query:"max_duration_ms,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxEnergy param.Opt[float64] `query:"max_energy,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxInstrumentalness param.Opt[float64] `query:"max_instrumentalness,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxKey param.Opt[int64] `query:"max_key,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxLiveness param.Opt[float64] `query:"max_liveness,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxLoudness param.Opt[float64] `query:"max_loudness,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxMode param.Opt[int64] `query:"max_mode,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxPopularity param.Opt[int64] `query:"max_popularity,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxSpeechiness param.Opt[float64] `query:"max_speechiness,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxTempo param.Opt[float64] `query:"max_tempo,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxTimeSignature param.Opt[int64] `query:"max_time_signature,omitzero" json:"-"`
	// For each tunable track attribute, a hard ceiling on the selected track
	// attribute’s value can be provided. See tunable track attributes below for the
	// list of available options. For example, `max_instrumentalness=0.35` would filter
	// out most tracks that are likely to be instrumental.
	MaxValence param.Opt[float64] `query:"max_valence,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinAcousticness param.Opt[float64] `query:"min_acousticness,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinDanceability param.Opt[float64] `query:"min_danceability,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinDurationMs param.Opt[int64] `query:"min_duration_ms,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinEnergy param.Opt[float64] `query:"min_energy,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinInstrumentalness param.Opt[float64] `query:"min_instrumentalness,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinKey param.Opt[int64] `query:"min_key,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinLiveness param.Opt[float64] `query:"min_liveness,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinLoudness param.Opt[float64] `query:"min_loudness,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinMode param.Opt[int64] `query:"min_mode,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinPopularity param.Opt[int64] `query:"min_popularity,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinSpeechiness param.Opt[float64] `query:"min_speechiness,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinTempo param.Opt[float64] `query:"min_tempo,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinTimeSignature param.Opt[int64] `query:"min_time_signature,omitzero" json:"-"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s
	// value can be provided. See tunable track attributes below for the list of
	// available options. For example, `min_tempo=140` would restrict results to only
	// those tracks with a tempo of greater than 140 beats per minute.
	MinValence param.Opt[float64] `query:"min_valence,omitzero" json:"-"`
	// A comma separated list of
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for seed
	// artists. Up to 5 seed values may be provided in any combination of
	// `seed_artists`, `seed_tracks` and `seed_genres`.<br/> _**Note**: only required
	// if `seed_genres` and `seed_tracks` are not set_.
	SeedArtists param.Opt[string] `query:"seed_artists,omitzero" json:"-"`
	// A comma separated list of any genres in the set of
	// [available genre seeds](/documentation/web-api/reference/get-recommendation-genres).
	// Up to 5 seed values may be provided in any combination of `seed_artists`,
	// `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists`
	// and `seed_tracks` are not set_.
	SeedGenres param.Opt[string] `query:"seed_genres,omitzero" json:"-"`
	// A comma separated list of
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for a seed
	// track. Up to 5 seed values may be provided in any combination of `seed_artists`,
	// `seed_tracks` and `seed_genres`.<br/> _**Note**: only required if `seed_artists`
	// and `seed_genres` are not set_.
	SeedTracks param.Opt[string] `query:"seed_tracks,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetAcousticness param.Opt[float64] `query:"target_acousticness,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetDanceability param.Opt[float64] `query:"target_danceability,omitzero" json:"-"`
	// Target duration of the track (ms)
	TargetDurationMs param.Opt[int64] `query:"target_duration_ms,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetEnergy param.Opt[float64] `query:"target_energy,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetInstrumentalness param.Opt[float64] `query:"target_instrumentalness,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetKey param.Opt[int64] `query:"target_key,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetLiveness param.Opt[float64] `query:"target_liveness,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetLoudness param.Opt[float64] `query:"target_loudness,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetMode param.Opt[int64] `query:"target_mode,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetPopularity param.Opt[int64] `query:"target_popularity,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetSpeechiness param.Opt[float64] `query:"target_speechiness,omitzero" json:"-"`
	// Target tempo (BPM)
	TargetTempo param.Opt[float64] `query:"target_tempo,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetTimeSignature param.Opt[int64] `query:"target_time_signature,omitzero" json:"-"`
	// For each of the tunable track attributes (below) a target value may be provided.
	// Tracks with the attribute values nearest to the target values will be preferred.
	// For example, you might request `target_energy=0.6` and
	// `target_danceability=0.8`. All target values will be weighed equally in ranking
	// results.
	TargetValence param.Opt[float64] `query:"target_valence,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecommendationGetParams]'s query parameters as
// `url.Values`.
func (r RecommendationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
