// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/pagination"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// MePlayerService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMePlayerService] method instead.
type MePlayerService struct {
	Options []option.RequestOption
	Queue   MePlayerQueueService
}

// NewMePlayerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMePlayerService(opts ...option.RequestOption) (r MePlayerService) {
	r = MePlayerService{}
	r.Options = opts
	r.Queue = NewMePlayerQueueService(opts...)
	return
}

// Get the object currently being played on the user's Spotify account.
func (r *MePlayerService) GetCurrentlyPlaying(ctx context.Context, query MePlayerGetCurrentlyPlayingParams, opts ...option.RequestOption) (res *MePlayerGetCurrentlyPlayingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/player/currently-playing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get information about a user’s available Spotify Connect devices. Some device
// models are not supported and will not be listed in the API response.
func (r *MePlayerService) GetDevices(ctx context.Context, opts ...option.RequestOption) (res *MePlayerGetDevicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/player/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get information about the user’s current playback state, including track or
// episode, progress, and active device.
func (r *MePlayerService) GetState(ctx context.Context, query MePlayerGetStateParams, opts ...option.RequestOption) (res *MePlayerGetStateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/player"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get tracks from the current user's recently played tracks. _**Note**: Currently
// doesn't support podcast episodes._
func (r *MePlayerService) ListRecentlyPlayed(ctx context.Context, query MePlayerListRecentlyPlayedParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MePlayerListRecentlyPlayedResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/player/recently-played"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get tracks from the current user's recently played tracks. _**Note**: Currently
// doesn't support podcast episodes._
func (r *MePlayerService) ListRecentlyPlayedAutoPaging(ctx context.Context, query MePlayerListRecentlyPlayedParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MePlayerListRecentlyPlayedResponse] {
	return pagination.NewCursorURLPageAutoPager(r.ListRecentlyPlayed(ctx, query, opts...))
}

// Pause playback on the user's account. This API only works for users who have
// Spotify Premium. The order of execution is not guaranteed when you use this API
// with other Player API endpoints.
func (r *MePlayerService) PausePlayback(ctx context.Context, body MePlayerPausePlaybackParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/pause"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Seeks to the given position in the user’s currently playing track. This API only
// works for users who have Spotify Premium. The order of execution is not
// guaranteed when you use this API with other Player API endpoints.
func (r *MePlayerService) SeekToPosition(ctx context.Context, body MePlayerSeekToPositionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/seek"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Set the repeat mode for the user's playback. This API only works for users who
// have Spotify Premium. The order of execution is not guaranteed when you use this
// API with other Player API endpoints.
func (r *MePlayerService) SetRepeatMode(ctx context.Context, body MePlayerSetRepeatModeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/repeat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Set the volume for the user’s current playback device. This API only works for
// users who have Spotify Premium. The order of execution is not guaranteed when
// you use this API with other Player API endpoints.
func (r *MePlayerService) SetVolume(ctx context.Context, body MePlayerSetVolumeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/volume"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Skips to next track in the user’s queue. This API only works for users who have
// Spotify Premium. The order of execution is not guaranteed when you use this API
// with other Player API endpoints.
func (r *MePlayerService) SkipNext(ctx context.Context, body MePlayerSkipNextParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/next"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Skips to previous track in the user’s queue. This API only works for users who
// have Spotify Premium. The order of execution is not guaranteed when you use this
// API with other Player API endpoints.
func (r *MePlayerService) SkipPrevious(ctx context.Context, body MePlayerSkipPreviousParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/previous"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Start a new context or resume current playback on the user's active device. This
// API only works for users who have Spotify Premium. The order of execution is not
// guaranteed when you use this API with other Player API endpoints.
func (r *MePlayerService) StartPlayback(ctx context.Context, params MePlayerStartPlaybackParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/play"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Toggle shuffle on or off for user’s playback. This API only works for users who
// have Spotify Premium. The order of execution is not guaranteed when you use this
// API with other Player API endpoints.
func (r *MePlayerService) ToggleShuffle(ctx context.Context, body MePlayerToggleShuffleParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player/shuffle"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Transfer playback to a new device and optionally begin playback. This API only
// works for users who have Spotify Premium. The order of execution is not
// guaranteed when you use this API with other Player API endpoints.
func (r *MePlayerService) Transfer(ctx context.Context, body MePlayerTransferParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/player"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type ContextObject struct {
	// External URLs for this context.
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The object type, e.g. "artist", "playlist", "album", "show".
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// context.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalURLs respjson.Field
		Href         respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextObject) RawJSON() string { return r.JSON.raw }
func (r *ContextObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceObject struct {
	// The device ID. This ID is unique and persistent to some extent. However, this is
	// not guaranteed and any cached `device_id` should periodically be cleared out and
	// refetched as necessary.
	ID string `json:"id,nullable"`
	// If this device is the currently active device.
	IsActive bool `json:"is_active"`
	// If this device is currently in a private session.
	IsPrivateSession bool `json:"is_private_session"`
	// Whether controlling this device is restricted. At present if this is "true" then
	// no Web API commands will be accepted by this device.
	IsRestricted bool `json:"is_restricted"`
	// A human-readable name for the device. Some devices have a name that the user can
	// configure (e.g. \"Loudest speaker\") and some devices have a generic name
	// associated with the manufacturer or device model.
	Name string `json:"name"`
	// If this device can be used to set the volume.
	SupportsVolume bool `json:"supports_volume"`
	// Device type, such as "computer", "smartphone" or "speaker".
	Type string `json:"type"`
	// The current volume in percent.
	VolumePercent int64 `json:"volume_percent,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsActive         respjson.Field
		IsPrivateSession respjson.Field
		IsRestricted     respjson.Field
		Name             respjson.Field
		SupportsVolume   respjson.Field
		Type             respjson.Field
		VolumePercent    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceObject) RawJSON() string { return r.JSON.raw }
func (r *DeviceObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerGetCurrentlyPlayingResponse struct {
	// Allows to update the user interface based on which playback actions are
	// available within the current context.
	Actions MePlayerGetCurrentlyPlayingResponseActions `json:"actions"`
	// A Context Object. Can be `null`.
	Context ContextObject `json:"context"`
	// The object type of the currently playing item. Can be one of `track`, `episode`,
	// `ad` or `unknown`.
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// If something is currently playing, return `true`.
	IsPlaying bool `json:"is_playing"`
	// The currently playing track or episode. Can be `null`.
	Item MePlayerGetCurrentlyPlayingResponseItemUnion `json:"item"`
	// Progress into the currently playing track or episode. Can be `null`.
	ProgressMs int64 `json:"progress_ms"`
	// Unix Millisecond Timestamp when data was fetched
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions              respjson.Field
		Context              respjson.Field
		CurrentlyPlayingType respjson.Field
		IsPlaying            respjson.Field
		Item                 respjson.Field
		ProgressMs           respjson.Field
		Timestamp            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerGetCurrentlyPlayingResponse) RawJSON() string { return r.JSON.raw }
func (r *MePlayerGetCurrentlyPlayingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to update the user interface based on which playback actions are
// available within the current context.
type MePlayerGetCurrentlyPlayingResponseActions struct {
	// Interrupting playback. Optional field.
	InterruptingPlayback bool `json:"interrupting_playback"`
	// Pausing. Optional field.
	Pausing bool `json:"pausing"`
	// Resuming. Optional field.
	Resuming bool `json:"resuming"`
	// Seeking playback location. Optional field.
	Seeking bool `json:"seeking"`
	// Skipping to the next context. Optional field.
	SkippingNext bool `json:"skipping_next"`
	// Skipping to the previous context. Optional field.
	SkippingPrev bool `json:"skipping_prev"`
	// Toggling repeat context flag. Optional field.
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	// Toggling repeat track flag. Optional field.
	TogglingRepeatTrack bool `json:"toggling_repeat_track"`
	// Toggling shuffle flag. Optional field.
	TogglingShuffle bool `json:"toggling_shuffle"`
	// Transfering playback between devices. Optional field.
	TransferringPlayback bool `json:"transferring_playback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterruptingPlayback  respjson.Field
		Pausing               respjson.Field
		Resuming              respjson.Field
		Seeking               respjson.Field
		SkippingNext          respjson.Field
		SkippingPrev          respjson.Field
		TogglingRepeatContext respjson.Field
		TogglingRepeatTrack   respjson.Field
		TogglingShuffle       respjson.Field
		TransferringPlayback  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerGetCurrentlyPlayingResponseActions) RawJSON() string { return r.JSON.raw }
func (r *MePlayerGetCurrentlyPlayingResponseActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerGetCurrentlyPlayingResponseItemUnion contains all possible properties
// and values from [shared.TrackObject], [shared.EpisodeObject].
//
// Use the [MePlayerGetCurrentlyPlayingResponseItemUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MePlayerGetCurrentlyPlayingResponseItemUnion struct {
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
	Restrictions MePlayerGetCurrentlyPlayingResponseItemUnionRestrictions `json:"restrictions"`
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

// anyMePlayerGetCurrentlyPlayingResponseItem is implemented by each variant of
// [MePlayerGetCurrentlyPlayingResponseItemUnion] to add type safety for the return
// type of [MePlayerGetCurrentlyPlayingResponseItemUnion.AsAny]
type anyMePlayerGetCurrentlyPlayingResponseItem interface {
	ImplMePlayerGetCurrentlyPlayingResponseItemUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MePlayerGetCurrentlyPlayingResponseItemUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MePlayerGetCurrentlyPlayingResponseItemUnion) AsAny() anyMePlayerGetCurrentlyPlayingResponseItem {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u MePlayerGetCurrentlyPlayingResponseItemUnion) AsTrack() (v shared.TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MePlayerGetCurrentlyPlayingResponseItemUnion) AsEpisode() (v shared.EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MePlayerGetCurrentlyPlayingResponseItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MePlayerGetCurrentlyPlayingResponseItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerGetCurrentlyPlayingResponseItemUnionRestrictions is an implicit subunion
// of [MePlayerGetCurrentlyPlayingResponseItemUnion].
// MePlayerGetCurrentlyPlayingResponseItemUnionRestrictions provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MePlayerGetCurrentlyPlayingResponseItemUnion].
type MePlayerGetCurrentlyPlayingResponseItemUnionRestrictions struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MePlayerGetCurrentlyPlayingResponseItemUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerGetDevicesResponse struct {
	Devices []DeviceObject `json:"devices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Devices     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerGetDevicesResponse) RawJSON() string { return r.JSON.raw }
func (r *MePlayerGetDevicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerGetStateResponse struct {
	// Allows to update the user interface based on which playback actions are
	// available within the current context.
	Actions MePlayerGetStateResponseActions `json:"actions"`
	// A Context Object. Can be `null`.
	Context ContextObject `json:"context"`
	// The object type of the currently playing item. Can be one of `track`, `episode`,
	// `ad` or `unknown`.
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// The device that is currently active.
	Device DeviceObject `json:"device"`
	// If something is currently playing, return `true`.
	IsPlaying bool `json:"is_playing"`
	// The currently playing track or episode. Can be `null`.
	Item MePlayerGetStateResponseItemUnion `json:"item"`
	// Progress into the currently playing track or episode. Can be `null`.
	ProgressMs int64 `json:"progress_ms"`
	// off, track, context
	RepeatState string `json:"repeat_state"`
	// If shuffle is on or off.
	ShuffleState bool `json:"shuffle_state"`
	// Unix Millisecond Timestamp when playback state was last changed (play, pause,
	// skip, scrub, new song, etc.).
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions              respjson.Field
		Context              respjson.Field
		CurrentlyPlayingType respjson.Field
		Device               respjson.Field
		IsPlaying            respjson.Field
		Item                 respjson.Field
		ProgressMs           respjson.Field
		RepeatState          respjson.Field
		ShuffleState         respjson.Field
		Timestamp            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerGetStateResponse) RawJSON() string { return r.JSON.raw }
func (r *MePlayerGetStateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows to update the user interface based on which playback actions are
// available within the current context.
type MePlayerGetStateResponseActions struct {
	// Interrupting playback. Optional field.
	InterruptingPlayback bool `json:"interrupting_playback"`
	// Pausing. Optional field.
	Pausing bool `json:"pausing"`
	// Resuming. Optional field.
	Resuming bool `json:"resuming"`
	// Seeking playback location. Optional field.
	Seeking bool `json:"seeking"`
	// Skipping to the next context. Optional field.
	SkippingNext bool `json:"skipping_next"`
	// Skipping to the previous context. Optional field.
	SkippingPrev bool `json:"skipping_prev"`
	// Toggling repeat context flag. Optional field.
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	// Toggling repeat track flag. Optional field.
	TogglingRepeatTrack bool `json:"toggling_repeat_track"`
	// Toggling shuffle flag. Optional field.
	TogglingShuffle bool `json:"toggling_shuffle"`
	// Transfering playback between devices. Optional field.
	TransferringPlayback bool `json:"transferring_playback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterruptingPlayback  respjson.Field
		Pausing               respjson.Field
		Resuming              respjson.Field
		Seeking               respjson.Field
		SkippingNext          respjson.Field
		SkippingPrev          respjson.Field
		TogglingRepeatContext respjson.Field
		TogglingRepeatTrack   respjson.Field
		TogglingShuffle       respjson.Field
		TransferringPlayback  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerGetStateResponseActions) RawJSON() string { return r.JSON.raw }
func (r *MePlayerGetStateResponseActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerGetStateResponseItemUnion contains all possible properties and values
// from [shared.TrackObject], [shared.EpisodeObject].
//
// Use the [MePlayerGetStateResponseItemUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MePlayerGetStateResponseItemUnion struct {
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
	Restrictions MePlayerGetStateResponseItemUnionRestrictions `json:"restrictions"`
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

// anyMePlayerGetStateResponseItem is implemented by each variant of
// [MePlayerGetStateResponseItemUnion] to add type safety for the return type of
// [MePlayerGetStateResponseItemUnion.AsAny]
type anyMePlayerGetStateResponseItem interface {
	ImplMePlayerGetStateResponseItemUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MePlayerGetStateResponseItemUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MePlayerGetStateResponseItemUnion) AsAny() anyMePlayerGetStateResponseItem {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u MePlayerGetStateResponseItemUnion) AsTrack() (v shared.TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MePlayerGetStateResponseItemUnion) AsEpisode() (v shared.EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MePlayerGetStateResponseItemUnion) RawJSON() string { return u.JSON.raw }

func (r *MePlayerGetStateResponseItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MePlayerGetStateResponseItemUnionRestrictions is an implicit subunion of
// [MePlayerGetStateResponseItemUnion].
// MePlayerGetStateResponseItemUnionRestrictions provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MePlayerGetStateResponseItemUnion].
type MePlayerGetStateResponseItemUnionRestrictions struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *MePlayerGetStateResponseItemUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerListRecentlyPlayedResponse struct {
	// The context the track was played from.
	Context ContextObject `json:"context"`
	// The date and time the track was played.
	PlayedAt time.Time `json:"played_at" format:"date-time"`
	// The track the user listened to.
	Track shared.TrackObject `json:"track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		PlayedAt    respjson.Field
		Track       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MePlayerListRecentlyPlayedResponse) RawJSON() string { return r.JSON.raw }
func (r *MePlayerListRecentlyPlayedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MePlayerGetCurrentlyPlayingParams struct {
	// A comma-separated list of item types that your client supports besides the
	// default `track` type. Valid types are: `track` and `episode`.<br/> _**Note**:
	// This parameter was introduced to allow existing clients to maintain their
	// current behaviour and might be deprecated in the future._<br/> In addition to
	// providing this parameter, make sure that your client properly handles cases of
	// new types in the future by checking against the `type` field of each object.
	AdditionalTypes param.Opt[string] `query:"additional_types,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [MePlayerGetCurrentlyPlayingParams]'s query parameters as
// `url.Values`.
func (r MePlayerGetCurrentlyPlayingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerGetStateParams struct {
	// A comma-separated list of item types that your client supports besides the
	// default `track` type. Valid types are: `track` and `episode`.<br/> _**Note**:
	// This parameter was introduced to allow existing clients to maintain their
	// current behaviour and might be deprecated in the future._<br/> In addition to
	// providing this parameter, make sure that your client properly handles cases of
	// new types in the future by checking against the `type` field of each object.
	AdditionalTypes param.Opt[string] `query:"additional_types,omitzero" json:"-"`
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
	paramObj
}

// URLQuery serializes [MePlayerGetStateParams]'s query parameters as `url.Values`.
func (r MePlayerGetStateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerListRecentlyPlayedParams struct {
	// A Unix timestamp in milliseconds. Returns all items after (but not including)
	// this cursor position. If `after` is specified, `before` must not be specified.
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// A Unix timestamp in milliseconds. Returns all items before (but not including)
	// this cursor position. If `before` is specified, `after` must not be specified.
	Before param.Opt[int64] `query:"before,omitzero" json:"-"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerListRecentlyPlayedParams]'s query parameters as
// `url.Values`.
func (r MePlayerListRecentlyPlayedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerPausePlaybackParams struct {
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerPausePlaybackParams]'s query parameters as
// `url.Values`.
func (r MePlayerPausePlaybackParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerSeekToPositionParams struct {
	// The position in milliseconds to seek to. Must be a positive number. Passing in a
	// position that is greater than the length of the track will cause the player to
	// start playing the next song.
	PositionMs int64 `query:"position_ms,required" json:"-"`
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerSeekToPositionParams]'s query parameters as
// `url.Values`.
func (r MePlayerSeekToPositionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerSetRepeatModeParams struct {
	// **track**, **context** or **off**.<br/> **track** will repeat the current
	// track.<br/> **context** will repeat the current context.<br/> **off** will turn
	// repeat off.
	State string `query:"state,required" json:"-"`
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerSetRepeatModeParams]'s query parameters as
// `url.Values`.
func (r MePlayerSetRepeatModeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerSetVolumeParams struct {
	// The volume to set. Must be a value from 0 to 100 inclusive.
	VolumePercent int64 `query:"volume_percent,required" json:"-"`
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerSetVolumeParams]'s query parameters as
// `url.Values`.
func (r MePlayerSetVolumeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerSkipNextParams struct {
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerSkipNextParams]'s query parameters as `url.Values`.
func (r MePlayerSkipNextParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerSkipPreviousParams struct {
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerSkipPreviousParams]'s query parameters as
// `url.Values`.
func (r MePlayerSkipPreviousParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerStartPlaybackParams struct {
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	// Optional. Spotify URI of the context to play. Valid contexts are albums, artists
	// & playlists. `{context_uri:"spotify:album:1Je1IMUlBXcx1Fz0WE7oPT"}`
	ContextUri param.Opt[string] `json:"context_uri,omitzero"`
	// Indicates from what position to start playback. Must be a positive number.
	// Passing in a position that is greater than the length of the track will cause
	// the player to start playing the next song.
	PositionMs param.Opt[int64] `json:"position_ms,omitzero"`
	// Optional. Indicates from where in the context playback should start. Only
	// available when context_uri corresponds to an album or playlist object "position"
	// is zero based and can’t be negative. Example: `"offset": {"position": 5}` "uri"
	// is a string representing the uri of the item to start at. Example:
	// `"offset": {"uri": "spotify:track:1301WleyT98MSxVHPZCA6M"}`
	Offset map[string]any `json:"offset,omitzero"`
	// Optional. A JSON array of the Spotify track URIs to play. For example:
	// `{"uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh", "spotify:track:1301WleyT98MSxVHPZCA6M"]}`
	Uris []string `json:"uris,omitzero"`
	paramObj
}

func (r MePlayerStartPlaybackParams) MarshalJSON() (data []byte, err error) {
	type shadow MePlayerStartPlaybackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MePlayerStartPlaybackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MePlayerStartPlaybackParams]'s query parameters as
// `url.Values`.
func (r MePlayerStartPlaybackParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerToggleShuffleParams struct {
	// **true** : Shuffle user's playback.<br/> **false** : Do not shuffle user's
	// playback.
	State bool `query:"state,required" json:"-"`
	// The id of the device this command is targeting. If not supplied, the user's
	// currently active device is the target.
	DeviceID param.Opt[string] `query:"device_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlayerToggleShuffleParams]'s query parameters as
// `url.Values`.
func (r MePlayerToggleShuffleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MePlayerTransferParams struct {
	// A JSON array containing the ID of the device on which playback should be
	// started/transferred.<br/>For
	// example:`{device_ids:["74ASZWbe4lXaubB36ztrGX"]}`<br/>_**Note**: Although an
	// array is accepted, only a single device_id is currently supported. Supplying
	// more than one will return `400 Bad Request`_
	DeviceIDs []string `json:"device_ids,omitzero,required"`
	// **true**: ensure playback happens on new device.<br/>**false** or not provided:
	// keep the current playback state.
	Play param.Opt[bool] `json:"play,omitzero"`
	paramObj
}

func (r MePlayerTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow MePlayerTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MePlayerTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
