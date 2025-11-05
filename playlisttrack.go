// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/spotted-go/internal/apijson"
	"github.com/stainless-sdks/spotted-go/internal/apiquery"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/pagination"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-go/packages/respjson"
	"github.com/stainless-sdks/spotted-go/shared"
)

// PlaylistTrackService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistTrackService] method instead.
type PlaylistTrackService struct {
	Options []option.RequestOption
}

// NewPlaylistTrackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaylistTrackService(opts ...option.RequestOption) (r PlaylistTrackService) {
	r = PlaylistTrackService{}
	r.Options = opts
	return
}

// Either reorder or replace items in a playlist depending on the request's
// parameters. To reorder items, include `range_start`, `insert_before`,
// `range_length` and `snapshot_id` in the request's body. To replace items,
// include `uris` as either a query parameter or in the request's body. Replacing
// items in a playlist will overwrite its existing items. This operation can be
// used for replacing or clearing items in a playlist. <br/> **Note**: Replace and
// reorder are mutually exclusive operations which share the same endpoint, but
// have different parameters. These operations can't be applied together in a
// single request.
func (r *PlaylistTrackService) Update(ctx context.Context, playlistID string, body PlaylistTrackUpdateParams, opts ...option.RequestOption) (res *PlaylistTrackUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/tracks", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get full details of the items of a playlist owned by a Spotify user.
func (r *PlaylistTrackService) List(ctx context.Context, playlistID string, query PlaylistTrackListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.PlaylistTrackObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/tracks", playlistID)
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

// Get full details of the items of a playlist owned by a Spotify user.
func (r *PlaylistTrackService) ListAutoPaging(ctx context.Context, playlistID string, query PlaylistTrackListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.PlaylistTrackObject] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, playlistID, query, opts...))
}

// Add one or more items to a user's playlist.
func (r *PlaylistTrackService) Add(ctx context.Context, playlistID string, body PlaylistTrackAddParams, opts ...option.RequestOption) (res *PlaylistTrackAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/tracks", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove one or more items from a user's playlist.
func (r *PlaylistTrackService) Remove(ctx context.Context, playlistID string, body PlaylistTrackRemoveParams, opts ...option.RequestOption) (res *PlaylistTrackRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/tracks", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type PlaylistTrackUpdateResponse struct {
	SnapshotID string `json:"snapshot_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SnapshotID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistTrackUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PlaylistTrackUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackAddResponse struct {
	SnapshotID string `json:"snapshot_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SnapshotID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistTrackAddResponse) RawJSON() string { return r.JSON.raw }
func (r *PlaylistTrackAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackRemoveResponse struct {
	SnapshotID string `json:"snapshot_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SnapshotID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistTrackRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *PlaylistTrackRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackUpdateParams struct {
	// The position where the items should be inserted.<br/>To reorder the items to the
	// end of the playlist, simply set _insert_before_ to the position after the last
	// item.<br/>Examples:<br/>To reorder the first item to the last position in a
	// playlist with 10 items, set _range_start_ to 0, and _insert_before_
	// to 10.<br/>To reorder the last item in a playlist with 10 items to the start of
	// the playlist, set _range_start_ to 9, and _insert_before_ to 0.
	InsertBefore param.Opt[int64] `json:"insert_before,omitzero"`
	// The amount of items to be reordered. Defaults to 1 if not set.<br/>The range of
	// items to be reordered begins from the _range_start_ position, and includes the
	// _range_length_ subsequent items.<br/>Example:<br/>To move the items at index
	// 9-10 to the start of the playlist, _range_start_ is set to 9, and _range_length_
	// is set to 2.
	RangeLength param.Opt[int64] `json:"range_length,omitzero"`
	// The position of the first item to be reordered.
	RangeStart param.Opt[int64] `json:"range_start,omitzero"`
	// The playlist's snapshot ID against which you want to make the changes.
	SnapshotID param.Opt[string] `json:"snapshot_id,omitzero"`
	Uris       []string          `json:"uris,omitzero"`
	paramObj
}

func (r PlaylistTrackUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistTrackUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistTrackUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackListParams struct {
	// A comma-separated list of item types that your client supports besides the
	// default `track` type. Valid types are: `track` and `episode`.<br/> _**Note**:
	// This parameter was introduced to allow existing clients to maintain their
	// current behaviour and might be deprecated in the future._<br/> In addition to
	// providing this parameter, make sure that your client properly handles cases of
	// new types in the future by checking against the `type` field of each object.
	AdditionalTypes param.Opt[string] `query:"additional_types,omitzero" json:"-"`
	// Filters for the query: a comma-separated list of the fields to return. If
	// omitted, all fields are returned. For example, to get just the total number of
	// items and the request limit:<br/>`fields=total,limit`<br/>A dot separator can be
	// used to specify non-reoccurring fields, while parentheses can be used to specify
	// reoccurring fields within objects. For example, to get just the added date and
	// user ID of the adder:<br/>`fields=items(added_at,added_by.id)`<br/>Use multiple
	// parentheses to drill down into nested objects, for
	// example:<br/>`fields=items(track(name,href,album(name,href)))`<br/>Fields can be
	// excluded by prefixing them with an exclamation mark, for
	// example:<br/>`fields=items.track.album(!external_urls,images)`
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 100.
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

// URLQuery serializes [PlaylistTrackListParams]'s query parameters as
// `url.Values`.
func (r PlaylistTrackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlaylistTrackAddParams struct {
	// The position to insert the items, a zero-based index. For example, to insert the
	// items in the first position: `position=0` ; to insert the items in the third
	// position: `position=2`. If omitted, the items will be appended to the playlist.
	// Items are added in the order they appear in the uris array. For example:
	// `{"uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh","spotify:track:1301WleyT98MSxVHPZCA6M"], "position": 3}`
	Position param.Opt[int64] `json:"position,omitzero"`
	// A JSON array of the
	// [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) to add. For
	// example:
	// `{"uris": ["spotify:track:4iV5W9uYEdYUVa79Axb7Rh","spotify:track:1301WleyT98MSxVHPZCA6M", "spotify:episode:512ojhOuo1ktJprKbVcKyQ"]}`<br/>A
	// maximum of 100 items can be added in one request. _**Note**: if the `uris`
	// parameter is present in the query string, any URIs listed here in the body will
	// be ignored._
	Uris []string `json:"uris,omitzero"`
	paramObj
}

func (r PlaylistTrackAddParams) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistTrackAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistTrackAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackRemoveParams struct {
	// An array of objects containing
	// [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) of the tracks
	// or episodes to remove. For example:
	// `{ "tracks": [{ "uri": "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" },{ "uri": "spotify:track:1301WleyT98MSxVHPZCA6M" }] }`.
	// A maximum of 100 objects can be sent at once.
	Tracks []PlaylistTrackRemoveParamsTrack `json:"tracks,omitzero,required"`
	// The playlist's snapshot ID against which you want to make the changes. The API
	// will validate that the specified items exist and in the specified positions and
	// make the changes, even if more recent changes have been made to the playlist.
	SnapshotID param.Opt[string] `json:"snapshot_id,omitzero"`
	paramObj
}

func (r PlaylistTrackRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistTrackRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistTrackRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackRemoveParamsTrack struct {
	// Spotify URI
	Uri param.Opt[string] `json:"uri,omitzero"`
	paramObj
}

func (r PlaylistTrackRemoveParamsTrack) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistTrackRemoveParamsTrack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistTrackRemoveParamsTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
