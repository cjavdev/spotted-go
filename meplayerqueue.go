// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"encoding/json"
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

// MePlayerQueueService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMePlayerQueueService] method instead.
type MePlayerQueueService struct {
	Options []option.RequestOption
}

// NewMePlayerQueueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMePlayerQueueService(opts ...option.RequestOption) (r MePlayerQueueService) {
	r = MePlayerQueueService{}
	r.Options = opts
	return
}

// Add an item to be played next in the user's current playback queue. This API
// only works for users who have Spotify Premium. The order of execution is not
// guaranteed when you use this API with other Player API endpoints.
func (r *MePlayerQueueService) Add(ctx context.Context, body MePlayerQueueAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/player/queue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the list of objects that make up the user's queue.
func (r *MePlayerQueueService) Get(ctx context.Context, opts ...option.RequestOption) (res *MePlayerQueueGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/player/queue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MePlayerQueueGetResponse struct {
	// The currently playing track or episode. Can be `null`.
	CurrentlyPlaying MePlayerQueueGetResponseCurrentlyPlayingUnion `json:"currently_playing"`
	// The tracks or episodes in the queue. Can be empty.
	Queue []MePlayerQueueGetResponseQueueUnion `json:"queue"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentlyPlaying respjson.Field
		Queue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerQueueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MePlayerQueueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerQueueGetResponseCurrentlyPlayingUnion contains all possible properties
// and values from [shared.TrackObject], [shared.EpisodeObject].
//
// Use the [MePlayerQueueGetResponseCurrentlyPlayingUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MePlayerQueueGetResponseCurrentlyPlayingUnion struct {
	ID string `json:"id"`
	// This field is from variant [shared.TrackObject].
	Album shared.TrackObjectAlbum `json:"album"`
	// This field is from variant [shared.TrackObject].
	Artists []shared.SimplifiedArtistObject `json:"artists"`
	// This field is from variant [shared.TrackObject].
	AvailableMarkets []string `json:"available_markets"`
	// This field is from variant [shared.TrackObject].
	DiscNumber int64 `json:"disc_number"`
	DurationMs int64 `json:"duration_ms"`
	Explicit   bool  `json:"explicit"`
	// This field is from variant [shared.TrackObject].
	ExternalIDs shared.ExternalIDObject `json:"external_ids"`
	// This field is from variant [shared.TrackObject].
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	Href         string                   `json:"href"`
	// This field is from variant [shared.TrackObject].
	IsLocal    bool `json:"is_local"`
	IsPlayable bool `json:"is_playable"`
	// This field is from variant [shared.TrackObject].
	LinkedFrom shared.LinkedTrackObject `json:"linked_from"`
	Name       string                   `json:"name"`
	// This field is from variant [shared.TrackObject].
	Popularity int64 `json:"popularity"`
	// This field is from variant [shared.TrackObject].
	PreviewURL string `json:"preview_url"`
	// This field is a union of [shared.TrackRestrictionObject],
	// [shared.EpisodeRestrictionObject]
	Restrictions MePlayerQueueGetResponseCurrentlyPlayingUnionRestrictions `json:"restrictions"`
	// This field is from variant [shared.TrackObject].
	TrackNumber int64 `json:"track_number"`
	// Any of "track", "episode".
	Type string `json:"type"`
	Uri  string `json:"uri"`
	// This field is from variant [shared.EpisodeObject].
	AudioPreviewURL string `json:"audio_preview_url"`
	// This field is from variant [shared.EpisodeObject].
	Description string `json:"description"`
	// This field is from variant [shared.EpisodeObject].
	HTMLDescription string `json:"html_description"`
	// This field is from variant [shared.EpisodeObject].
	Images []shared.ImageObject `json:"images"`
	// This field is from variant [shared.EpisodeObject].
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// This field is from variant [shared.EpisodeObject].
	Languages []string `json:"languages"`
	// This field is from variant [shared.EpisodeObject].
	ReleaseDate string `json:"release_date"`
	// This field is from variant [shared.EpisodeObject].
	ReleaseDatePrecision shared.EpisodeObjectReleaseDatePrecision `json:"release_date_precision"`
	// This field is from variant [shared.EpisodeObject].
	Show shared.ShowBase `json:"show"`
	// This field is from variant [shared.EpisodeObject].
	Language string `json:"language"`
	// This field is from variant [shared.EpisodeObject].
	ResumePoint shared.ResumePointObject `json:"resume_point"`
	JSON        struct {
		ID                   respjson.Field
		Album                respjson.Field
		Artists              respjson.Field
		AvailableMarkets     respjson.Field
		DiscNumber           respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalIDs          respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		IsLocal              respjson.Field
		IsPlayable           respjson.Field
		LinkedFrom           respjson.Field
		Name                 respjson.Field
		Popularity           respjson.Field
		PreviewURL           respjson.Field
		Restrictions         respjson.Field
		TrackNumber          respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		Languages            respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Show                 respjson.Field
		Language             respjson.Field
		ResumePoint          respjson.Field
		raw                  string
	} `json:"-"`
}

// anyMePlayerQueueGetResponseCurrentlyPlaying is implemented by each variant of
// [MePlayerQueueGetResponseCurrentlyPlayingUnion] to add type safety for the
// return type of [MePlayerQueueGetResponseCurrentlyPlayingUnion.AsAny]
type anyMePlayerQueueGetResponseCurrentlyPlaying interface {
	ImplMePlayerQueueGetResponseCurrentlyPlayingUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MePlayerQueueGetResponseCurrentlyPlayingUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MePlayerQueueGetResponseCurrentlyPlayingUnion) AsAny() anyMePlayerQueueGetResponseCurrentlyPlaying {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u MePlayerQueueGetResponseCurrentlyPlayingUnion) AsTrack() (v shared.TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MePlayerQueueGetResponseCurrentlyPlayingUnion) AsEpisode() (v shared.EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MePlayerQueueGetResponseCurrentlyPlayingUnion) RawJSON() string { return u.JSON.raw }

func (r *MePlayerQueueGetResponseCurrentlyPlayingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerQueueGetResponseCurrentlyPlayingUnionRestrictions is an implicit
// subunion of [MePlayerQueueGetResponseCurrentlyPlayingUnion].
// MePlayerQueueGetResponseCurrentlyPlayingUnionRestrictions provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MePlayerQueueGetResponseCurrentlyPlayingUnion].
type MePlayerQueueGetResponseCurrentlyPlayingUnionRestrictions struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MePlayerQueueGetResponseCurrentlyPlayingUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerQueueGetResponseQueueUnion contains all possible properties and values
// from [shared.TrackObject], [shared.EpisodeObject].
//
// Use the [MePlayerQueueGetResponseQueueUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MePlayerQueueGetResponseQueueUnion struct {
	ID string `json:"id"`
	// This field is from variant [shared.TrackObject].
	Album shared.TrackObjectAlbum `json:"album"`
	// This field is from variant [shared.TrackObject].
	Artists []shared.SimplifiedArtistObject `json:"artists"`
	// This field is from variant [shared.TrackObject].
	AvailableMarkets []string `json:"available_markets"`
	// This field is from variant [shared.TrackObject].
	DiscNumber int64 `json:"disc_number"`
	DurationMs int64 `json:"duration_ms"`
	Explicit   bool  `json:"explicit"`
	// This field is from variant [shared.TrackObject].
	ExternalIDs shared.ExternalIDObject `json:"external_ids"`
	// This field is from variant [shared.TrackObject].
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	Href         string                   `json:"href"`
	// This field is from variant [shared.TrackObject].
	IsLocal    bool `json:"is_local"`
	IsPlayable bool `json:"is_playable"`
	// This field is from variant [shared.TrackObject].
	LinkedFrom shared.LinkedTrackObject `json:"linked_from"`
	Name       string                   `json:"name"`
	// This field is from variant [shared.TrackObject].
	Popularity int64 `json:"popularity"`
	// This field is from variant [shared.TrackObject].
	PreviewURL string `json:"preview_url"`
	// This field is a union of [shared.TrackRestrictionObject],
	// [shared.EpisodeRestrictionObject]
	Restrictions MePlayerQueueGetResponseQueueUnionRestrictions `json:"restrictions"`
	// This field is from variant [shared.TrackObject].
	TrackNumber int64 `json:"track_number"`
	// Any of "track", "episode".
	Type string `json:"type"`
	Uri  string `json:"uri"`
	// This field is from variant [shared.EpisodeObject].
	AudioPreviewURL string `json:"audio_preview_url"`
	// This field is from variant [shared.EpisodeObject].
	Description string `json:"description"`
	// This field is from variant [shared.EpisodeObject].
	HTMLDescription string `json:"html_description"`
	// This field is from variant [shared.EpisodeObject].
	Images []shared.ImageObject `json:"images"`
	// This field is from variant [shared.EpisodeObject].
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// This field is from variant [shared.EpisodeObject].
	Languages []string `json:"languages"`
	// This field is from variant [shared.EpisodeObject].
	ReleaseDate string `json:"release_date"`
	// This field is from variant [shared.EpisodeObject].
	ReleaseDatePrecision shared.EpisodeObjectReleaseDatePrecision `json:"release_date_precision"`
	// This field is from variant [shared.EpisodeObject].
	Show shared.ShowBase `json:"show"`
	// This field is from variant [shared.EpisodeObject].
	Language string `json:"language"`
	// This field is from variant [shared.EpisodeObject].
	ResumePoint shared.ResumePointObject `json:"resume_point"`
	JSON        struct {
		ID                   respjson.Field
		Album                respjson.Field
		Artists              respjson.Field
		AvailableMarkets     respjson.Field
		DiscNumber           respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalIDs          respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		IsLocal              respjson.Field
		IsPlayable           respjson.Field
		LinkedFrom           respjson.Field
		Name                 respjson.Field
		Popularity           respjson.Field
		PreviewURL           respjson.Field
		Restrictions         respjson.Field
		TrackNumber          respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		Languages            respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Show                 respjson.Field
		Language             respjson.Field
		ResumePoint          respjson.Field
		raw                  string
	} `json:"-"`
}

// anyMePlayerQueueGetResponseQueue is implemented by each variant of
// [MePlayerQueueGetResponseQueueUnion] to add type safety for the return type of
// [MePlayerQueueGetResponseQueueUnion.AsAny]
type anyMePlayerQueueGetResponseQueue interface {
	ImplMePlayerQueueGetResponseQueueUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MePlayerQueueGetResponseQueueUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MePlayerQueueGetResponseQueueUnion) AsAny() anyMePlayerQueueGetResponseQueue {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u MePlayerQueueGetResponseQueueUnion) AsTrack() (v shared.TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MePlayerQueueGetResponseQueueUnion) AsEpisode() (v shared.EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MePlayerQueueGetResponseQueueUnion) RawJSON() string { return u.JSON.raw }

func (r *MePlayerQueueGetResponseQueueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerQueueGetResponseQueueUnionRestrictions is an implicit subunion of
// [MePlayerQueueGetResponseQueueUnion].
// MePlayerQueueGetResponseQueueUnionRestrictions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MePlayerQueueGetResponseQueueUnion].
type MePlayerQueueGetResponseQueueUnionRestrictions struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MePlayerQueueGetResponseQueueUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerQueueAddParams struct {
	// The uri of the item to add to the queue. Must be a track or an episode uri.
	Uri string `query:"uri,required" json:"-"`
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerQueueAddParams]'s query parameters as `url.Values`.
func (r MePlayerQueueAddParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
