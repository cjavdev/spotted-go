// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/respjson"
)

// AudioFeatureService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioFeatureService] method instead.
type AudioFeatureService struct {
	Options []option.RequestOption
}

// NewAudioFeatureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioFeatureService(opts ...option.RequestOption) (r AudioFeatureService) {
	r = AudioFeatureService{}
	r.Options = opts
	return
}

// Get audio feature information for a single track identified by its unique
// Spotify ID.
//
// Deprecated: deprecated
func (r *AudioFeatureService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AudioFeatureGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("audio-features/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get audio features for multiple tracks based on their Spotify IDs.
//
// Deprecated: deprecated
func (r *AudioFeatureService) BulkGet(ctx context.Context, query AudioFeatureBulkGetParams, opts ...option.RequestOption) (res *AudioFeatureBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audio-features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AudioFeatureGetResponse struct {
	// The Spotify ID for the track.
	ID string `json:"id"`
	// A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0
	// represents high confidence the track is acoustic.
	Acousticness float64 `json:"acousticness"`
	// A URL to access the full audio analysis of this track. An access token is
	// required to access this data.
	AnalysisURL string `json:"analysis_url"`
	// Danceability describes how suitable a track is for dancing based on a
	// combination of musical elements including tempo, rhythm stability, beat
	// strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is
	// most danceable.
	Danceability float64 `json:"danceability"`
	// The duration of the track in milliseconds.
	DurationMs int64 `json:"duration_ms"`
	// Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of
	// intensity and activity. Typically, energetic tracks feel fast, loud, and noisy.
	// For example, death metal has high energy, while a Bach prelude scores low on the
	// scale. Perceptual features contributing to this attribute include dynamic range,
	// perceived loudness, timbre, onset rate, and general entropy.
	Energy float64 `json:"energy"`
	// Predicts whether a track contains no vocals. "Ooh" and "aah" sounds are treated
	// as instrumental in this context. Rap or spoken word tracks are clearly "vocal".
	// The closer the instrumentalness value is to 1.0, the greater likelihood the
	// track contains no vocal content. Values above 0.5 are intended to represent
	// instrumental tracks, but confidence is higher as the value approaches 1.0.
	Instrumentalness float64 `json:"instrumentalness"`
	// The key the track is in. Integers map to pitches using standard
	// [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1
	// = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.
	Key int64 `json:"key"`
	// Detects the presence of an audience in the recording. Higher liveness values
	// represent an increased probability that the track was performed live. A value
	// above 0.8 provides strong likelihood that the track is live.
	Liveness float64 `json:"liveness"`
	// The overall loudness of a track in decibels (dB). Loudness values are averaged
	// across the entire track and are useful for comparing relative loudness of
	// tracks. Loudness is the quality of a sound that is the primary psychological
	// correlate of physical strength (amplitude). Values typically range between -60
	// and 0 db.
	Loudness float64 `json:"loudness"`
	// Mode indicates the modality (major or minor) of a track, the type of scale from
	// which its melodic content is derived. Major is represented by 1 and minor is 0.
	Mode int64 `json:"mode"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Speechiness detects the presence of spoken words in a track. The more
	// exclusively speech-like the recording (e.g. talk show, audio book, poetry), the
	// closer to 1.0 the attribute value. Values above 0.66 describe tracks that are
	// probably made entirely of spoken words. Values between 0.33 and 0.66 describe
	// tracks that may contain both music and speech, either in sections or layered,
	// including such cases as rap music. Values below 0.33 most likely represent music
	// and other non-speech-like tracks.
	Speechiness float64 `json:"speechiness"`
	// The overall estimated tempo of a track in beats per minute (BPM). In musical
	// terminology, tempo is the speed or pace of a given piece and derives directly
	// from the average beat duration.
	Tempo float64 `json:"tempo"`
	// An estimated time signature. The time signature (meter) is a notational
	// convention to specify how many beats are in each bar (or measure). The time
	// signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int64 `json:"time_signature"`
	// A link to the Web API endpoint providing full details of the track.
	TrackHref string `json:"track_href"`
	// The object type.
	//
	// Any of "audio_features".
	Type AudioFeatureGetResponseType `json:"type"`
	// The Spotify URI for the track.
	Uri string `json:"uri"`
	// A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a
	// track. Tracks with high valence sound more positive (e.g. happy, cheerful,
	// euphoric), while tracks with low valence sound more negative (e.g. sad,
	// depressed, angry).
	Valence float64 `json:"valence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Acousticness     respjson.Field
		AnalysisURL      respjson.Field
		Danceability     respjson.Field
		DurationMs       respjson.Field
		Energy           respjson.Field
		Instrumentalness respjson.Field
		Key              respjson.Field
		Liveness         respjson.Field
		Loudness         respjson.Field
		Mode             respjson.Field
		Published        respjson.Field
		Speechiness      respjson.Field
		Tempo            respjson.Field
		TimeSignature    respjson.Field
		TrackHref        respjson.Field
		Type             respjson.Field
		Uri              respjson.Field
		Valence          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioFeatureGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioFeatureGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type AudioFeatureGetResponseType string

const (
	AudioFeatureGetResponseTypeAudioFeatures AudioFeatureGetResponseType = "audio_features"
)

type AudioFeatureBulkGetResponse struct {
	AudioFeatures []AudioFeatureBulkGetResponseAudioFeature `json:"audio_features,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioFeatures respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioFeatureBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioFeatureBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioFeatureBulkGetResponseAudioFeature struct {
	// The Spotify ID for the track.
	ID string `json:"id"`
	// A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0
	// represents high confidence the track is acoustic.
	Acousticness float64 `json:"acousticness"`
	// A URL to access the full audio analysis of this track. An access token is
	// required to access this data.
	AnalysisURL string `json:"analysis_url"`
	// Danceability describes how suitable a track is for dancing based on a
	// combination of musical elements including tempo, rhythm stability, beat
	// strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is
	// most danceable.
	Danceability float64 `json:"danceability"`
	// The duration of the track in milliseconds.
	DurationMs int64 `json:"duration_ms"`
	// Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of
	// intensity and activity. Typically, energetic tracks feel fast, loud, and noisy.
	// For example, death metal has high energy, while a Bach prelude scores low on the
	// scale. Perceptual features contributing to this attribute include dynamic range,
	// perceived loudness, timbre, onset rate, and general entropy.
	Energy float64 `json:"energy"`
	// Predicts whether a track contains no vocals. "Ooh" and "aah" sounds are treated
	// as instrumental in this context. Rap or spoken word tracks are clearly "vocal".
	// The closer the instrumentalness value is to 1.0, the greater likelihood the
	// track contains no vocal content. Values above 0.5 are intended to represent
	// instrumental tracks, but confidence is higher as the value approaches 1.0.
	Instrumentalness float64 `json:"instrumentalness"`
	// The key the track is in. Integers map to pitches using standard
	// [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1
	// = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.
	Key int64 `json:"key"`
	// Detects the presence of an audience in the recording. Higher liveness values
	// represent an increased probability that the track was performed live. A value
	// above 0.8 provides strong likelihood that the track is live.
	Liveness float64 `json:"liveness"`
	// The overall loudness of a track in decibels (dB). Loudness values are averaged
	// across the entire track and are useful for comparing relative loudness of
	// tracks. Loudness is the quality of a sound that is the primary psychological
	// correlate of physical strength (amplitude). Values typically range between -60
	// and 0 db.
	Loudness float64 `json:"loudness"`
	// Mode indicates the modality (major or minor) of a track, the type of scale from
	// which its melodic content is derived. Major is represented by 1 and minor is 0.
	Mode int64 `json:"mode"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Speechiness detects the presence of spoken words in a track. The more
	// exclusively speech-like the recording (e.g. talk show, audio book, poetry), the
	// closer to 1.0 the attribute value. Values above 0.66 describe tracks that are
	// probably made entirely of spoken words. Values between 0.33 and 0.66 describe
	// tracks that may contain both music and speech, either in sections or layered,
	// including such cases as rap music. Values below 0.33 most likely represent music
	// and other non-speech-like tracks.
	Speechiness float64 `json:"speechiness"`
	// The overall estimated tempo of a track in beats per minute (BPM). In musical
	// terminology, tempo is the speed or pace of a given piece and derives directly
	// from the average beat duration.
	Tempo float64 `json:"tempo"`
	// An estimated time signature. The time signature (meter) is a notational
	// convention to specify how many beats are in each bar (or measure). The time
	// signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int64 `json:"time_signature"`
	// A link to the Web API endpoint providing full details of the track.
	TrackHref string `json:"track_href"`
	// The object type.
	//
	// Any of "audio_features".
	Type string `json:"type"`
	// The Spotify URI for the track.
	Uri string `json:"uri"`
	// A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a
	// track. Tracks with high valence sound more positive (e.g. happy, cheerful,
	// euphoric), while tracks with low valence sound more negative (e.g. sad,
	// depressed, angry).
	Valence float64 `json:"valence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Acousticness     respjson.Field
		AnalysisURL      respjson.Field
		Danceability     respjson.Field
		DurationMs       respjson.Field
		Energy           respjson.Field
		Instrumentalness respjson.Field
		Key              respjson.Field
		Liveness         respjson.Field
		Loudness         respjson.Field
		Mode             respjson.Field
		Published        respjson.Field
		Speechiness      respjson.Field
		Tempo            respjson.Field
		TimeSignature    respjson.Field
		TrackHref        respjson.Field
		Type             respjson.Field
		Uri              respjson.Field
		Valence          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioFeatureBulkGetResponseAudioFeature) RawJSON() string { return r.JSON.raw }
func (r *AudioFeatureBulkGetResponseAudioFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioFeatureBulkGetParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the tracks.
	// Maximum: 100 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [AudioFeatureBulkGetParams]'s query parameters as
// `url.Values`.
func (r AudioFeatureBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
