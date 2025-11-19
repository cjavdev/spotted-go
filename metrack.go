// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
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

// MeTrackService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeTrackService] method instead.
type MeTrackService struct {
	Options []option.RequestOption
}

// NewMeTrackService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeTrackService(opts ...option.RequestOption) (r MeTrackService) {
	r = MeTrackService{}
	r.Options = opts
	return
}

// Get a list of the songs saved in the current Spotify user's 'Your Music'
// library.
func (r *MeTrackService) List(ctx context.Context, query MeTrackListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MeTrackListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/tracks"
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

// Get a list of the songs saved in the current Spotify user's 'Your Music'
// library.
func (r *MeTrackService) ListAutoPaging(ctx context.Context, query MeTrackListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MeTrackListResponse] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

// Check if one or more tracks is already saved in the current Spotify user's 'Your
// Music' library.
func (r *MeTrackService) Check(ctx context.Context, query MeTrackCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/tracks/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove one or more tracks from the current user's 'Your Music' library.
func (r *MeTrackService) Remove(ctx context.Context, body MeTrackRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/tracks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Save one or more tracks to the current user's 'Your Music' library.
func (r *MeTrackService) Save(ctx context.Context, body MeTrackSaveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/tracks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type MeTrackListResponse struct {
	// The date and time the track was saved. Timestamps are returned in ISO 8601
	// format as Coordinated Universal Time (UTC) with a zero offset:
	// YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an
	// album release), an additional field indicates the precision; see for example,
	// release_date in an album object.
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Information about the track.
	Track shared.TrackObject `json:"track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		Track       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeTrackListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeTrackListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeTrackListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
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
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeTrackListParams]'s query parameters as `url.Values`.
func (r MeTrackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeTrackCheckParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=4iV5W9uYEdYUVa79Axb7Rh,1301WleyT98MSxVHPZCA6M`. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeTrackCheckParams]'s query parameters as `url.Values`.
func (r MeTrackCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeTrackRemoveParams struct {
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `["4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M"]`<br/>A maximum of 50 items
	// can be specified in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeTrackRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeTrackRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeTrackRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeTrackSaveParams struct {
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `["4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M"]`<br/>A maximum of 50 items
	// can be specified in one request. _**Note**: if the `timestamped_ids` is present
	// in the body, any IDs listed in the query parameters (deprecated) or the `ids`
	// field in the body will be ignored._
	IDs []string `json:"ids,omitzero,required"`
	// A JSON array of objects containing track IDs with their corresponding
	// timestamps. Each object must include a track ID and an `added_at` timestamp.
	// This allows you to specify when tracks were added to maintain a specific
	// chronological order in the user's library.<br/>A maximum of 50 items can be
	// specified in one request. _**Note**: if the `timestamped_ids` is present in the
	// body, any IDs listed in the query parameters (deprecated) or the `ids` field in
	// the body will be ignored._
	TimestampedIDs []MeTrackSaveParamsTimestampedID `json:"timestamped_ids,omitzero"`
	paramObj
}

func (r MeTrackSaveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeTrackSaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeTrackSaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, AddedAt are required.
type MeTrackSaveParamsTimestampedID struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	ID string `json:"id,required"`
	// The timestamp when the track was added to the library. Use ISO 8601 format with
	// UTC timezone (e.g., `2023-01-15T14:30:00Z`). You can specify past timestamps to
	// insert tracks at specific positions in the library's chronological order. The
	// API uses minute-level granularity for ordering, though the timestamp supports
	// millisecond precision.
	AddedAt time.Time `json:"added_at,required" format:"date-time"`
	paramObj
}

func (r MeTrackSaveParamsTimestampedID) MarshalJSON() (data []byte, err error) {
	type shadow MeTrackSaveParamsTimestampedID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeTrackSaveParamsTimestampedID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
